package process

import (
	"fmt"
	"os/exec"
	"strconv"
)

// GetVmList get a list of VMs
func GetVmList(state string) (list []string) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VM | where {$_.State -eq '"+state+"'} | Format-Table Name").Output()
	if err != nil {
		panic(err)
	}
	list = listingOfExecuteResults(res, "Name")
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
	if CheckVmParam(newVm) {
		args := "New-VM -Name " + newVm.Name + " -Generation " + strconv.Itoa(newVm.Generation) + " -Path " + newVm.Path + " -MemoryStartupBytes " + newVm.Memory.Size
		cmd := exec.Command("powershell", "-NoProfile", args)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		SetVmProcessor(newVm.Name, newVm.Cpu)
	}
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

func CheckVmParam(newVm Vm) (passCheck bool) {
	passCheck = true

	if GetVmState(newVm.Name) != "NotFound" {
		fmt.Print("error: " + newVm.Name + " is already existed\n")
		passCheck = false
	} else if newVm.Generation < 1 || newVm.Generation > 2 {
		fmt.Print("error: Generation is not a valid value\n")
		passCheck = false
	} else if isFolderExist(newVm.Path) {
		fmt.Print("error: " + newVm.Path + " is already exist\n")
		passCheck = false
	} else if newVm.Image != "" && !isFileExist(newVm.Image) {
		fmt.Print("error: " + newVm.Image + " does not exist\n")
		passCheck = false
	} else {
		for _, disk := range newVm.Disk {
			if !isFileExist(disk) {
				fmt.Print("error: " + disk + " does not exist\n")
				passCheck = false
			}
		}
		for _, network := range newVm.Network {
			if GetSwitchType(network) == "NotFound" {
				fmt.Print("error: " + network + " does not exist\n")
				passCheck = false
			}
		}
	}
	return
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
