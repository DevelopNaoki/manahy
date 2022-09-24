package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/manahy/hyperv"
	"github.com/DevelopNaoki/manahy/internal"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "management vm on Hyper-V",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need valid command")
	},
}

var vmListOption struct {
	All      bool
	Active   bool
	Inactive bool
	Saved    bool
	Paused   bool
}
var vmListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print VM list",
	Args:  cobra.RangeArgs(0, 0),
	RunE: func(cmd *cobra.Command, args []string) error {
		// If options are given, disable default options
		if vmListOption.All || vmListOption.Inactive || vmListOption.Saved || vmListOption.Paused {
			vmListOption.Active = false
		}

		vmList, err := hyperv.GetVmList()
		if err != nil {
			return err
		}

		// Find the number of characters of the longest VM name in the list of VMs obtained.
		nameSize := len("Vm Name")
		for i := range vmList {
			if len(vmList[i].VmName) > nameSize {
				nameSize = len(vmList[i].VmName)
			}
		}

		// Add trailing blanks to other VM names to match the number of characters in the longest VM name
		for i := range vmList {
			vmList[i].VmName = internal.SizeAdjustment(vmList[i].VmName, nameSize)
		}

		// List the header string and size and call the function to display the header
		header := []string{
			"VmID",
			"Vm Name",
			"State",
			"Processor",
			"Memory",
		}
		headerSize := []int{
			len(vmList[0].VmId),
			len(vmList[0].VmName),
			len(vmList[0].State),
			len(vmList[0].Processor),
			len(vmList[0].Memory),
		}
		internal.PrintHeader(header, headerSize)

		// Print vm list
		for i := range vmList {
			// Do not display results that do not match the options
			r_disable := vmList[i].State == "Running" && (!vmListOption.Active || !vmListOption.All)
			o_disable := vmList[i].State == "Off" && (vmListOption.Active || vmListOption.Saved || vmListOption.Paused)
			p_disable := vmList[i].State == "Paused" && (!vmListOption.Paused && !vmListOption.All)
                        s_disable := vmList[i].State == "Saved" && (!vmListOption.Saved && !vmListOption.All)
			if r_disable || o_disable || p_disable || s_disable {
				continue
			} else {
				fmt.Printf("%s\t", vmList[i].VmId)
				fmt.Printf("%s\t", vmList[i].VmName)
				fmt.Printf("%s\t", vmList[i].State)
				fmt.Printf("\t%s", vmList[i].Processor)
				fmt.Printf("\t%s\n", vmList[i].Memory)
			}
		}
		return nil
	},
}
