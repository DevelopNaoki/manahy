package process

import (
	"os/exec"
	"regexp"
	"strings"
	"strconv"
)

func IsFileExist(path string) (status bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path -Path "+path+" -PathType Leaf")
	res, err := cmd.Output()
        if err != nil {
                panic(err)
	}
	status, _ = strconv.ParseBool(string(res))
	return
}

func IsFolderExist(path string) (status bool) {
        cmd := exec.Command("powershell", "-NoProfile", "Test-Path -Path "+path)
        res, err := cmd.Output()
        if err != nil {
                panic(err)
        }
        status, _ = strconv.ParseBool(string(res))
        return
}

func ListingOfExecuteResults(res []byte, flag string) (list []string) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i, _ := range split {
		split[i] = strings.Trim(split[i], " ")
		if split[i] != flag && !regexp.MustCompile("^-*$").Match([]byte(split[i])) && split[i] != "" {
			list = append(list, split[i])
		}
	}
	return
}

// ---------- ---------- ---------- //

func GetVMList(state string) (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-VM | where {$_.State -eq '"+state+"'} | Select-Object Name")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	list = ListingOfExecuteResults(res, "Name")
	return
}

func GetVmState(name string) (state string) {
	args := "Get-VM '" + name + "' | Select-Object State"
	cmd := exec.Command("powershell", "-NoProfile", args)
	res, err := cmd.Output()
	if err != nil {
		state = "NotFound"
	} else {
		list := ListingOfExecuteResults(res, "State")
		if len(list) == 1 {
			state = list[0]
		} else {
			fmt.Print("error: Unknown error\n")
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
