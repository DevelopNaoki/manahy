package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var switchOptionCfgCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure switch option",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var switchType string
var switchChangeOptionCfgType = &cobra.Command{
	Use:   "type",
	Short: "Configure switch type",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		err := process.ChangeSwitchType(args[0], switchType)
		if err != nil {
			fmt.Print(err)
		}
	},
}

var netAdapter string
var switchOptionCfgNetAdapter = &cobra.Command{
	Use:   "adapter",
	Short: "Configure network adapter",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		err := process.ChangeSwitchNetAdapter(args[0], netAdapter)
		if err != nil {
			fmt.Print(err)
		}
	},
}
