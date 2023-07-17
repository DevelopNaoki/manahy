package cmd

import (
	"fmt"
	"strings"

	"github.com/DevelopNaoki/manahy/hyperv"
	"github.com/DevelopNaoki/manahy/internal"
	"github.com/spf13/cobra"
)

var vmswitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "management vmswitch on Hyper-V",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var vmswitchListOption struct {
	All      bool
	External bool
	internal bool
	Private  bool
	Type     string
}
var vmswitchListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print VMSwitch list",
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {
		// If options are given, disable default options
		if vmswitchListOption.External || vmswitchListOption.internal || vmswitchListOption.Private {
			vmswitchListOption.All = false
		}

		fmt.Print("Get VMSwitch List...")
		vmswitchList, err := hyperv.GetVmswitchList()
		if err != nil {
			fmt.Print("\rGet VMSwitch List...\t [\x1b[31mFalse\x1b[0m]\n")
			return err
		}
		fmt.Print("\rGet VMSwitch List...\t [\x1b[32mOK\x1b[0m]\n")

		// Find the number of characters of the longest VMSwitch
		// name in the list of VMSwitchs obtained.
		nameSize := len("VMSwitch Name")
		for i := range vmswitchList {
			if len(vmswitchList[i].VmswitchName) > nameSize {
				nameSize = len(vmswitchList[i].VmswitchName)
			}
		}

		// Add trailing blanks to other VMSwitch names to match
		// the number of characters in the longest VMSwitch name
		for i := range vmswitchList {
			vmswitchList[i].VmswitchName = internal.SizeAdjustment(vmswitchList[i].VmswitchName, nameSize)
		}

		// List the header string and size and call the function to display the header
		header := []string{"VMSwitch ID", "VMSwitch Name", "VMSwitch Type"}
		headerSize := []int{
			len(vmswitchList[0].VmswitchId),
			len(vmswitchList[0].VmswitchName),
			len(vmswitchList[0].VmswitchType),
		}
		internal.PrintHeader(header, headerSize)

		// Print vmswitch list
		for i := range vmswitchList {
			if !strings.EqualFold(vmswitchList[i].VmswitchType, vmswitchListOption.Type) && !strings.EqualFold(vmswitchListOption.Type, "all") {
				continue
			}

			fmt.Printf("%s\t", vmswitchList[i].VmswitchId)
			fmt.Printf("%s\t", vmswitchList[i].VmswitchName)
			fmt.Printf("%s\t", vmswitchList[i].VmswitchType)

			fmt.Print("\n")
		}
		return nil
	},
}

var vmswitchCreateOption hyperv.NewVmswitch
var vmswitchCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create VMSwitch list",
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Creating VMSwitch...")
		err := hyperv.CreateVmswitch(vmswitchCreateOption)
		if err != nil {
			fmt.Print("Create VMSwitch...\t [\x1b[31mFalse\x1b[0m]\n")
			return err
		}
		fmt.Print("Create VMSwitch...\t [\x1b[32mOK\x1b[0m]\n")
		return nil
	},
}
