package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var switchChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change switch option",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.Error(1)
		return nil
	},
}

var switchChangeType = &cobra.Command{
	Use:   "type",
	Short: "Change switch type",
	Args:  cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			process.ChangeSwitchType(args[0], args[1])
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}
