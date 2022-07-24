package modules

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func isFileExist(path string) (exist bool, err error) {
	res, e := exec.Command("powershell", "-NoProfile", "Test-Path \""+path+"\"").Output()
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

func vmListingOfExecuteResults(res []byte) (vmList VmList, err error) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Name") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			state := regexp.MustCompile("Running$|Saved$|Off$|Paused$").FindString(split[i])
			split[i] = regexp.MustCompile("Running$|Saved$|Off$|Paused$").ReplaceAllString(split[i], "")
			split[i] = strings.TrimSpace(split[i])

			switch state {
			case "Running":
				vmList.Running = append(vmList.Running, split[i])
			case "Saved":
				vmList.Saved = append(vmList.Saved, split[i])
			case "Off":
				vmList.Off = append(vmList.Off, split[i])
			case "Paused":
				vmList.Paused = append(vmList.Paused, split[i])
			default:
				return vmList, fmt.Errorf("Unknown status for vm list")
			}
		}
	}
	return vmList, nil
}

func switchListingOfExecuteResults(res []byte) (switchList SwitchList, err error) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Name") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			switchType := regexp.MustCompile("External$|Internal$|Private$").FindString(split[i])
			split[i] = regexp.MustCompile("External$|Internal$|Private$").ReplaceAllString(split[i], "")
			split[i] = strings.TrimSpace(split[i])

			switch switchType {
			case "External":
				switchList.External = append(switchList.External, split[i])
			case "Internal":
				switchList.Internal = append(switchList.Internal, split[i])
			case "Private":
				switchList.Private = append(switchList.Private, split[i])
			default:
				return switchList, fmt.Errorf("Unknown error for switch list")
			}
		}
	}
	return switchList, nil
}

func storageListingOfExecuteResults(res []byte) (storageList StorageList, err error) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Number") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			storageList.Number = append(storageList.Number, regexp.MustCompile("^[0-9]+").FindString(split[i]))
			split[i] = regexp.MustCompile("^[0-9]+").ReplaceAllString(split[i], "")

			storageSize, storageSizeUnit, err := computCapacity(regexp.MustCompile("[0-9]+$").FindString(split[i]))
			if err != nil {
				return storageList, err
			}
			storageList.Size = append(storageList.Size, storageSize)
			storageList.SizeUnit = append(storageList.SizeUnit, storageSizeUnit)
			split[i] = regexp.MustCompile("[0-9]+$").ReplaceAllString(split[i], "")

			split[i] = strings.TrimSpace(split[i])
			storageList.FriendlyName = append(storageList.FriendlyName, split[i])
		}
	}
	return storageList, nil
}
