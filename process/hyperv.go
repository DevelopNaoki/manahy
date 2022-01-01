package process

import (
	"fmt"
)

func VmList(option string) {
	if option == "running" || option == "all" {
		fmt.Print("Running VM's\n")
		activeVms := GetVMList("Running")
		for _, l := range activeVms {
			fmt.Printf("- %s\n", l)
		}
		fmt.Print("\n")
	}
	if option == "save" || option == "all" {
		fmt.Print("Saved VM's\n")
		savedVms := GetVMList("Saved")
		for _, l := range savedVms {
			fmt.Printf("- %s\n", l)
		}
		fmt.Print("\n")
	}
	if option == "off" || option == "all" {
		fmt.Print("Stop VM's\n")
		offVms := GetVMList("Off")
		for _, l := range offVms {
			fmt.Printf("- %s\n", l)
		}
		fmt.Print("\n")
	}
}

func VmState(name string) {
	fmt.Print(GetVmState(name) + "\n")
}

func VmOperation(name string, operation string) {
	state := GetVmState(name)
	switch operation {
	case "start":
		if state != "Running" {
			StartVm(name)
		}
        case "saved":
                if state == "Running" {
                        SaveVm(name)
                }
        }
	case "shutdown":
		if state == "Running" {
			StopVm(name)
		}
	case "destroy":
		if state == "Running" {
			DestroyVm(name)
		}
	}
}
