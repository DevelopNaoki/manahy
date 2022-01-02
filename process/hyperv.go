package process

import (
	"os/exec"
)

func GetVMList(state string) (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VM | where {$_.State -eq '"+state+"'} | Select-Object Name")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	list = listingOfExecuteResults(res, "Name")
	return
}

func GetVmState(name string) (state string) {
	args := "Get-VM '" + name + "' | Select-Object State"
	cmd := exec.Command("powershell", "-NoProfile", args)
	res, err := cmd.Output()
	if err != nil {
		state = "NotFound"
	} else {
		list := listingOfExecuteResults(res, "State")
		if len(list) == 1 {
			state = list[0]
		}
	}
	return
}

func StartVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func StopVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func DestroyVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"' -Force")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func SaveVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
