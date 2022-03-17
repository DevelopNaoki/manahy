package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "management vm on Hyper-V",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var vmListOption struct {
	active   bool
	saved    bool
	inactive bool
	paused   bool
	all      bool
}

var vmList = &cobra.Command{
	Use:   "list",
	Short: "Print VM list",
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
		if vmListOption.saved || vmListOption.inactive || vmListOption.paused || vmListOption.all {
			vmListOption.active = false
		}

		vmList, err := process.GetVmList()

		if err != nil {
			fmt.Print(err)
		}

		if vmListOption.active || vmListOption.all {
			fmt.Print("Running VM's\n")
			for i := range vmList.Running {
				fmt.Printf("- %s\n", vmList.Running[i])
			}
			fmt.Print("\n")
		}

		if vmListOption.saved || vmListOption.all {
			fmt.Print("Saved VM's\n")
			for i := range vmList.Saved {
				fmt.Printf("- %s\n", vmList.Saved[i])
			}
			fmt.Print("\n")
		}

		if vmListOption.paused || vmListOption.all {
			fmt.Print("Paused VM's\n")
			for i := range vmList.Paused {
				fmt.Printf("- %s\n", vmList.Paused[i])
			}
			fmt.Print("\n")
		}

		if vmListOption.inactive || vmListOption.all {
			fmt.Print("Inactive VM's\n")
			for i := range vmList.Off {
				fmt.Printf("- %s\n", vmList.Off[i])
			}
			fmt.Print("\n")
		}
	},
}

var vmState = &cobra.Command{
	Use:   "state",
	Short: "Print VM state",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(process.GetVmState(args[0]) + "\n")
	},
}

var vmCreate = &cobra.Command{
	Use:   "create",
	Short: "create VM",
	Args:  cobra.RangeArgs(0, 0),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var vmRemove = &cobra.Command{
	Use:   "remove",
	Short: "remove VM",
	Args:  cobra.RangeArgs(1, 100),
	Run: func(cmd *cobra.Command, args []string) {
		for _, index := range args {
			err := process.RemoveVm(index)
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var newVmName string
var vmRename = &cobra.Command{
	Use:   "rename",
	Short: "rename VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if newVmName == "" {
			fmt.Print("error: need new vm name\n")
		} else {
			err := process.RenameVm(args[0], newVmName)
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}
