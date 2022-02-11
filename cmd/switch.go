package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "management switch on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		process.Error(1)
	},
}

var switchListOption struct {
	external bool
	internal bool
	private  bool
	all      bool
}

var switchList = &cobra.Command{
	Use:   "list",
	Short: "Print switch list",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if switchListOption.external || switchListOption.internal || switchListOption.private {
			switchListOption.all = false
		}

		if switchListOption.external || switchListOption.all {
			fmt.Print("External Switch's\n")
			externalSwitchs := process.GetSwitchList("External")
			for _, externamSwitch := range externalSwitchs {
				fmt.Printf("- %s\n", externamSwitch)
			}
			fmt.Print("\n")
		}

		if switchListOption.internal || switchListOption.all {
			fmt.Print("Internal Switch's\n")
			internalSwitchs := process.GetSwitchList("Internal")
			for _, internalSwitch := range internalSwitchs {
				fmt.Printf("- %s\n", internalSwitch)
			}
			fmt.Print("\n")
		}

		if switchListOption.private || switchListOption.all {
			fmt.Print("Private switch's\n")
			privateSwitchs := process.GetSwitchList("Private")
			for _, privateSwitch := range privateSwitchs {
				fmt.Printf("- %s\n", privateSwitch)
			}
			fmt.Print("\n")
		}

		return nil
	},
}


var switchCreateOption process.Network
var switchCreate = &cobra.Command{
	Use:   "create",
	Short: "Create switch",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if switchCreateOption.Name == "" || switchCreateOption.Type == "" {
			fmt.Print("error: Please specify switch name and switch type\n")
		} else if switchCreateOption.Type == "external" && switchCreateOption.ExternameInterface == "" {
			fmt.Print("error: Please specify an external interface\n")
		} else {
			process.CreateSwitch(switchCreateOption)
		}

		return nil
	},
}

var switchRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove switch",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			process.RemoveSwitch(args[0])
		}

		return nil
	},
}
