// hyperv package is manage Hyper-V
package hyperv

import (
	"fmt"
	"os/exec"
	"strconv"
)

// ------ //
// Get VM
// ------ //

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

// IsVmExist
func IsVmExist(name string) error {
	state := GetVmState(name)
	if state == "Unknown" {
		return fmt.Errorf("failed get vm state")
	} else if state == "NotFound" {
		return fmt.Errorf("%s is not found", name)
	}
	return nil
}

// IsNotVmExist
func IsNotVmExist(name string) error {
	state := GetVmState(name)
	if state == "Unknown" {
		return fmt.Errorf("failed get vm state")
	} else if state != "NotFound" {
		return fmt.Errorf("%s is already found", name)
	}
	return nil
}

// -------------- //
//  Set VM Option
// -------------- //

// SetVmProcessor is set vm processor
func SetVmProcessor(name string, cpu Cpu) error {
	err := IsVmExist(name)
	if err != nil {
		return err
	}
	err = checkVmProcessorParam(cpu)
	if err != nil {
		return err
	}

	cmd := "Set-VMProcessor '" + name + "'"
	cmd += " -Count " + strconv.Itoa(cpu.Thread)
	cmd += " -ExposeVirtualizationExtensions $" + strconv.FormatBool(cpu.Nested)

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}

func SetVmMemory(name string, memory Memory) error {
	err := IsVmExist(name)
	if err != nil {
		return err
	}
	err = checkVmMemoryParam(memory)
	if err != nil {
		return err
	}

	cmd := "Set-VMMemory -VMName '" + name + "'"
	cmd += " -StartupBytes " + memory.Size
	cmd += " -DynamicMemoryEnabled $" + strconv.FormatBool(memory.Dynamic)

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}

func SetVmHardDisk(name string, disks []string) error {
	err := IsVmExist(name)
	if err != nil {
		return err
	}

	for index := range disks {
		err := isFileExist(disks[index])
		if err != nil {
			return err
		}

		cmd := "Add-VMHardDiskDrive -VMName '" + name + "'"
		cmd += " -Path " + disks[index]

		fmt.Printf("%s\n", cmd)

		err = exec.Command("powershell", "-NoProfile", cmd).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func SetVmImageFile(name string, image string) error {
	err := IsVmExist(name)
	if err != nil {
		return err
	}
	err = isFileExist(image)
	if err != nil {
		return err
	}

	cmd := "Add-VMDvdDrive -VMName '" + name + "'"
	cmd += " -Path " + image

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

// ---------------- //
// Create/Remove VM
//  ---------------- //

func CreateVm(newVm Vm, output bool) error {
	err := checkVmParam(newVm)
	if output {
		PrintError("Check Vm Param", err)
	}
	if err != nil {
		return err
	}

	cmd := "New-VM -Name " + newVm.Name
	cmd += " -Generation " + strconv.Itoa(newVm.Generation)
	cmd += " -Path " + newVm.Path

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if output {
		PrintError("Create Vm", err)
	}
	if err != nil {
		return err
	}

	err = SetVmProcessor(newVm.Name, newVm.Cpu)
	if output {
		PrintError("Set Processor", err)
	}
	if err != nil {
		return err
	}

	err = SetVmMemory(newVm.Name, newVm.Memory)
	if output {
		PrintError("Set Memory", err)
	}
	if err != nil {
		return err
	}

	err = SetVmHardDisk(newVm.Name, newVm.Disks)
	if output {
		PrintError("Set HardDisk", err)
	}
	if err != nil {
		return err
	}

	err = SetVmImageFile(newVm.Name, newVm.Image)
	if output {
		PrintError("Set Image File", err)
	}
	if err != nil {
		return err
	}

	err = SetVmSwitch(newVm.Name, newVm.Networks)
	if output {
		PrintError("Set VMSwitch", err)
	}
	if err != nil {
		return err
	}

	return nil
}

func RemoveVm(name string, output bool) error {
	err := IsVmExist(name)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "Remove-VM -Name '"+name+"' -Force").Run()
	if output {
		PrintError("Remove Vm", err)
	}
	if err != nil {
		return err
	}
	return nil
}

// ---------------- //
// Update Vm Option
// ---------------- //

func RenameVm(name string, newName string) error {
	err := IsVmExist(name)
	if err != nil {
		return err
	}
	err = IsNotVmExist(newName)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "Rename-VM -Name '"+name+"' -NewName '"+newName+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

// ------------ //
// VM Operation
// ------------ //

// ConnectVm connect the VM
func ConnectVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "vmconnect localhost '"+name+"'").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is not running", name)
}

// StartVm start the VM
func StartVm(name string) error {
	if GetVmState(name) != "Running" {
		err := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is already running", name)
}

// StopVm stop the VM
func StopVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is not running", name)
}

// DestroyVm force stop VM
func DestroyVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Stop-VM -Force -Name '"+name+"'").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is not running", name)
}

// SaveVm save VM
func SaveVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is not running", name)
}

// SuspendVm suspend vm
func SuspendVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Suspend-VM -Name '"+name+"'").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is not running", name)
}

// RestartVM restart vm
func RestartVm(name string) error {
	if GetVmState(name) == "Running" {
		err := exec.Command("powershell", "-NoProfile", "Restart-VM -Name '"+name+"' -Force").Run()
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("%s is not running", name)
}

// --------------- //
// Check VM Option
// --------------- //

func checkVmParam(newVm Vm) error {
	err := IsNotVmExist(newVm.Name)
	if err != nil {
		return err
	}

	err = checkVmGeneration(newVm.Generation)
	if err != nil {
		return err
	}

	err = checkVmPath(newVm.Name, newVm.Path)
	if err != nil {
		return err
	}

	if newVm.Image != "" {
		err := isFileExist(newVm.Image)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkVmGeneration(generation int) error {
	if generation < 1 || generation > 2 {
		return fmt.Errorf("generation is not a valid value\n")
	}
	return nil
}

func checkVmPath(name string, path string) error {
	err := isNotFileExist(path + "\\" + name)
	if err != nil {
		return err
	}
	return nil
}

func checkVmProcessorParam(cpu Cpu) error {
	if cpu.Thread < 0 {
		return fmt.Errorf("thread does not valid value\n")
	}
	return nil
}

func checkVmMemoryParam(memory Memory) error {
	return nil
}
