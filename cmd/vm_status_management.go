package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/hyperv"
	"github.com/spf13/cobra"
)

var vmId string
var vmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Starting VM...")
		if vmId != "" {
			err := hyperv.StartVmById(vmId)
			if err != nil {
				fmt.Print("Starting VM...\t [\x1b[31mFalse\x1b[0m]\n")
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.StartVm(args[i])
			if err != nil {
				fmt.Print("Starting VM...\t [\x1b[31mFalse\x1b[0m]\n")
				return err
			}
		}
		fmt.Print("Starting VM...\t [\x1b[32mOK\x1b[0m]\n")
		return nil
	},
}

var vmResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "resume VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmId != "" {
			err := hyperv.ResumeVmById(vmId)
			if err != nil {
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.ResumeVm(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var vmShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "shutdown VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmId != "" {
			err := hyperv.ShutdownVmById(vmId)
			if err != nil {
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.ShutdownVm(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var vmDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmId != "" {
			err := hyperv.DestroyVmById(vmId)
			if err != nil {
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.DestroyVm(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var vmSaveCmd = &cobra.Command{
	Use:   "save",
	Short: "save VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmId != "" {
			err := hyperv.SaveVmById(vmId)
			if err != nil {
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.SaveVm(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var vmSuspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "suspend VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		if vmId != "" {
			err := hyperv.SuspendVmById(vmId)
			if err != nil {
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.SuspendVm(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

var vmForceReboot bool
var vmRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot VM",
	Args:  cobra.RangeArgs(0, 100),
	RunE: func(cmd *cobra.Command, args []string) error {
		// vmForceReboot is force reboot
		if vmId != "" {
			err := hyperv.RebootVmById(vmId, vmForceReboot)
			if err != nil {
				return err
			}
		}
		// Multiple VM name specification supported
		for i := range args {
			err := hyperv.RebootVm(args[i], vmForceReboot)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
