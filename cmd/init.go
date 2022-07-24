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
}

func init() {
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
		vmCreate,
	)

	vmList.Flags().BoolVarP(&vmListOption.active, "active", "", true, "list active vm's")
	vmList.Flags().BoolVarP(&vmListOption.inactive, "inactive", "i", false, "list inactive vm's")
	vmList.Flags().BoolVarP(&vmListOption.saved, "saved", "s", false, "list save vm's")
	vmList.Flags().BoolVarP(&vmListOption.paused, "paused", "p", false, "list pause vm's")
	vmList.Flags().BoolVarP(&vmListOption.all, "all", "a", false, "list all vm's")

	vmRename.Flags().StringVarP(&newVmName, "new-name", "n", "", "new vm name")

	vmCreate.Flags().StringVarP(&vm.Name, "name", "n", "", "new vm name")
        vmCreate.Flags().IntVarP(&vm.Generation, "generation", "g", 1, "set vm generation")
        vmCreate.Flags().IntVarP(&vm.Cpu.Thread, "vcpus", "v", 1, "set vm vcpus")
        vmCreate.Flags().BoolVarP(&vm.Cpu.Nested, "nested", "", false, "enable nested virtualization")
        vmCreate.Flags().StringVarP(&vm.Memory.Size, "memory", "m", "", "set vm memory")
        vmCreate.Flags().BoolVarP(&vm.Memory.Dynamic, "nodynamic", "", false, "disable dynamic memory")
        vmCreate.Flags().StringVarP(&vm.Path, "path", "p", "", "new vm path")
        vmCreate.Flags().StringVarP(&vm.Image, "image", "i", "", "image path")
        vmCreate.Flags().StringVarP(&vmDisk, "disk", "d", "", "disk path")
        vmCreate.Flags().StringVarP(&vmSwitch, "network", "s", "", "switch name")
}

func init() {
	switchCmd.AddCommand(
		switchList,
		switchOptionCfgCmd,
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

	switchOptionCfgCmd.AddCommand(
		switchOptionCfgType,
		switchOptionCfgNetAdapter,
	)

	switchOptionCfgType.Flags().StringVarP(&switchType, "type", "t", "", "change switch type")

	switchOptionCfgNetAdapter.Flags().StringVarP(&switchType, "type", "t", "", "change switch type")
	switchOptionCfgNetAdapter.Flags().StringVarP(&netAdapter, "net-adapter", "n", "", "change network adapter")
}

func init() {
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
