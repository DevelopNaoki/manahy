package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "management storage",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var storageList = &cobra.Command{
	Use:   "list",
	Short: "listing all storage",
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		storageList, err := process.GetStorageList()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print("Storage\n")
		for i := range storageList.Number {
			fmt.Printf("- %s: %s: %.2f %s\n", storageList.Number[i], storageList.FriendlyName[i], storageList.Size[i], storageList.SizeUnit[i])
		}
		fmt.Print("\n")
		fmt.Print("More infomation, execute 'Get-Disk'\n")
	},
}
