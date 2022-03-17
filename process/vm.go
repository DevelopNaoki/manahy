package process

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// GetVmList get a list of VMs
func GetVmList() (vmList VmList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VM | Sort-Object State | Format-Table Name, State").Output()
	if err != nil {
		return vmList, err
	}
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Name") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			state := regexp.MustCompile("Running$|Saved$|Off$|Paused$").FindString(split[i])
			split[i] = regexp.MustCompile("Running$|Saved$|Off$|Paused$").ReplaceAllString(split[i], "")
			split[i] = strings.TrimSpace(split[i])

			switch state {
			case "Running":
				vmList.Running = append(vmList.Running, split[i])
			case "Saved":
				vmList.Saved = append(vmList.Saved, split[i])
			case "Off":
				vmList.Off = append(vmList.Off, split[i])
			case "Paused":
				vmList.Paused = append(vmList.Paused, split[i])
			default:
				return vmList, fmt.Errorf("Unknown error for vm list")
			}
		}
	}
	return
}

// GetVmState get a VM state
func GetVmState(name string) (state string) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VM '"+name+"' | Format-Table State").Output()
	if err != nil {
		state = "NotFound"
	} else {
		vmState := listingOfExecuteResults(res, "State")
		if len(vmState) == 1 {
			state = vmState[0]
		}
	}
	return
}

func SetVmProcessor(name string, cpu Cpu) error {
	var args string

	if GetVmState(name) != "NotFound" {
		args = "Set-VMProcessor " + name + " "
	}
	if cpu.Thread > 0 {
		args = args + "-Count " + strconv.Itoa(cpu.Thread) + " "
	}
	if cpu.Reserve <= 100 && cpu.Reserve >= 0 {
		args = args + " -Reserve " + strconv.Itoa(cpu.Reserve) + " "
	}
	if cpu.Maximum <= 100 && cpu.Maximum >= 0 {
		args = args + " -Maximum " + strconv.Itoa(cpu.Maximum) + " "
	}
	if cpu.RelativeWeight <= 10000 && cpu.RelativeWeight > 0 {
		args = args + " -RelativeWeight " + strconv.Itoa(cpu.RelativeWeight) + " "
	}
	args = args + " -ExposeVirtualizationExtensions " + strconv.FormatBool(cpu.Nested)

	err := exec.Command("powershell", "-NoProfile", args).Run()

	if err != nil {
		return err
	}
	return nil
}

func SetVmMemory(name string) {
}

func CreateVm(newVm Vm) error {
	err := CheckVmParam(newVm)
	if err != nil {
		return err
	}

	args := "New-VM -Name " + newVm.Name + " -Generation " + strconv.Itoa(newVm.Generation) + " -Path " + newVm.Path + " -MemoryStartupBytes " + newVm.Memory.Size
	err = exec.Command("powershell", "-NoProfile", args).Run()
	if err != nil {
		return err
	}
	SetVmProcessor(newVm.Name, newVm.Cpu)

	return nil
}

func RemoveVm(name string) error {
	if GetVmState(name) == "NotFound" {
		fmt.Print("error: this vm does not exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Remove-VM -Name "+name+" -Force").Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func RenameVm(name string, newName string) error {
	if GetVmState(name) == "NotFound" {
		fmt.Print("error: this vm does not exist\n")
	} else if GetVmState(newName) != "NotFound" {
		fmt.Print("error: New vm name already exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VM -Name "+name+" -NewName "+newName).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
