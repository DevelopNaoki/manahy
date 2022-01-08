package process

import (
	"fmt"
	"os/exec"
	"strconv"
)

func GetDiskList() (list []string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-Disk | Format-Table FriendlyName")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	list = listingOfExecuteResults(res, "FriendlyName")
	return
}

func GetDiskId(diskName string) (diskId string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-Disk | where {$_.FriendlyName -eq '"+diskName+"'} | Format-Table Number")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	id := listingOfExecuteResults(res, "Number")
	if len(id) == 1 {
		diskId = id[0]
	}
	return
}

func GetDiskSize(diskName string) (diskSize float32, sizeUnit string) {
	cmd := exec.Command("powershell", "-NoProfile", "Get-Disk | where {$_.FriendlyName -eq '"+diskName+"'} | Format-Table Size")
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	size := listingOfExecuteResults(res, "Size")
	if len(size) == 1 {
		sizeInt, _ := strconv.Atoi(size[0])
		diskSize, sizeUnit = computCapacity(sizeInt)
	}
	return
}

func CreateDisk(newDisk Disk) {
	var args string

	if isFileExist(newDisk.Path) {
		fmt.Print("error: Disk is already exist\n")
	} else {
		switch newDisk.Type {
			case "dynamic":
				args = "New-VHD -Path '"+newDisk.Path+"' -SizeBytes "+newDisk.Size
			case "fixed":
				args = "New-VHD -Path "+newDisk.Path+" -SizeBytes "+newDisk.Size+" -SourceDisk "+strconv.Itoa(newDisk.SourceDisk)+" -Fixed"
			case "differencing":
				args = "New-VHD -ParentPath "+newDisk.ParentPath+" -Path "+newDisk.Path+" -Differencing"
		}
		cmd := exec.Command("powershell", "-NoProfile", args)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
