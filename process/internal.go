package process

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func isFileExist(path string) (exist bool) {
	res, err := exec.Command("powershell", "-NoProfile", "Test-Path "+path).Output()
	if err != nil {
		panic(err)
	}
	exist, _ = strconv.ParseBool(strings.Replace(string(res), "\r\n", "", -1))
	return
}

func computCapacity(raw string) (processing float64, unit string) {
	processing, _ = strconv.ParseFloat(raw, 64)
	unit = "B"
	for processing >= 1024 {
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
