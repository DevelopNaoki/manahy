package process

import (
	"os/exec"
)

// GetSwitchList get a list of Switch
func GetSwitchList(state string) (list []string){
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
        cmd := exec.Command("powershell", "-NoProfile", "Get-VMSwitch '" +name+ "' | Format-Table SwitchType")
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
	cmd := exec.Command("powershell", "-NoProfile", "Set-VMSwitch '" +name+ "' -SwitchType "+switchType+"")
        err := cmd.Run()
        if err != nil {
		panic(err)
	}
        return
}
