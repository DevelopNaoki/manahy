package process

import (
	"os/exec"
	"strconv"
)

func IsExist(name string) (exist bool) {
	if GetVmState(name) == "NotFound" {
		exist = false
	} else {
		exist = true
	}
	return
}

func IsFileExist(path string) (exist bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path -Path "+path+" -PathType Leaf")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	exist, _ = strconv.ParseBool(string(res))
	return
}

func IsFolderExist(path string) (exist bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path -Path "+path)
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	exist, _ = strconv.ParseBool(string(res))
	return
}
