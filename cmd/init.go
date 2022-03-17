package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		vmCmd,
		switchCmd,
		diskCmd,
		storageCmd,
		build,
		remove,
	)

	// ---------------------------------- //
	vmCmd.AddCommand(
		vmList,
		vmState,
		vmSave,
		vmStart,
		vmShutdown,
		vmDestroy,
		vmSuspend,
		vmRestart,
		vmConnect,
		vmRemove,
		vmRename,
	)

	vmList.Flags().BoolVarP(&vmListOption.active, "active", "", true, "list active vm's")
	vmList.Flags().BoolVarP(&vmListOption.inactive, "inactive", "i", false, "list inactive vm's")
	vmList.Flags().BoolVarP(&vmListOption.saved, "saved", "s", false, "list save vm's")
	vmList.Flags().BoolVarP(&vmListOption.paused, "paused", "p", false, "list pause vm's")
	vmList.Flags().BoolVarP(&vmListOption.all, "all", "a", false, "list all vm's")

	vmRename.Flags().StringVarP(&newVmName, "new-name", "n", "", "new vm name")

	// ---------------------------------- //
	switchCmd.AddCommand(
		switchList,
		switchChangeCmd,
		switchCreate,
		switchRemove,
		switchRename,
	)

	switchList.Flags().BoolVarP(&switchListOption.external, "external", "e", false, "list external vm's")
	switchList.Flags().BoolVarP(&switchListOption.internal, "internal", "i", false, "list internal vm's")
	switchList.Flags().BoolVarP(&switchListOption.private, "private", "p", false, "list private vm's")
	switchList.Flags().BoolVarP(&switchListOption.all, "all", "a", true, "list all vm's")

	switchCreate.Flags().StringVarP(&switchCreateOption.Name, "name", "n", "", "set name")
	switchCreate.Flags().StringVarP(&switchCreateOption.Type, "type", "t", "", "set type")
	switchCreate.Flags().StringVarP(&switchCreateOption.ExternameInterface, "extername-interface", "", "", "set extername interface")
	switchCreate.Flags().BoolVarP(&switchCreateOption.AllowManagementOs, "allow-management-os", "", false, "set allow management os")

	switchRename.Flags().StringVarP(&newSwitchName, "new-name", "n", "", "rename switch")

	// ---------------------------------- //
	switchChangeCmd.AddCommand(
		switchChangeType,
		switchChangeNetAdapter,
	)

	switchChangeType.Flags().StringVarP(&switchType, "type", "t", "", "change switch type")

	switchChangeNetAdapter.Flags().StringVarP(&switchType, "type", "t", "", "change switch type")
	switchChangeNetAdapter.Flags().StringVarP(&netAdapter, "net-adapter", "n", "", "change network adapter")

	// ---------------------------------- //
	diskCmd.AddCommand(
		diskCreate,
	)

	diskCreate.Flags().StringVarP(&diskCreateOption.Path, "path", "p", "", "set path")
	diskCreate.Flags().StringVarP(&diskCreateOption.Size, "size", "s", "", "set size")
	diskCreate.Flags().StringVarP(&diskCreateOption.Type, "type", "t", "dynamic", "set type")
	diskCreate.Flags().StringVarP(&diskCreateOption.ParentPath, "parent-path", "", "", "set parent path")
	diskCreate.Flags().IntVarP(&diskCreateOption.SourceDisk, "source-disk", "", 0, "set source disk")

	storageCmd.AddCommand(
		storageList,
	)
}
