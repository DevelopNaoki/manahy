package modules

import (
	"os/exec"
)

func GetStorageList() (storageList StorageList, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "Get-Disk | Format-Table Number,FriendlyName,Size").Output()
	if err != nil {
		return storageList, err
	}

	storageList, err = storageListingOfExecuteResults(res)
	if err != nil {
		return storageList, err
	}

	return storageList, nil
}
