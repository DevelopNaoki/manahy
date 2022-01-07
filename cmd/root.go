package cmd

import (
	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "manahy",
	Short: "manahy is management tool on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		process.Error(1)
	},
}
