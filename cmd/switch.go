package cmd

import (
	"fmt"

	"github.com/DevelopNaoki/manahy/process"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "management switch on Hyper-V",
	Run: func(cmd *cobra.Command, args []string) {
		process.Error(1)
	},
}

var switchListOption struct {
        external bool
        internal bool
        private  bool
        all      bool
}

var switchList = &cobra.Command{
        Use:   "list",
        Short: "Print switch list",
        Args:  cobra.RangeArgs(0, 1),
        RunE: func(cmd *cobra.Command, args []string) error {
                if switchListOption.external || switchListOption.internal || switchListOption.private {
                        switchListOption.all = false
                }

                if switchListOption.external ||switchListOption.all {
                        fmt.Print("External Switch's\n")
                        externalSwitchs := process.GetSwitchList("External")
                        for _, l := range externalSwitchs {
                                fmt.Printf("- %s\n", l)
                        }
                        fmt.Print("\n")
                }

                if switchListOption.internal || switchListOption.all {
                        fmt.Print("Internal Switch's\n")
                        internalSwitchs := process.GetSwitchList("Internal")
                        for _, l := range internalSwitchs {
                                fmt.Printf("- %s\n", l)
                        }
                        fmt.Print("\n")
                }

                if switchListOption.private || switchListOption.all {
                        fmt.Print("Private switch's\n")
                        privateSwitchs := process.GetSwitchList("Private")
                        for _, l := range privateSwitchs {
                                fmt.Printf("- %s\n", l)
                        }
                        fmt.Print("\n")
                }

                return nil
        },
}

var switchType = &cobra.Command{
        Use:   "type",
        Short: "Print switch type",
        Args:  cobra.RangeArgs(0, 1),
        RunE: func(cmd *cobra.Command, args []string) error {
                if len(args) == 1 {
                        fmt.Print(process.GetSwitchType(args[0]) + "\n")
                } else {
                        fmt.Print("error: Too many or too few arguments\n")
                }

                return nil
        },
}
