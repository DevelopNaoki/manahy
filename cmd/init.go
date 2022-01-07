package cmd

import(
        "github.com/spf13/cobra"
)

func init() {
        cobra.OnInitialize()
        RootCmd.AddCommand(
                exitCodeList,
		vmCmd,
		switchCmd,
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
	)

	switchList.Flags().BoolVarP(&switchListOption.external, "external", "", false, "list external vm's")
        switchList.Flags().BoolVarP(&switchListOption.internal, "internal", "", false, "list internal vm's")
        switchList.Flags().BoolVarP(&switchListOption.private, "private", "", false, "list private vm's")
        switchList.Flags().BoolVarP(&switchListOption.all, "all", "", true, "list all vm's")

	switchChangeCmd.AddCommand(
		switchChangeType,
	)
}
