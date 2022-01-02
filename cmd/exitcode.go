package cmd

import (
	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
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
