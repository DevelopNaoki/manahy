package cmd

import (
	"github.com/spf13/cobra"
	"github.com/DevelopNaoki/manahy/process"
)


var saved = &cobra.Command{
	Use:   "saved",
	Short: "saved VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.VmOperation(args[0], "saved")

		return nil
	},
}
