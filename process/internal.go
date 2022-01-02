package process

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func isFileExist(path string) (status bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path -Path "+path+" -PathType Leaf")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	status, _ = strconv.ParseBool(string(res))
	return
}

func isFolderExist(path string) (status bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path -Path "+path)
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	status, _ = strconv.ParseBool(string(res))
	return
}

func listingOfExecuteResults(res []byte, flag string) (list []string) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i, _ := range split {
		split[i] = strings.Trim(split[i], " ")
		if split[i] != flag && !regexp.MustCompile("^-*$").Match([]byte(split[i])) && split[i] != "" {
			list = append(list, split[i])
		}
	}
	return
}
