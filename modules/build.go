package modules

import (
	"fmt"
	"strconv"
)

func BuildByStruct(summarize Summarize) error {
	for _, disk := range summarize.Disks {
		state, e := isFileExist(disk.Path)
		if e != nil {
			return e
		}
		if !state && !disk.Import {
			err := CreateDisk(disk)
			if err != nil {
				return err
			}
		} else if !state && disk.Import {
			return fmt.Errorf("%s is not exist", disk.Path)
		}
	}

	for _, network := range summarize.Networks {
		state := GetSwitchType(network.Name)
		if state == "NotFound" {
			err := CreateSwitch(network)
			if err != nil {
				return err
			}
		}
	}

	for _, vm := range summarize.Vms {
		if vm.Count == 0 {
			vm.Count = 1
		}
		for i := 1; i <= vm.Count; i++ {
			if vm.Count != 1 {
				vm.Name += strconv.Itoa(i)
			}
			err := CreateVm(vm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
