// hyperv package is manage Hyper-V
package hyperv

import (
	"fmt"
	"regexp"
	"strings"
)

var newLine = "\r\n|\n"
var spaceChar = "^[-\\s]*$"

func listingOfExecuteResults(res []byte, flag string) (list []string) {
	split := regexp.MustCompile(newLine).Split(string(res), -1)
	for i := range split {
		split[i] = strings.Trim(split[i], " ")
		if split[i] != flag && !regexp.MustCompile("^-*$").Match([]byte(split[i])) && split[i] != "" {
			list = append(list, split[i])
		}
	}
	return
}

func vmListingOfExecuteResults(res []byte) (vmList VmList, err error) {
	split := regexp.MustCompile(newLine).Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Name") && !regexp.MustCompile(spaceChar).Match([]byte(split[i])) {
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
				return vmList, fmt.Errorf("unknown status for listing vm's")
			}
		}
	}
	return vmList, nil
}

func switchListingOfExecuteResults(res []byte) (switchList SwitchList, err error) {
	split := regexp.MustCompile(newLine).Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Name") && !regexp.MustCompile(spaceChar).Match([]byte(split[i])) {
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
				return switchList, fmt.Errorf("unknown error forlisting  switch's")
			}
		}
	}
	return switchList, nil
}
