package modules

import (
	"strconv"
)

func BuildByStruct(summarize Summarize) error {
	for _, disk := range summarize.Disks {
		err := CreateDisk(disk)
		if err != nil {
			return err
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
