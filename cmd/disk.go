package cmd

import (
	"fmt"
	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "disk is management virtual disk",
	Run: func(cmd *cobra.Command, args []string) {
		process.Error(1)
	},
}

var diskList = &cobra.Command{
	Use:   "list",
	Short: "listing all storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Storage\n")
		diskList := process.GetDiskList()
		for _, l := range diskList {
			diskSize, sizeUnit := process.GetDiskSize(l)
			fmt.Printf("- %s: %s: %.2f %s\n", process.GetDiskId(l), l, diskSize, sizeUnit)
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
