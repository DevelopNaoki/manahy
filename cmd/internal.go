package cmd

import (
	"fmt"
	"github.com/DevelopNaoki/manahy/modules"
)

func displayList(list []string, message string) {
	fmt.Println(message)
	for i := range list {
		fmt.Println("-", list[i])
	}
	fmt.Println()
}

func displayStorageList(storageList modules.StorageList) {
	fmt.Print("Storage\n")
	for i := range storageList.Number {
		fmt.Printf("- %s: %s: %.2f %s\n", storageList.Number[i], storageList.FriendlyName[i], storageList.Size[i], storageList.SizeUnit[i])
	}
	fmt.Print("\n")
	fmt.Print("More infomation, execute 'Get-Disk'\n")
}
