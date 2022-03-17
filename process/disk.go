package process

import (
	"os/exec"
	"strconv"
)

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
