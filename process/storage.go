package process

import (
	"os/exec"
	"regexp"
	"strings"
)

func GetStorageList() (storageList StorageList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-Disk | Format-Table Number,FriendlyName,Size").Output()
	if err != nil {
		return storageList, err
	}

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
	return
}
