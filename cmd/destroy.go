package cmd

import (
	"github.com/spf13/cobra"
	"github.com/DevelopNaoki/manahy/process"
)


var destroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.VmOperation(args[0], "destroy")

		return nil
	},
}
