package cmd

import (
	"github.com/spf13/cobra"
        "github.com/DevelopNaoki/manahy/process"
)

var RootCmd = &cobra.Command{
	Use:   "manahy",
	Short: "manahy is management tool on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		process.Error(1)
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
		exitCodeList,
	)

	list.Flags().Bool("active", true, "list active vm's")
	list.Flags().Bool("inactive", false, "list inactive vm's")
        list.Flags().Bool("saved", false, "list save vm's")
	list.Flags().Bool("all", false, "list all vm's")
}
