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
	if GetVmState(name) == "NotFound" {
		return fmt.Errorf("%s is not exist", name)
	}
	err := checkVmProcessorParam(cpu)
	if err != nil {
		return err
	}

	cmd := "Set-VMProcessor " + name
	cmd += " -Count " + strconv.Itoa(cpu.Thread)
	/*	cmd += " -Reserve "+strconv.Itoa(cpu.Reserve)
		cmd += " -Maximum "+strconv.Itoa(cpu.Maximum)
		cmd += " -RelativeWeight "+strconv.Itoa(cpu.RelativeWeight)
	*/cmd += " -ExposeVirtualizationExtensions $" + strconv.FormatBool(cpu.Nested)

	fmt.Printf("%s\n", cmd)

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}

func SetVmMemory(name string, memory Memory) error {
	if GetVmState(name) == "NotFound" {
		return fmt.Errorf("%s is not exist", name)
	}
	err := checkVmMemoryParam(memory)
	if err != nil {
		return err
	}

	cmd := "Set-VMMemory -VMName " + name
	cmd += " -StartupBytes " + memory.Size
	cmd += " -DynamicMemoryEnabled $" + strconv.FormatBool(memory.Dynamic)

	fmt.Printf("%s\n", cmd)

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}

func SetVmHardDisk(name string, disks []string) error {
	if GetVmState(name) == "NotFound" {
		return fmt.Errorf("%s is not exist", name)
	}

	for _, disk := range disks {
		diskExist, err := isFileExist(disk)
		if err != nil {
			return err
		} else if !diskExist {
			return fmt.Errorf("%s is not exist", disk)
		}

		cmd := "Add-VMHardDiskDrive -VMName " + name
		cmd += " -Path " + disk

		fmt.Printf("%s\n", cmd)

		err = exec.Command("powershell", "-NoProfile", cmd).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func SetVmImageFile(name string, image string) error {
	if GetVmState(name) == "NotFound" {
		return fmt.Errorf("%s is not exist", name)
	}

	fileExist, err := isFileExist(image)
	if err != nil {
		return err
	} else if !fileExist {
		return fmt.Errorf("%s is not exist", image)
	}

	cmd := "Add-VMDvdDrive -VMName " + name
	cmd += " -Path " + image

	fmt.Printf("%s\n", cmd)

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}

func SetVmSwitch(name string, networks []string) error {
	for _, network := range networks {
		switchExist := GetSwitchType(network)
		if switchExist == "NotFound" || switchExist == "Unknown" {
			return fmt.Errorf("%s is not exist", network)
		} else if switchExist == "Unknown" {
			return fmt.Errorf("%s is Unknown error", network)
		}

		cmd := "Add-VMNetworkAdapter -VMName " + name
		cmd += " -SwitchName " + network

		fmt.Printf("%s\n", cmd)

		err := exec.Command("powershell", "-NoProfile", cmd).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateVm(newVm Vm) error {
	err := checkVmParam(newVm)
	if err != nil {
		return err
	}

	cmd := "New-VM -Name " + newVm.Name
	cmd += " -Generation " + strconv.Itoa(newVm.Generation)
	cmd += " -Path " + newVm.Path

	fmt.Printf("%s\n", cmd)

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	for true {
		state := GetVmState(newVm.Name)
		if state != "NotFound" && state != "Unknown" {
			break
		}
	}

	err = SetVmProcessor(newVm.Name, newVm.Cpu)
	if err != nil {
		return err
	}
	err = SetVmMemory(newVm.Name, newVm.Memory)
	if err != nil {
		return err
	}
	err = SetVmHardDisk(newVm.Name, newVm.Disks)
	if err != nil {
		return err
	}
	err = SetVmImageFile(newVm.Name, newVm.Image)
	if err != nil {
		return err
	}
	err = SetVmSwitch(newVm.Name, newVm.Networks)
	if err != nil {
		return err
	}

	return nil
}

func RemoveVm(name string) error {
	if GetVmState(name) == "NotFound" {
		return fmt.Errorf("%s is not exist", name)
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
		return fmt.Errorf("%s is not exist", name)
	} else if GetVmState(newName) != "NotFound" {
		return fmt.Errorf("%s is already exist", newName)
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

// SuspendVm suspend vm
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

// RestartVM restart vm
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

func checkVmParam(newVm Vm) error {
	if GetVmState(newVm.Name) != "NotFound" {
		return fmt.Errorf("error: " + newVm.Name + " is already existed\n")
	}

	if newVm.Generation < 1 || newVm.Generation > 2 {
		return fmt.Errorf("error: Generation is not a valid value\n")
	}

	vmPathExist, err := isFileExist(newVm.Path)
	if err != nil {
		return err
	}
	if vmPathExist {
		return fmt.Errorf("error: " + newVm.Path + " is already exist\n")
	}

	if newVm.Image != "" {
		imageExist, err := isFileExist(newVm.Image)
		if err != nil {
			return err
		} else if !imageExist {
			return fmt.Errorf("error: " + newVm.Image + " is doesnt exist\n")
		}
	}

	return nil
}

func checkVmProcessorParam(cpu Cpu) error {
	if cpu.Thread < 0 {
		return fmt.Errorf("error: thread does not valid value\n")
	}
	/*	if cpu.Reserve > 100 && cpu.Reserve < 0 {
			return fmt.Errorf("error: cpu reserve does not valid value\n")
		}
		if cpu.Maximum > 100 && cpu.Maximum < 0 {
			return fmt.Errorf("error: cpu maximum does not valid value\n")
		}
		if cpu.RelativeWeight > 10000 && cpu.RelativeWeight < 0 {
			return fmt.Errorf("error: cpu relative weight does not valid value\n")
		}
	*/return nil
}

func checkVmMemoryParam(memory Memory) error {
	return nil
}
