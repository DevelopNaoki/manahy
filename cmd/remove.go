package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/modules"
	"github.com/spf13/cobra"
)

var remove = &cobra.Command{
	Use:   "remove",
	Short: "remove vm, disk and switch from manahy.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := modules.UnmarshalYaml("manahy.yaml")
		if err != nil {
			fmt.Printf("%w\n", err)
		}

		for i := range data.Disks {
			err = modules.RemoveDisk(data.Disks[i].Path)
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}

		for i := range data.Networks {
			err = modules.RemoveSwitch(data.Networks[i].Name)
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}
	},
}
