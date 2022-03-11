package process

import (
	"fmt"
	"os"
	"regexp"
)

func CheckDiskParam(newDisk Disk) {
	if isFileExist(newDisk.Path) {
		fmt.Print("error: Disk is already exist\n")
		os.Exit(1)
	}
	if !newDisk.Import {
		if newDisk.Type != "dynamic" && newDisk.Type != "fixed" && newDisk.Type != "differencing" {
			fmt.Print("error: Undefined DiskType\n")
			os.Exit(1)
		}
		if newDisk.Type == "differencing" && !isFileExist(newDisk.ParentPath) {
			fmt.Print("error: Disk doesnot exist\n")
			os.Exit(1)
		}
		diskSize := regexp.MustCompile("^[0-9]*[TGM]B$").FindString(newDisk.Size)
		if diskSize == "" {
			fmt.Print("error: undefined size\n")
			os.Exit(1)
		}
	}
}

func CheckSwitchParam(newSwitch Network) {
	if GetSwitchType(newSwitch.Name) != "NotFound" {
		fmt.Print("error: " + newSwitch.Name + " is already exist\n")
		os.Exit(1)
	}
	if newSwitch.Type != "external" && newSwitch.Type != "internal" && newSwitch.Type != "private" {
		fmt.Print("error: undefined switch type \n")
		os.Exit(1)
	}
}

func CheckVmParam(newVm Vm) {
	if GetVmState(newVm.Name) != "NotFound" {
		fmt.Print("error: " + newVm.Name + " is already existed\n")
		os.Exit(1)
	} else if newVm.Generation < 1 || newVm.Generation > 2 {
		fmt.Print("error: Generation is not a valid value\n")
		os.Exit(1)
	} else if isFileExist(newVm.Path) {
		fmt.Print("error: " + newVm.Path + " is already exist\n")
		os.Exit(1)
	} else if newVm.Image != "" && !isFileExist(newVm.Image) {
		fmt.Print("error: " + newVm.Image + " does not exist\n")
		os.Exit(1)
	}
	return
}

func CheckVmProcessorParam(cpu Cpu) {
	if cpu.Thread < 0 {
	}
	if cpu.Reserve > 100 && cpu.Reserve < 0 {
	}
	if cpu.Maximum > 100 && cpu.Maximum < 0 {
	}
	if cpu.RelativeWeight > 10000 && cpu.RelativeWeight < 0 {
	}
}
