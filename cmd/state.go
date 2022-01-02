package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var state = &cobra.Command{
	Use:   "state",
	Short: "Print VM state",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			fmt.Print(process.GetVmState(args[0]) + "\n")
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}
