package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/modules"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "management virtual disk",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var diskCreateOption modules.Disk
var diskCreate = &cobra.Command{
	Use:   "create",
	Short: "Create virtual disk",
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		err := modules.CreateDisk(diskCreateOption)
		if err != nil {
			fmt.Print(err)
		}
		return
	},
}
