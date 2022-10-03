package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// Initializing cobra and setting commands
	cobra.OnInitialize()
	RootCmd.AddCommand(
		versionCmd,
		hypervCmd,
		groupMemberCmd,
                vmCmd,
	)

	// Setting subcommands

	// Add subcommannds for groupCmd
	groupMemberCmd.AddCommand(
                groupMemberListCmd,
                groupMemberAddCmd,
                groupMemberRemoveCmd,
        )

	// Add subcommannds for hypervCmd
	hypervCmd.AddCommand(
		hypervCheckCmd,
		hypervEnableCmd,
	)

	// Add subcommands for vmCmd
	vmCmd.AddCommand(
		vmListCmd,
		vmStartCmd,
		vmResumeCmd,
		vmShutdownCmd,
		vmDestroyCmd,
		vmSaveCmd,
		vmSuspendCmd,
		vmRebootCmd,
	)

	// Setting Options

	// Option of vmListCmd
	vmListCmd.Flags().BoolVarP(&vmListOption.Active, "active", "", true, "display running vm")
	vmListCmd.Flags().BoolVarP(&vmListOption.Inactive, "inactive", "", false, "display power off vm")
	vmListCmd.Flags().BoolVarP(&vmListOption.Saved, "saved", "", false, "display saved vm")
	vmListCmd.Flags().BoolVarP(&vmListOption.Paused, "paused", "", false, "display paused vm")
	vmListCmd.Flags().BoolVarP(&vmListOption.All, "all", "", false, "display all vm")
	// Option of vmStartCmd
	vmStartCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Start vm by id")
	// Option of vmResumeCmd
	vmResumeCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Resume vm by id")
	// Option of vmShutdownCmd
	vmShutdownCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Shutdown vm by id")
	// Option of vmDestroyCmd
	vmDestroyCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Destroy vm by id")
	// Option of vmSaveCmd
	vmSaveCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Save vm by id")
	// Option of vmSuspendCmd
	vmSuspendCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Suspend vm by id")
	// Option of vmRebootCmd
	vmRebootCmd.Flags().StringVarP(&vmId, "vmid", "i", "", "Reboot vm by id")
	vmRebootCmd.Flags().BoolVarP(&vmHardReboot, "force", "f", false, "Hard Reboot vm")
}
