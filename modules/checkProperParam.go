package modules

import (
	"fmt"
)

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
	} else if !imagePathExist {
		return fmt.Errorf("error: " + newVm.Image + " is doesnt exist\n")
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
