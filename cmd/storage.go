package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/modules"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		storageList, err := modules.GetStorageList()
		if err != nil {
			return err
		}
		displayStorageList(storageList)
		return nil
	},
}
