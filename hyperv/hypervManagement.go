// Hyper-v operation and management package
package hyperv

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsAdmin() (isAdmin bool, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)").Output()
	if err != nil {
		return false, fmt.Errorf("Failed to execute the command to check if the user has administrative privileges.")
	}
	if strings.TrimSpace(string(res)) == "True" {
		return true, nil
	}
	return false, nil
}

func CheckHypervEnabled() (enabled bool, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V).State").Output()
	if err != nil {
		return false, fmt.Errorf("Failed to execute command to check if Hyper-V is enabled")
	}
	if strings.TrimSpace(string(res)) == "Enabled" {
		return true, nil
	}
	return false, nil
}

func EnableHyperv() error {
	isAdmin, err := IsAdmin()
	if err != nil {
		return err
	}
	if isAdmin {
		enabled, err := CheckHypervEnabled()
		if err != nil {
			return err
		}
		if enabled {
			return fmt.Errorf("Hyper-V is already enabled")
		}
	} else {
		return fmt.Errorf("You do not have execute permission")
	}

	err = exec.Command("powershell", "-NoProfile", "Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V -All").Run()
	if err != nil {
		return fmt.Errorf("Failed to execute command to enable Hyper-V")
	}

	enabled, err := CheckHypervEnabled()
	if err != nil {
		return err
	}
	if !enabled {
		return fmt.Errorf("Failed to enable Hyper-V")
	}
	return nil
}
