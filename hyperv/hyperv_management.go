// Hyper-v operation and management package
package hyperv

import (
	"fmt"
	"os/exec"
	"strings"
)

// Check is Hyper-V enable
func IsEnabled() (enabled bool, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V).State").Output()
	if err != nil {
		return false, fmt.Errorf("failed check if Hyper-V is enabled: command execution error")
	}
	if strings.TrimSpace(string(res)) == "Enabled" {
		return true, nil
	}
	return false, nil
}

func EnableHyperv() error {
	isEnabled, err := IsEnabled()
	if err != nil {
		return err
	}
	if isEnabled {
		return fmt.Errorf("Hyper-V is already enabled")
	}

	err = exec.Command("powershell", "-NoProfile", "Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V -All").Run()
	if err != nil {
		return fmt.Errorf("failed to enable Hyper-V: command execution error")
	}
	return nil
}
