package process

import (
	"fmt"
	"os"
	"regexp"
)

func CheckDiskParam(newDisk Disk) error {
	fileExist, err := isFileExist(newDisk.Path)
	if err != nil {
		return err
	} else if fileExist {
		return fmt.Errorf("error: Disk is already exist")
	}
	if !newDisk.Import {
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
		return fmt.Errorf("error: %s is already exist", newSwitch.Name)
	}
	if newSwitch.Type != "external" && newSwitch.Type != "internal" && newSwitch.Type != "private" {
		return fmt.Errorf("error: undefined switch type \n")
	}
	return nil
}

func CheckVmParam(newVm Vm) {
	if GetVmState(newVm.Name) != "NotFound" {
		fmt.Print("error: " + newVm.Name + " is already existed\n")
		os.Exit(1)
	}
	if newVm.Generation < 1 || newVm.Generation > 2 {
		fmt.Print("error: Generation is not a valid value\n")
		os.Exit(1)
	}
	fileExist, err := isFileExist(newVm.Path)
	if err != nil {
		os.Exit(1)
	}
	if fileExist {
		fmt.Print("error: " + newVm.Path + " is already exist\n")
		os.Exit(1)
	}
	fileExist, err = isFileExist(newVm.Image)
	if err != nil {
		os.Exit(1)
	} else if newVm.Image != "" && !fileExist {
		fmt.Print("error: " + newVm.Image + " does not exist\n")
		os.Exit(1)
	}
	return
}

func CheckVmProcessorParam(cpu Cpu) {
	if cpu.Thread < 0 {
		fmt.Print("error: thread does not valid value\n")
		os.Exit(1)
	}
	if cpu.Reserve > 100 && cpu.Reserve < 0 {
		fmt.Print("error: cpu reserve does not valid value\n")
		os.Exit(1)
	}
	if cpu.Maximum > 100 && cpu.Maximum < 0 {
		fmt.Print("error: cpu maximum does not valid value\n")
		os.Exit(1)
	}
	if cpu.RelativeWeight > 10000 && cpu.RelativeWeight < 0 {
		fmt.Print("error: cpu relative weight does not valid value\n")
		os.Exit(1)
	}
}
