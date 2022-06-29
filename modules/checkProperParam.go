package modules

import (
	"fmt"
	"regexp"
)

func CheckDiskParam(newDisk Disk) error {
	diskExist, err := isFileExist(newDisk.Path)
	if err != nil {
		return err
	} else if !diskExist && newDisk.Import {
		return fmt.Errorf("error: %s does not exist", newDisk.Path)
	} else if diskExist && !newDisk.Import {
		return fmt.Errorf("error: %s is already exist", newDisk.Path)
	} else {
		switch newDsik.Type {
		case "dynamic":
		case "fixed":
		case "differencing":
			parentDiskExist, err := isFileExist(newDisk.ParentPath)
			if err != nil {
				return err
			} else if !parentDiskExist {
				return fmt.Errorf("error: Parent disk doesnot exist")
			}
		default:
			return fmt.Errorf("error: Undefined disk type")
		}
		
		diskSize := regexp.MustCompile("^[0-9]*[TGM]B$").FindString(newDisk.Size)
		if diskSize == "" {
			return fmt.Errorf("error: unknown size")
		}
	}
	return nil
}

func CheckSwitchParam(newSwitch Network) error {
	if GetSwitchType(newSwitch.Name) != "NotFound" {
		return fmt.Errorf("%s is already exist", newSwitch.Name)
	}
	switch newSwitch.Type {
	case "external":
	case "internal":
	case "private":
	default:
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
	vmPathExist, err := isFileExist(newVm.Path)
	if err != nil {
		return err
	}
	if vmPathExist {
		return fmt.Errorf("error: " + newVm.Path + " is already exist\n")
	}
	imagePathExist, err := isFileExist(newVm.Image)
	if err != nil {
		return err
	} if !imagePathExist {
		return fmt.Errorf("error: "+newVm.Image+" is doesnt exist\n")	
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
