package modules

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

func CreateDisk(newDisk Disk) error {
	err := checkDiskParam(newDisk)
	if err != nil {
		return err
	}

	cmd := "New-VHD -Path " + newDisk.Path
	switch newDisk.Type {
	case "dynamic":
		cmd += " -SizeBytes " + newDisk.Size
	case "fixed":
		cmd += " -SizeBytes " + newDisk.Size
		cmd += " -SourceDisk " + strconv.Itoa(newDisk.SourceDisk)
		cmd += " -Fixed"
	case "differencing":
		cmd += " -ParentPath " + newDisk.ParentPath
		cmd += " -Differencing"
	}

	err = exec.Command("powershell", "-NoProfile", cmd).Run()
	if err != nil {
		return err
	}

	return nil
}

func RemoveDisk(path string) error {
	exist, err := isFileExist(path)
	if err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("error: %s is not exist", path)
	}

	err = exec.Command("powershell", "-NoProfile", "rm '"+path+"'").Run()
	if err != nil {
		return fmt.Errorf("failed delete %s", path)
	}

	return nil
}

func checkDiskParam(newDisk Disk) error {
	diskExist, err := isFileExist(newDisk.Path)
	if err != nil {
		return err
	} else if !diskExist && newDisk.Import {
		return fmt.Errorf("error: %s is not exist", newDisk.Path)
	} else if diskExist && !newDisk.Import {
		return fmt.Errorf("error: %s is already exist", newDisk.Path)
	}

	err = checkDiskTypeParam(newDisk.Type)
	if err != nil {
		return err
	}

	switch newDisk.Type {
	case "differencing":
		parentDiskExist, err := isFileExist(newDisk.ParentPath)
		if err != nil {
			return err
		} else if !parentDiskExist {
			return fmt.Errorf("error: Parent disk is not exist")
		}
	case "Fixed":
	}

	err = checkDiskSizeParam(newDisk.Size)
	if err != nil {
		return err
	}

	return nil
}

func checkDiskTypeParam(diskType string) error {
	switch diskType {
	case "dynamic":
	case "fixed":
	case "differencing":
	default:
		return fmt.Errorf("error: Undefined disk type")
	}

	return nil
}

func checkDiskSizeParam(diskSize string) error {
	diskSize = regexp.MustCompile("^[0-9]*[TGM]B$").FindString(diskSize)
	if diskSize == "" {
		return fmt.Errorf("error: unknown size")
	}

	return nil
}
