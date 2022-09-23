package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/hyperv"
	"github.com/spf13/cobra"
)

var hypervCmd = &cobra.Command{
	Use:   "hyperv",
	Short: "management Hyper-V",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var hypervCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check Hyper-V Enabled",
	RunE: func(cmd *cobra.Command, args []string) error {
		isEnable, err := hyperv.CheckHypervEnabled()
		if err != nil {
			return err
		}

		if isEnable {
			fmt.Printf("Check Hyper-V is Enable...\t\t[\x1b[32mEnabled\x1b[0m]\n")
		} else {
			fmt.Printf("Check Hyper-V is Enable...\t\t[\x1b[31mDisabled\x1b[0m]\n")
		}

		return nil
	},
}