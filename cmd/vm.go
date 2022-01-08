package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "vm is management vm on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		process.Error(1)
	},
}

var vmListOption struct {
	active   bool
	saved    bool
	inactive bool
	all      bool
}

var vmList = &cobra.Command{
	Use:   "list",
	Short: "Print VM list",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmListOption.saved || vmListOption.inactive || vmListOption.all {
			vmListOption.active = false
		}

		if vmListOption.active || vmListOption.all {
			fmt.Print("Running VM's\n")
			activeVms := process.GetVmList("Running")
			for _, l := range activeVms {
				fmt.Printf("- %s\n", l)
			}
			fmt.Print("\n")
		}

		if vmListOption.saved || vmListOption.all {
			fmt.Print("Saved VM's\n")
			savedVms := process.GetVmList("Saved")
			for _, l := range savedVms {
				fmt.Printf("- %s\n", l)
			}
			fmt.Print("\n")
		}

		if vmListOption.inactive || vmListOption.all {
			fmt.Print("Inactive VM's\n")
			inactiveVms := process.GetVmList("Off")
			for _, l := range inactiveVms {
				fmt.Printf("- %s\n", l)
			}
			fmt.Print("\n")
		}

		return nil
	},
}

var vmState = &cobra.Command{
	Use:   "state",
	Short: "Print VM state",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			fmt.Print(process.GetVmState(args[0]) + "\n")
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}

var vmStart = &cobra.Command{
	Use:   "start",
	Short: "start VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) != "Running" {
			process.StartVm(args[0])
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}

var vmSave = &cobra.Command{
	Use:   "save",
	Short: "save VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) == "Running" {
			process.SaveVm(args[0])
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}

var vmShutdown = &cobra.Command{
	Use:   "shutdown",
	Short: "shutdown VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) == "Running" {
			process.StopVm(args[0])
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}

var vmDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) == "Running" {
			process.DestroyVm(args[0])
		} else {
			fmt.Print("error: Too many or too few arguments\n")
		}

		return nil
	},
}

var vmCreate = &cobra.Command{
	Use:   "create",
	Short: "create VM",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
