package cmd

import (
	"github.com/spf13/cobra"
	"github.com/DevelopNaoki/manahy/process"
)


var exitCodeList = &cobra.Command{
	Use:   "exitcode",
	Short: "show exit code list",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.ExitCodeList()
		return nil
	},
}
