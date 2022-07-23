package cmd

import (
	"github.com/DevelopNaoki/manahy/modules"
	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "create vm, disk and switch from manahy.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := modules.UnmarshalYaml("manahy.yaml")
		if err != nil {
			return err
		}

		for i := range data.Disks {
			err = modules.CreateDisk(data.Disks[i])
			if err != nil {
				return err
			}
		}

		for i := range data.Networks {
			err = modules.CreateSwitch(data.Networks[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
}
