package process

import (
	"fmt"
	"os/exec"
	"strconv"
)

// GetSwitchList get a list of Switch
func GetSwitchList(state string) (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VMSwitch * | where {$_.SwitchType -eq '"+state+"'} | Format-Table Name")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	list = listingOfExecuteResults(res, "Name")
	return
}

// GetSwitchType get a Switch Type
func GetSwitchType(name string) (state string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VMSwitch '"+name+"' | Format-Table SwitchType")
	res, err := cmd.Output()
	if err != nil {
		state = "NotFound"
	} else {
		list := listingOfExecuteResults(res, "SwitchType")
		if len(list) == 1 {
			state = list[0]
		}
	}
	return
}

func ChangeSwitchType(name string, switchType string) {
	cmd := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -SwitchType "+switchType)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func CreateSwitch(newSwitch Network) {
	var args string

	if GetSwitchType(newSwitch.Name) != "NotFound" {
		fmt.Print("error: " + newSwitch.Name + " is already exist\n")
	} else {
		if newSwitch.Type == "external" {
			err := exec.Command("powershell", "-NoProfile", "Get-NetAdapter "+newSwitch.ExternameInterface)
			if err != nil {
				fmt.Print("error: " + newSwitch.ExternameInterface + " is not found\n")
			} else {
				args = "New-VMSwitch -name '" + newSwitch.Name + "' -NetAdapterName '" + newSwitch.ExternameInterface + "' -AllowManagementOS $" + strconv.FormatBool(newSwitch.AllowManagementOs)
			}
		} else if newSwitch.Type == "internal" || newSwitch.Type == "private" {
			args = "New-VMSwitch -name '" + newSwitch.Name + "' -SwitchType " + newSwitch.Type
		}

		cmd := exec.Command("powershell", "-NoProfile", args)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func RemoveSwitch(name string) {
	if GetSwitchType(name) == "NotFound" {
		fmt.Print("error: " + name + " is not exist\n")
	} else {
		cmd := exec.Command("powershell", "-NoProfile", "Remove-VMSwitch '"+name+"' -Force")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
