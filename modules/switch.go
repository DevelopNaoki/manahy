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

// ChangeSwitchType change switch type
func ChangeSwitchType(name string, switchType string) error {
	nameType := GetSwitchType(name)
	switch nameType {
	case "NotFound":
		return fmt.Errorf("%s is not exist", name)
	case "Unknown":
		return fmt.Errorf("Unknown error")
	default:
		err := checkSwitchTypeParam(switchType)
		if err != nil {
			return err
		} else if nameType == switchType {
			return fmt.Errorf("%s's now type already %s", name, switchType)
		}

		err = exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -SwitchType "+switchType).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// ChangeSwitchNetAdapter change netpadapter of external switch
func ChangeSwitchNetAdapter(name string, netAdapter string) error {
	nameType := GetSwitchType(name)
	switch nameType {
	case "NotFound":
		return fmt.Errorf("%s is not exist", name)
	case "Unknown":
		return fmt.Errorf("Unknown error")
	default:
		err := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '"+name+"' -NetAdapterName '"+netAdapter+"'").Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateSwitch create new switch
func CreateSwitch(newSwitch Network) error {
	var cmd string
	err := checkSwitchParam(newSwitch)
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

// RemoveSwitch remove switch
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

// RenameSwitch rename switch
func RenameSwitch(name string, newName string) error {
	if GetSwitchType(name) == "NotFound" {
		return fmt.Errorf("error: %s is not exist", name)
	} else if GetSwitchType(newName) != "NotFound" {
		return fmt.Errorf("error: %s is already exist", newName)
	} else {
		err := exec.Command("powershell", "-NoProfile", "Rename-VMSwitch '"+name+"' -NewName "+newName).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func checkSwitchParam(newSwitch Network) error {
	if GetSwitchType(newSwitch.Name) != "NotFound" {
		return fmt.Errorf("%s is already exist", newSwitch.Name)
	}

	err := checkSwitchTypeParam(newSwitch.Type)
	if err != nil {
		return err
	}

	err = checkSwitchParamIntegrity(newSwitch.Type, newSwitch.ExternameInterface)
	if err != nil {
		return err
	}
	return nil
}

func checkSwitchTypeParam(switchType string) error {
	switch switchType {
	case "external":
	case "internal":
	case "private":
	default:
		return fmt.Errorf("error: undefined switch type \n")
	}
	return nil
}

func checkSwitchParamIntegrity(switchType string, externameInterface string) error {
	if switchType == "external" && externameInterface == "" {
		return fmt.Errorf("Need externameInterface option")
	} else if switchType != "external" && externameInterface != "" {
		return fmt.Errorf("Don't need externameInterface option")
	}
	return nil
}
