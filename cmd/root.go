package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "manahy",
	Short: "manahy is management tool on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		Error(1)
		os.Exit(1)
	},
}
