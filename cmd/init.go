package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		exitCodeList,
		vmCmd,
		switchCmd,
		diskCmd,
	)

	vmCmd.AddCommand(
		vmList,
		vmState,
		vmSave,
		vmStart,
		vmShutdown,
		vmDestroy,
	)

	vmList.Flags().BoolVarP(&vmListOption.active, "active", "", true, "list active vm's")
	vmList.Flags().BoolVarP(&vmListOption.inactive, "inactive", "", false, "list inactive vm's")
	vmList.Flags().BoolVarP(&vmListOption.saved, "saved", "", false, "list save vm's")
	vmList.Flags().BoolVarP(&vmListOption.all, "all", "", false, "list all vm's")

	switchCmd.AddCommand(
		switchList,
		switchType,
		switchChangeCmd,
		switchCreate,
		switchRemove,
	)

	switchList.Flags().BoolVarP(&switchListOption.external, "external", "", false, "list external vm's")
	switchList.Flags().BoolVarP(&switchListOption.internal, "internal", "", false, "list internal vm's")
	switchList.Flags().BoolVarP(&switchListOption.private, "private", "", false, "list private vm's")
	switchList.Flags().BoolVarP(&switchListOption.all, "all", "", true, "list all vm's")

	switchCreate.Flags().StringVarP(&switchCreateOption.Name, "name", "", "", "set name")
	switchCreate.Flags().StringVarP(&switchCreateOption.Type, "type", "", "", "set type")
	switchCreate.Flags().StringVarP(&switchCreateOption.ExternameInterface, "extername-interface", "", "", "set extername interface")
	switchCreate.Flags().BoolVarP(&switchCreateOption.AllowManagementOs, "allow-management-os", "", false, "set allow management os")

	switchChangeCmd.AddCommand(
		switchChangeType,
	)

	diskCmd.AddCommand(
		diskList,
		diskCreate,
	)

	diskCreate.Flags().StringVarP(&diskCreateOption.Path, "path", "", "", "set path")
        diskCreate.Flags().StringVarP(&diskCreateOption.Size, "size", "", "", "set size")
        diskCreate.Flags().StringVarP(&diskCreateOption.Type, "type", "", "dynamic", "set type")
        diskCreate.Flags().StringVarP(&diskCreateOption.ParentPath, "parent-path", "", "", "set parent path")
        diskCreate.Flags().IntVarP(&diskCreateOption.SourceDisk, "source-disk", "", 0, "set source disk")
}
