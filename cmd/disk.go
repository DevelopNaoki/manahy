package cmd

import (
	"fmt"
	"os"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "disk is management virtual disk",
	Run: func(cmd *cobra.Command, args []string) {
		Error(1)
		os.Exit(1)
	},
}

var diskList = &cobra.Command{
	Use:   "list",
	Short: "listing all storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Storage\n")
		volumeList := process.GetDiskList()
		for i := range volumeList.Number {
			fmt.Printf("- %s: %s: %.2f %s\n", volumeList.Number[i], volumeList.FriendlyName[i], volumeList.Size[i], volumeList.SizeUnit[i])
		}
		fmt.Print("\n")
	},
}

var diskCreateOption process.Disk
var diskCreate = &cobra.Command{
	Use:   "create",
	Short: "Create virtual disk",
	Run: func(cmd *cobra.Command, args []string) {
		process.CreateDisk(diskCreateOption)
	},
}
