package cmd

import (
        "fmt"
        "github.com/spf13/cobra"

        "github.com/DevelopNaoki/manahy/hyperv"
)

var groupMemberCmd = &cobra.Command{
        Use:   "member",
        Short: "management Group Member on Hyper-V Administrators",
        RunE: func(cmd *cobra.Command, args []string) error {
                return fmt.Errorf("need valid command")
        },
}

var groupMemberListCmd = &cobra.Command{
        Use:   "list",
        Short: "Show group member on Hyper-V Administrators",
        RunE: func(cmd *cobra.Command, args []string) error {
                groupMembers, err := hyperv.GetGroupMember()
		if err != nil {
			return err
		}

		fmt.Printf("Hyper-V Administrators\n")
		for i := range groupMembers {
			fmt.Printf("- %s\n", groupMembers[i])
		}
		return nil
        },
}

var groupMemberAddCmd = &cobra.Command{
        Use:   "add",
        Short: "Add group member on Hyper-V Administrators",
        Args:  cobra.RangeArgs(0, 100),
        RunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			err := hyperv.AddGroupMember(args[i])
			if err != nil {
				return err
			}
		}
		return nil
        },
}

var groupMemberRemoveCmd = &cobra.Command{
        Use:   "remove",
        Short: "Remove group member on Hyper-V Administrators",
        Args:  cobra.RangeArgs(0, 100),
        RunE: func(cmd *cobra.Command, args []string) error {
                for i := range args {
                        err := hyperv.RemoveGroupMember(args[i])
                        if err != nil {
                                return err
                        }
                }
                return nil
        },
}

