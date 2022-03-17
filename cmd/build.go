package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "create vm, disk and switch from manahy.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := process.UnmarshalYaml("manahy.yaml")
		if err != nil {
			fmt.Printf("%w\n", err)
		}

		for i := range data.Disks {
			err = process.CreateDisk(data.Disks[i])
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}

		for i := range data.Networks {
			err = process.CreateSwitch(data.Networks[i])
			if err != nil {
				fmt.Printf("%w\n", err)
			}
		}
	},
}
