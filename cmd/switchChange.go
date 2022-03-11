package cmd

import (
	"os"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var switchChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change switch option",
	Run: func(cmd *cobra.Command, args []string) {
		Error(1)
		os.Exit(1)
	},
}

var switchType string
var switchChangeType = &cobra.Command{
	Use:   "type",
	Short: "Change switch type",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.ChangeSwitchType(args[0], switchType)

		return nil
	},
}

var netAdapter string
var switchChangeNetAdapter = &cobra.Command{
	Use:   "adapter",
	Short: "Change network adapter",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.ChangeSwitchNetAdapter(args[0], netAdapter)

		return nil
	},
}
