package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var vmStart = &cobra.Command{
	Use:   "start",
	Short: "start VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if process.GetVmState(args[0]) != "Running" {
			err := process.StartVm(args[0])
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var vmSave = &cobra.Command{
	Use:   "save",
	Short: "save VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if process.GetVmState(args[0]) == "Running" {
			err := process.SaveVm(args[0])
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var vmShutdown = &cobra.Command{
	Use:   "shutdown",
	Short: "shutdown VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if process.GetVmState(args[0]) == "Running" {
			err := process.StopVm(args[0])
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var vmDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if process.GetVmState(args[0]) == "Running" {
			err := process.DestroyVm(args[0])
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var vmSuspend = &cobra.Command{
	Use:   "suspend",
	Short: "suspend VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if process.GetVmState(args[0]) == "Running" {
			err := process.SuspendVm(args[0])
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var vmRestart = &cobra.Command{
	Use:   "restart",
	Short: "restart VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if process.GetVmState(args[0]) == "Running" {
			err := process.RestartVm(args[0])
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

var vmConnect = &cobra.Command{
	Use:   "connect",
	Short: "connect VM",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		err := process.ConnectVm(args[0])
		if err != nil {
			fmt.Print(err)
		}
	},
}
