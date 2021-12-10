package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "manahy",
	Short: "manahy is management tool on hyper-v",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("error: need subcommand\n")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
	)
}
