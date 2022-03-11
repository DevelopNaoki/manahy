package process

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetDiskList() (diskList DiskList) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-Disk | Format-Table Number,FriendlyName,Size").Output()
	if err != nil {
		panic(err)
	}

	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if !strings.Contains(split[i], "Number") && !regexp.MustCompile("^[-\\s]*$").Match([]byte(split[i])) {
			diskList.Number = append(diskList.Number, regexp.MustCompile("^[0-9]+").FindString(split[i]))
			split[i] = regexp.MustCompile("^[0-9]+").ReplaceAllString(split[i], "")

			diskSize, diskSizeUnit := computCapacity(regexp.MustCompile("[0-9]+$").FindString(split[i]))
			diskList.Size = append(diskList.Size, diskSize)
			diskList.SizeUnit = append(diskList.SizeUnit, diskSizeUnit)
			split[i] = regexp.MustCompile("[0-9]+$").ReplaceAllString(split[i], "")

			split[i] = strings.TrimSpace(split[i])
			diskList.FriendlyName = append(diskList.FriendlyName, split[i])
		}
	}

	return
}

func CreateDisk(newDisk Disk) {
	var args, diskSize string
	CheckDiskParam(newDisk)

	switch newDisk.Type {
	case "dynamic":
		args = "New-VHD -Path " + newDisk.Path + " -SizeBytes " + newDisk.Size
	case "fixed":
		args = "New-VHD -Path " + newDisk.Path + " -SizeBytes " + diskSize + " -SourceDisk " + strconv.Itoa(newDisk.SourceDisk) + " -Fixed"
	case "differencing":
		args = "New-VHD -ParentPath " + newDisk.ParentPath + " -Path " + newDisk.Path + " -Differencing"
	}

	err := exec.Command("powershell", "-NoProfile", args).Run()
	if err != nil {
		panic(err)
	}
}

func CheckDiskParam(newDisk Disk) {
	if isFileExist(newDisk.Path) {
		fmt.Print("error: Disk is already exist\n")
		os.Exit(1)
	}
	if newDisk.Type != "dynamic" && newDisk.Type != "fixed" && newDisk.Type != "differencing" {
		fmt.Print("error:  Undefined DiskType\n")
		os.Exit(1)
	}
	if newDisk.Type == "differencing" && !isFileExist(newDisk.ParentPath) {
		fmt.Print("error: Disk doesnot exist\n")
		os.Exit(1)
	}
	diskSize := regexp.MustCompile("^[0-9]*[TGM]B$").FindString(newDisk.Size)
	if diskSize == "" {
		fmt.Print("error: undefined size\n")
		os.Exit(1)
	}
}
