package process

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetDiskList() (diskList DiskList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-Disk | Format-Table Number,FriendlyName,Size").Output()
	if err != nil {
		return diskList, err
	}

	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Number") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			diskList.Number = append(diskList.Number, regexp.MustCompile("^[0-9]+").FindString(split[i]))
			split[i] = regexp.MustCompile("^[0-9]+").ReplaceAllString(split[i], "")

			diskSize, diskSizeUnit, err := computCapacity(regexp.MustCompile("[0-9]+$").FindString(split[i]))
			if err != nil {
				return diskList, err
			}
			diskList.Size = append(diskList.Size, diskSize)
			diskList.SizeUnit = append(diskList.SizeUnit, diskSizeUnit)
			split[i] = regexp.MustCompile("[0-9]+$").ReplaceAllString(split[i], "")

			split[i] = strings.TrimSpace(split[i])
			diskList.FriendlyName = append(diskList.FriendlyName, split[i])
		}
	}
	return
}

func CreateDisk(newDisk Disk) error {
	var args, diskSize string
	err := CheckDiskParam(newDisk)

	if err != nil {
		return err
	}
	switch newDisk.Type {
	case "dynamic":
		args = "New-VHD -Path " + newDisk.Path + " -SizeBytes " + newDisk.Size
	case "fixed":
		args = "New-VHD -Path " + newDisk.Path + " -SizeBytes " + diskSize + " -SourceDisk " + strconv.Itoa(newDisk.SourceDisk) + " -Fixed"
	case "differencing":
		args = "New-VHD -ParentPath " + newDisk.ParentPath + " -Path " + newDisk.Path + " -Differencing"
	}

	err = exec.Command("powershell", "-NoProfile", args).Run()
	if err != nil {
		return err
	}
	return nil
}
