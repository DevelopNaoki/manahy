package process

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func isFileExist(path string) (exist bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path "+path+" -PathType Leaf")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	exist, _ = strconv.ParseBool(string(res))
	return
}

func isFolderExist(path string) (exist bool) {
	cmd := exec.Command("powershell", "-NoProfile", "Test-Path "+path)
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	exist, _ = strconv.ParseBool(string(res))
	return
}

func computCapacity(raw int) (processing float32, unit string) {
	processing = float32(raw)
	unit = "B"
	for processing > 1024 {
		processing = processing / 1024
		switch unit {
		case "B":
			unit = "KB"
		case "KB":
			unit = "MB"
		case "MB":
			unit = "GB"
		case "GB":
			unit = "TB"
		case "TB":
			unit = "PB"
		}
	}
	return
}

func listingOfExecuteResults(res []byte, flag string) (list []string) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.Trim(split[i], " ")
		if split[i] != flag && !regexp.MustCompile("^-*$").Match([]byte(split[i])) && split[i] != "" {
			list = append(list, split[i])
		}
	}
	return
}
