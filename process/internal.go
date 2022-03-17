package process

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func isFileExist(path string) (exist bool, err error) {
	res, e := exec.Command("powershell", "-NoProfile", "Test-Path '"+path+"'").Output()
	if e != nil {
		return false, fmt.Errorf("error: failed execute 'Test-Path'")
	}
	exist, _ = strconv.ParseBool(strings.Replace(string(res), "\r\n", "", -1))
	return exist, nil
}

func computCapacity(raw string) (processing float64, unit string, err error) {
	processing, err = strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, "", fmt.Errorf("error: in type conversion")
	} else {
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
			default:
				return 0, "", fmt.Errorf("error: %s is undefine", unit)
			}
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
