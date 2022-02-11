package process

import (
	"fmt"
	"os/exec"
	"strconv"
)

// GetVmList get a list of VMs
func GetVmList(state string) (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VM | where {$_.State -eq '"+state+"'} | Format-Table Name")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	list = listingOfExecuteResults(res, "Name")
	return
}

// GetVmState get a VM state
func GetVmState(name string) (state string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VM '"+name+"' | Format-Table State")
	res, err := cmd.Output()
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

func SetVmMemory(name string, memory Memory) {

}

func CreateVm(newVm Vm) {
	if CheckVmOption(newVm) {
		args := "New-VM -Name " + newVm.Name + " -Generation " + strconv.Itoa(newVm.Generation) + " -Path " + newVm.Path + " -MemoryStartupBytes " + newVm.Memory.Size
		cmd := exec.Command("powershell", "-NoProfile", args)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		SetVmProcessor(newVm.Name, newVm.Cpu)
	}
}

func CheckVmOption(newVm Vm) (passCheck bool) {
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
// StartVm start the VM
func StartVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// StopVm stop the VM
func StopVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// DestroyVm force stop VM
func DestroyVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Stop-VM -Force -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// SaveVm save VM
func SaveVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
