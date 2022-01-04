package process

import (
	"os/exec"
)

// GetVmList get a list of VMs
func GetVmList(state string) (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VM | where {$_.State -eq '"+state+"'} | Select-Object Name")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	list = listingOfExecuteResults(res, "Name")
	return
}

// GetVmState get a VM state
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

// StartVm start the VM
func StartVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// StopVm stop the VM
func StopVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// DestroyVm force stop VM
func DestroyVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"' -Force")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// SaveVm save VM
func SaveVm(name string) {
	cmd := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
