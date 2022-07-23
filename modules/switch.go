package modules

import (
	"fmt"
	"os/exec"
	"strconv"
)

// GetSwitchList get a list of Switch
func GetSwitchList() (switchList SwitchList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-VMSwitch | Sort-Object SwitchType | Format-Table Name, SwitchType").Output()
	if err != nil {
		return switchList, fmt.Errorf("failed get vm switch list")
	}

	switchList, err = switchListingOfExecuteResults(res)
	if err != nil {
		return switchList, err
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
		} else {
			state = "Unknown"
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
	var cmd string
	err := CheckSwitchParam(newSwitch)
	if err != nil {
		return err
	}

	if newSwitch.Type == "external" {
		cmd = "New-VMSwitch -name '" + newSwitch.Name + "' -NetAdapterName '" + newSwitch.ExternameInterface + "' -AllowManagementOS $" + strconv.FormatBool(newSwitch.AllowManagementOs)
	} else {
		cmd = "New-VMSwitch -name '" + newSwitch.Name + "' -SwitchType " + newSwitch.Type
	}

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return fmt.Errorf("failed create new switch")
	}
	return nil
}

func RemoveSwitch(name string) error {
	if GetSwitchType(name) == "NotFound" {
		return fmt.Errorf("error: %s does not exist", name)
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
		return fmt.Errorf("error: %s does not exist", name)
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VMSwitch '"+name+"' -NewName "+newName).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
