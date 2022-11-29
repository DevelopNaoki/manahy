package hyperv

import (
	"fmt"
	"os/exec"
	"regexp"
)

func GetGroupMember() (groupMenbers []string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-LocalGroupMember -Group 'Hyper-V Administrators' | Format-Table Name | Out-String).Trim()").Output()
	if err != nil {
		return groupMenbers, fmt.Errorf("failed get Hyper-V Administrators")
	}

	members := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range members {
		if i < 2 {
			continue
		}

		if members[i] != "" {
			groupMenbers = append(groupMenbers, members[i])
		}
	}
	return groupMenbers, nil
}

func AddGroupMember(name string) error {
	_, err := exec.Command("powershell", "-NoProfile", "Add-LocalGroupMember -Group 'Hyper-V Administrators'-Member "+name).Output()
	if err != nil {
		return fmt.Errorf("failed add group member on Hyper-V Administrators")
	}
	return nil
}

func RemoveGroupMember(name string) error {
	_, err := exec.Command("powershell", "-NoProfile", "Remove-LocalGroupMember -Group 'Hyper-V Administrators'-Member "+name).Output()
	if err != nil {
		return fmt.Errorf("failed remove group member on Hyper-V Administrators")
	}
	return nil
}
