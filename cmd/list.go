package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var listOption struct {
	active   bool
	saved    bool
	inactive bool
	all      bool
}

var list = &cobra.Command{
	Use:   "list",
	Short: "Print VM status",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if listOption.saved || listOption.inactive || listOption.all {
			listOption.active = false
		}

		if listOption.active || listOption.all {
			fmt.Print("Running VM's\n")
			activeVms := process.GetVmList("Running")
			for _, l := range activeVms {
				fmt.Printf("- %s\n", l)
			}
			fmt.Print("\n")
		}

		if listOption.saved || listOption.all {
			fmt.Print("Saved VM's\n")
			activeVms := process.GetVmList("Saved")
			for _, l := range activeVms {
				fmt.Printf("- %s\n", l)
			}
			fmt.Print("\n")
		}

		if listOption.inactive || listOption.all {
			fmt.Print("Inactive VM's\n")
			activeVms := process.GetVmList("Off")
			for _, l := range activeVms {
				fmt.Printf("- %s\n", l)
			}
			fmt.Print("\n")
		}

		return nil
	},
}
