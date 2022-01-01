package cmd

import (
	"github.com/spf13/cobra"
	"github.com/DevelopNaoki/manahy/process"
)


var list = &cobra.Command{
	Use:   "list",
	Short: "Print VM status",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var option string

		active, _ := cmd.Flags().GetBool("active")
		if active {
			option = "running"
		}

		inactive, _ := cmd.Flags().GetBool("inactive")
		if inactive {
			option = "off"
		}

                save, _ := cmd.Flags().GetBool("saved")
                if save {
                        option = "save"
                }

		all, _ := cmd.Flags().GetBool("all")
		if all {
			option = "all"
		}

		process.VmList(option)

		return nil
	},
}
