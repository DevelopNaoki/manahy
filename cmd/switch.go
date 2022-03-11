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
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if switchListOption.external || switchListOption.internal || switchListOption.private {
			switchListOption.all = false
		}

		switchList := process.GetSwitchLists()

		if switchListOption.external || switchListOption.all {
                        fmt.Print("External Switch's\n")
			for i := range switchList.External {
				fmt.Printf("- %s\n", switchList.External[i])
			}
			fmt.Print("\n")
		}

		if switchListOption.internal || switchListOption.all {
                        fmt.Print("Internal Switch's\n")
			for i := range switchList.Internal {
				fmt.Printf("- %s\n", switchList.Internal[i])
			}
			fmt.Print("\n")
		}

		if switchListOption.private || switchListOption.all {
                        fmt.Print("Private switch's\n")
			for i := range switchList.Private {
				fmt.Printf("- %s\n", switchList.Private[i])
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
	Args:  cobra.RangeArgs(0, 0),
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
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.RemoveSwitch(args[0])

		return nil
	},
}

var newSwitchName string
var switchRename = &cobra.Command{
	Use:   "rename",
	Short: "Rename switch",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if newSwitchName == "" {
			fmt.Print("error: need new switch name\n")
		} else {
			process.RenameSwitch(args[0], newSwitchName)
		}

		return nil
	},
}
