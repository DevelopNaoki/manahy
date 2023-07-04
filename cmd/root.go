package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "manahy",
	Short: "manahy is management tool on Hyper-V",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print manahy version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("manahy version 0.0.0 (beta)\n")
	},
}
