package process

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// GetSwitchList get a list of Switch
func GetSwitchLists() (switchList SwitchList) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VMSwitch | Sort-Object SwitchType | Format-Table Name, SwitchType").Output()
	if err != nil {
		panic(err)
	}
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Name") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			switchType := regexp.MustCompile("External$|Internal$|Private$").FindString(split[i])
			split[i] = regexp.MustCompile("External$|Internal$|Private$").ReplaceAllString(split[i], "")
			split[i] = strings.TrimSpace(split[i])

			switch switchType {
			case "External":
				switchList.External = append(switchList.External, split[i])
			case "Internal":
				switchList.Internal = append(switchList.Internal, split[i])
			case "Private":
				switchList.Private = append(switchList.Private, split[i])
			}
		}
	}
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
