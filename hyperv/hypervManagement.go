// Hyper-v operation and management package
package hyperv

import (
	"fmt"
	"os/exec"
	"strings"
)

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
        hasEnabled, err := CheckHypervEnabled()
	if err != nil {
		return err
	}
	if hasEnabled {
		return fmt.Errorf("Hyper-V is already enabled")
	}

	err = exec.Command("powershell", "-NoProfile", "Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V -All").Run()
	if err != nil {
		return fmt.Errorf("Failed to execute command to enable Hyper-V")
	}
	return nil
}
