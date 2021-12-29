package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "manahy",
	Short: "manahy is management tool on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("error: need subcommand\n")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		list,
		state,
		start,
		shutdown,
		destroy,
		saved,
	)

	list.Flags().Bool("active", true, "list active vm's")
	list.Flags().Bool("inactive", false, "list inactive vm's")
        list.Flags().Bool("save", false, "list save vm's")
	list.Flags().Bool("all", false, "list all vm's")
}
