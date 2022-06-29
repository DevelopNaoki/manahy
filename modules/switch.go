package modules

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// GetSwitchList get a list of Switch
func GetSwitchList() (switchList SwitchList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VMSwitch | Sort-Object SwitchType | Format-Table Name, SwitchType").Output()
	if err != nil {
		return switchList, fmt.Errorf("failed get vm switch list")
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
			default:
				return switchList, fmt.Errorf("Unknown error for switch list")
			}
		}
	}
	return switchList, nil
}

// GetSwitchType get a Switch Type
func GetSwitchType(name string) (state string) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VMSwitch '"+name+"' | Format-Table SwitchType").Output()
	if err != nil {
		state = "NotFound"
	} else {
		switchType := listingOfExecuteResults(res, "SwitchType")
		if len(switchType) == 1 {
			state = switchType[0]
		}
	}
	return state
}

func ChangeSwitchType(name string, switchType string) error {
	err := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -SwitchType "+switchType).Run()
	if err != nil {
		return err
	}
	return nil
}

func ChangeSwitchNetAdapter(name string, netAdapter string) error {
	err := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -NetAdapterName '"+netAdapter+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

func CreateSwitch(newSwitch Network) error {
	var args string
	err := CheckSwitchParam(newSwitch)
	if err != nil {
		return err
	}

	if newSwitch.Type == "external" {
		args = "New-VMSwitch -name '" + newSwitch.Name + "' -NetAdapterName '" + newSwitch.ExternameInterface + "' -AllowManagementOS $" + strconv.FormatBool(newSwitch.AllowManagementOs)
	} else {
		args = "New-VMSwitch -name '" + newSwitch.Name + "' -SwitchType " + newSwitch.Type
	}

	err = exec.Command("powershell", "-NoProfile", args).Run()
	if err != nil {
		return fmt.Errorf("failed create new switch")
	}
	return nil
}

func RemoveSwitch(name string) error {
	if GetSwitchType(name) == "NotFound" {
		fmt.Print("error: " + name + " does not exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Remove-VMSwitch '"+name+"' -Force").Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func RenameSwitch(name string, newName string) error {
	if GetSwitchType(name) == "NotFound" {
		fmt.Print("error: " + name + " does not exist\n")
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VMSwitch '"+name+"' -NewName "+newName).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
