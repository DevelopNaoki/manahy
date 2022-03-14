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
	res, e := exec.Command("powershell", "-NoProfile", "Get-VM | Sort-Object State | Format-Table Name, State").Output()
	if e != nil {
		return vmList, e
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
		list := listingOfExecuteResults(res, "State")
		if len(list) == 1 {
			state = list[0]
		}
	}
	return
}

func SetVmProcessor(name string, cpu Cpu) {
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
		panic(err)
	}
}

func SetVmMemory(name string) {
}

func CreateVm(newVm Vm) {
	CheckVmParam(newVm)

	args := "New-VM -Name " + newVm.Name + " -Generation " + strconv.Itoa(newVm.Generation) + " -Path " + newVm.Path + " -MemoryStartupBytes " + newVm.Memory.Size
	cmd := exec.Command("powershell", "-NoProfile", args)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	SetVmProcessor(newVm.Name, newVm.Cpu)
}

func RemoveVm(name string) {
	if GetVmState(name) == "NotFound" {
		fmt.Print("error: this vm does not exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Remove-VM -Name "+name+" -Force").Run()
		if err != nil {
			panic(err)
		}
	}
}

func RenameVm(name string, newName string) {
	if GetVmState(name) == "NotFound" {
		fmt.Print("error: this vm does not exist\n")
	} else if GetVmState(newName) != "NotFound" {
		fmt.Print("error: New vm name already exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VM -Name "+name+" -NewName "+newName).Run()
		if err != nil {
			panic(err)
		}
	}
}

// -------------------- VM Operation --------------------
// ConnectVm connect VM
func ConnectVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "vmconnect localhost "+name).Run()
	if err != nil {
		panic(err)
	}
}

// StartVm start the VM
func StartVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'").Run()
	if err != nil {
		panic(err)
	}
}

// StopVm stop the VM
func StopVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'").Run()
	if err != nil {
		panic(err)
	}
}

// DestroyVm force stop VM
func DestroyVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "Stop-VM -Force -Name '"+name+"'").Run()
	if err != nil {
		panic(err)
	}
}

// SaveVm save VM
func SaveVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'").Run()
	if err != nil {
		panic(err)
	}
}

func SuspendVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "Suspend-VM -Name '"+name+"'").Run()
	if err != nil {
		panic(err)
	}
}

func RestartVm(name string) {
	err := exec.Command("powershell", "-NoProfile", "Restart-VM -Name '"+name+"' -Force").Run()
	if err != nil {
		panic(err)
	}
}
