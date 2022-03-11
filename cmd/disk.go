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
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		diskList := process.GetDiskList()
		fmt.Print("Storage\n")
		for i := range diskList.Number {
			fmt.Printf("- %s: %s: %.2f %s\n", diskList.Number[i], diskList.FriendlyName[i], diskList.Size[i], diskList.SizeUnit[i])
		}
		fmt.Print("\n")
		fmt.Print("More infomation, execute 'Get-Disk'\n")
	},
}

var diskCreateOption process.Disk
var diskCreate = &cobra.Command{
	Use:   "create",
	Short: "Create virtual disk",
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		process.CreateDisk(diskCreateOption)
	},
}