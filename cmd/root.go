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

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		list,
		state,
		start,
		shutdown,
		destroy,
		save,
		exitCodeList,
	)

	list.Flags().BoolVarP(&listOption.active, "active", "", true, "list active vm's")
	list.Flags().BoolVarP(&listOption.inactive, "inactive", "", false, "list inactive vm's")
	list.Flags().BoolVarP(&listOption.saved, "saved", "", false, "list save vm's")
	list.Flags().BoolVarP(&listOption.all, "all", "", false, "list all vm's")
}
