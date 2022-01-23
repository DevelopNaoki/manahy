package process

import (
	"fmt"
	"os/exec"
	"strconv"
)

type DiskList struct {
	Number       []string
	FriendlyName []string
	Size         []float32
	SizeUnit     []string
}

func GetDiskList() (diskList DiskList) {
	diskList.Number = GetDIskListInfo("Number")
	diskList.FriendlyName = GetDIskListInfo("FriendlyName")
	diskSize := GetDIskListInfo("Size")
	for i, _ := range diskSize {
		sizeInt, _ := strconv.Atoi(diskSize[i])
		diskSizeFloat, diskSizeUnit := computCapacity(sizeInt)
		diskList.Size = append(diskList.Size, diskSizeFloat)
		diskList.SizeUnit = append(diskList.SizeUnit, diskSizeUnit)
	}
	return
}

func GetDIskListInfo(parameter string) (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "$res = Get-Disk; echo $res | Format-Table "+parameter)
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	list = listingOfExecuteResults(res, parameter)
	return
}

func CreateDisk(newDisk Disk) {
	var args string

	if isFileExist(newDisk.Path) {
		fmt.Print("error: Disk is already exist\n")
	} else {
		switch newDisk.Type {
		case "dynamic":
			args = "New-VHD -Path '" + newDisk.Path + "' -SizeBytes " + newDisk.Size
		case "fixed":
			args = "New-VHD -Path " + newDisk.Path + " -SizeBytes " + newDisk.Size + " -SourceDisk " + strconv.Itoa(newDisk.SourceDisk) + " -Fixed"
		case "differencing":
			args = "New-VHD -ParentPath " + newDisk.ParentPath + " -Path " + newDisk.Path + " -Differencing"
		}
		cmd := exec.Command("powershell", "-NoProfile", args)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
