package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "start VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) != "Running" {
			process.StartVm(args[0])
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}
