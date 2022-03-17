package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var remove = &cobra.Command{
	Use:   "remove",
	Short: "remove vm, disk and switch from manahy.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := process.UnmarshalYaml("manahy.yaml")
		if err != nil {
			fmt.Printf("%w\n", err)
		}

		for i := range data.Disks {
			err = process.RemoveDisk(data.Disks[i].Path)
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}

		for i := range data.Networks {
			err = process.RemoveSwitch(data.Networks[i].Name)
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}
	},
}
