package process

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// GetSwitchList get a list of Switch
func GetSwitchList(state string) (list []string) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VMSwitch * | where {$_.SwitchType -eq '"+state+"'} | Format-Table Name").Output()
	if err != nil {
		panic(err)
	}
	list = listingOfExecuteResults(res, "Name")
	return
}

// GetSwitchType get a Switch Type
func GetSwitchType(name string) (state string) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VMSwitch '"+name+"' | Format-Table SwitchType").Output()
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
		err := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -SwitchType "+switchType).Run()
		if err != nil {
			panic(err)
		}
}

func ChangeSwitchNetAdapter(name string, netAdapter string) {
	err := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -NetAdapterName '"+netAdapter+"'").Run()
	if err != nil {
		panic(err)
	}
}

func CreateSwitch(newSwitch Network) {
	var args string
	CheckSwitchParam(newSwitch)

	if newSwitch.Type == "external" {
		args = "New-VMSwitch -name '" + newSwitch.Name + "' -NetAdapterName '" + newSwitch.ExternameInterface + "' -AllowManagementOS $" + strconv.FormatBool(newSwitch.AllowManagementOs)
	} else {
		args = "New-VMSwitch -name '" + newSwitch.Name + "' -SwitchType " + newSwitch.Type
	}

	err := exec.Command("powershell", "-NoProfile", args).Run()
	if err != nil {
		panic(err)
	}
}

func RemoveSwitch(name string) {
	if GetSwitchType(name) == "NotFound" {
		fmt.Print("error: " + name + " is not exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Remove-VMSwitch '"+name+"' -Force").Run()
		if err != nil {
			panic(err)
		}
	}
}

func RenameSwitch(name string, newName string) {
	if GetSwitchType(name) == "NotFound" {
		fmt.Print("error: " + name + " is not exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VMSwitch '"+name+"' -NewName "+newName).Run()
		if err != nil {
			panic(err)
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
	} /*
		if newSwitch.Type == "external" {
			err := exec.Command("powershell", "-NoProfile", "Get-NetAdapter '"+newSwitch.ExternameInterface+"'").Run()
			if err != nil {
				fmt.Print("error: "+newSwitch.ExternameInterface+" undefined interface\n")
				os.Exit(1)
			}
		}*/
}
