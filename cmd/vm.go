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
	paused   bool
	all      bool
}

var vmList = &cobra.Command{
	Use:   "list",
	Short: "Print VM list",
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmListOption.saved || vmListOption.inactive || vmListOption.paused || vmListOption.all {
			vmListOption.active = false
		}

		vmList, err := process.GetVmList()

		if err != nil {
			return err
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

		return nil
	},
}

var vmState = &cobra.Command{
	Use:   "state",
	Short: "Print VM state",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(process.GetVmState(args[0]) + "\n")

		return nil
	},
}

var vmStart = &cobra.Command{
	Use:   "start",
	Short: "start VM",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) != "Running" {
			process.StartVm(args[0])
		}

		return nil
	},
}

var vmSave = &cobra.Command{
	Use:   "save",
	Short: "save VM",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) == "Running" {
			process.SaveVm(args[0])
		}

		return nil
	},
}

var vmShutdown = &cobra.Command{
	Use:   "shutdown",
	Short: "shutdown VM",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) == "Running" {
			process.StopVm(args[0])
		}

		return nil
	},
}

var vmDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy VM",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if process.GetVmState(args[0]) == "Running" {
			process.DestroyVm(args[0])
		}

		return nil
	},
}

var vmSuspend = &cobra.Command{
        Use:   "suspend",
        Short: "suspend VM",
        Args:  cobra.RangeArgs(1, 1),
        RunE: func(cmd *cobra.Command, args []string) error {
                if process.GetVmState(args[0]) == "Running" {
                        process.SuspendVm(args[0])
                }

                return nil
        },
}

var vmRestart = &cobra.Command{
        Use:   "restart",
        Short: "restart VM",
        Args:  cobra.RangeArgs(1, 1),
        RunE: func(cmd *cobra.Command, args []string) error {
                if process.GetVmState(args[0]) == "Running" {
                        process.RestartVm(args[0])
                }

                return nil
        },
}

var vmConnect = &cobra.Command{
	Use:   "connect",
	Short: "connect VM",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.ConnectVm(args[0])

		return nil
	},
}

var vmCreate = &cobra.Command{
	Use:   "create",
	Short: "create VM",
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

var vmRemove = &cobra.Command{
	Use:   "remove",
	Short: "remove VM",
	Args:  cobra.RangeArgs(1, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, index := range args {
			process.RemoveVm(index)
		}
		return nil
	},
}

var newVmName string
var vmRename = &cobra.Command{
	Use:   "rename",
	Short: "rename VM",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if newVmName == "" {
			fmt.Print("error: need new vm name\n")
		} else {
			process.RenameVm(args[0], newVmName)
		}
		return nil
	},
}
