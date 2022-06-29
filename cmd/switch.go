package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "management switch on Hyper-V",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
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
	Run: func(cmd *cobra.Command, args []string) {
		if switchListOption.external || switchListOption.internal || switchListOption.private {
			switchListOption.all = false
		}

		switchList, err := process.GetSwitchList()
		if err != nil {
			fmt.Print(err)
		}

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
	},
}

var switchCreateOption process.Network
var switchCreate = &cobra.Command{
	Use:   "create",
	Short: "Create switch",
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		if switchCreateOption.Name == "" || switchCreateOption.Type == "" {
			fmt.Print("error: Please specify switch name and switch type\n")
		} else if switchCreateOption.Type == "external" && switchCreateOption.ExternameInterface == "" {
			fmt.Print("error: Please specify an external interface\n")
		} else {
			err := process.CreateSwitch(switchCreateOption)
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var switchRemove = &cobra.Command{
	Use:   "remove",
	Short: "Remove switch",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		err := process.RemoveSwitch(args[0])
		if err != nil {
			fmt.Print(err)
		}
	},
}

var newSwitchName string
var switchRename = &cobra.Command{
	Use:   "rename",
	Short: "Rename switch",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if newSwitchName == "" {
			fmt.Print("error: need new switch name\n")
		} else {
			err := process.RenameSwitch(args[0], newSwitchName)
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

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
