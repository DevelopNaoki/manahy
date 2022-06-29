package process

import (
	"fmt"
	"regexp"
)

func CheckDiskParam(newDisk Disk) error {
	fileExist, err := isFileExist(newDisk.Path)
	if err != nil {
		return err
	} else if newDisk.Import && !fileExist {
		return fmt.Errorf("error: %s does not exist", newDisk.Path)
	} else {
		if fileExist {
			return fmt.Errorf("error: %s is already exist", newDisk.Path)
		}
		if newDisk.Type != "dynamic" && newDisk.Type != "fixed" && newDisk.Type != "differencing" {
			return fmt.Errorf("error: Undefined disk type")
		}
		fileExist, err = isFileExist(newDisk.ParentPath)
		if newDisk.Type == "differencing" && err != nil {
			return err
		} else if newDisk.Type == "differencing" && !fileExist {
			return fmt.Errorf("error: Parent disk doesnot exist")
		}
		diskSize := regexp.MustCompile("^[0-9]*[TGM]B$").FindString(newDisk.Size)
		if diskSize == "" {
			return fmt.Errorf("error: undefined size")
		}
	}
	return nil
}

func CheckSwitchParam(newSwitch Network) error {
	if GetSwitchType(newSwitch.Name) != "NotFound" {
		return fmt.Errorf("%s is already exist", newSwitch.Name)
	}
	if newSwitch.Type != "external" && newSwitch.Type != "internal" && newSwitch.Type != "private" {
		return fmt.Errorf("error: undefined switch type \n")
	}
	return nil
}

func CheckVmParam(newVm Vm) error {
	if GetVmState(newVm.Name) != "NotFound" {
		return fmt.Errorf("error: " + newVm.Name + " is already existed\n")
	}
	if newVm.Generation < 1 || newVm.Generation > 2 {
		return fmt.Errorf("error: Generation is not a valid value\n")
	}
	fileExist, err := isFileExist(newVm.Path)
	if err != nil {
		return err
	}
	if fileExist {
		return fmt.Errorf("error: " + newVm.Path + " is already exist\n")
	}
	fileExist, err = isFileExist(newVm.Image)
	if err != nil {
		return err
	}
	return nil
}

func CheckVmProcessorParam(cpu Cpu) error {
	if cpu.Thread < 0 {
		return fmt.Errorf("error: thread does not valid value\n")
	}
	if cpu.Reserve > 100 && cpu.Reserve < 0 {
		return fmt.Errorf("error: cpu reserve does not valid value\n")
	}
	if cpu.Maximum > 100 && cpu.Maximum < 0 {
		return fmt.Errorf("error: cpu maximum does not valid value\n")
	}
	if cpu.RelativeWeight > 10000 && cpu.RelativeWeight < 0 {
		return fmt.Errorf("error: cpu relative weight does not valid value\n")
	}
	return nil
}
