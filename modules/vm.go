package modules

import (
	"fmt"
	"os/exec"
	"strconv"
)

// GetVmList get a list of VMs
func GetVmList() (vmList VmList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VM | Sort-Object State | Format-Table Name, State").Output()
	if err != nil {
		return vmList, err
	}

	vmList, err = vmListingOfExecuteResults(res)
	if err != nil {
		return vmList, err
	}

	return vmList, nil
}

// GetVmState get a VM state
func GetVmState(name string) (state string) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VM '"+name+"' | Format-Table State").Output()
	if err != nil {
		return "NotFound"
	} else {
		vmState := listingOfExecuteResults(res, "State")
		if len(vmState) == 1 {
			return vmState[0]
		}
	}
	return "Unknown"
}

func SetVmProcessor(name string, cpu Cpu) error {
	var cmd string

	if GetVmState(name) != "NotFound" {
		cmd = "Set-VMProcessor " + name + " "
		if cpu.Thread > 0 {
			cmd += "-Count " + strconv.Itoa(cpu.Thread) + " "
		}
		if cpu.Reserve <= 100 && cpu.Reserve >= 0 {
			cmd += " -Reserve " + strconv.Itoa(cpu.Reserve) + " "
		}
		if cpu.Maximum <= 100 && cpu.Maximum >= 0 {
			cmd += " -Maximum " + strconv.Itoa(cpu.Maximum) + " "
		}
		if cpu.RelativeWeight <= 10000 && cpu.RelativeWeight > 0 {
			cmd += " -RelativeWeight " + strconv.Itoa(cpu.RelativeWeight) + " "
		}
		cmd += " -ExposeVirtualizationExtensions " + strconv.FormatBool(cpu.Nested)

		err := exec.Command("powershell", "-NoProfile", cmd).Run()

		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not found")
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

	cmd := "New-VM -Name " + newVm.Name + " -Generation " + strconv.Itoa(newVm.Generation) + " -Path " + newVm.Path + " -MemoryStartupBytes " + newVm.Memory.Size
	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}
	SetVmProcessor(newVm.Name, newVm.Cpu)

	return nil
}

func RemoveVm(name string) error {
	if GetVmState(name) == "NotFound" {
		return fmt.Errorf("error: %s is not exist", name)
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
		return fmt.Errorf("error: %s is not exist", name)
	} else if GetVmState(newName) != "NotFound" {
		return fmt.Errorf("error: %s is already exist", newName)
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VM -Name "+name+" -NewName "+newName).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// ConnectVm connect the VM
func ConnectVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "vmconnect localhost '"+name+"'").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not running", name)
	}
	return nil
}

// StartVm start the VM
func StartVm(name string) error {
	if GetVmState(name) != "Running" {
		err := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is already running", name)
	}
	return nil
}

// StopVm stop the VM
func StopVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not running", name)
	}
	return nil
}

// DestroyVm force stop VM
func DestroyVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Stop-VM -Force -Name '"+name+"'").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not running", name)
	}
	return nil
}

// SaveVm save VM
func SaveVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not running", name)
	}
	return nil
}

func SuspendVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Suspend-VM -Name '"+name+"'").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not running", name)
	}
	return nil
}

func RestartVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Restart-VM -Name '"+name+"' -Force").Run()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s is not running", name)
	}
	return nil
}
