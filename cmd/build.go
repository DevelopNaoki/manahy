package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/modules"
	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "create vm, disk and switch from manahy.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := modules.UnmarshalYaml("manahy.yaml")
		if err != nil {
			fmt.Printf("%w\n", err)
		}

		for i := range data.Disks {
			err = modules.CreateDisk(data.Disks[i])
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}

		for i := range data.Networks {
			err = modules.CreateSwitch(data.Networks[i])
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}
	},
}
