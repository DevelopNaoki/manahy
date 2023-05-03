// Hyper-v operation and management package
package hyperv

import (
	"fmt"
	"os/exec"
)

func ConnectVm(vmName string) error {
	_, err := GetVmState(vmName)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "vmconnect localhost '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed connect %s: command execution error", vmName)
	}
	return nil
}

func ConnectVmById(vmid string) error {
	_, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "vmconnect localhost -G '"+vmid+"'").Run()
	if err != nil {
		return fmt.Errorf("failed connect %s: command execution error", vmid)
	}
	return nil
}

func StartVm(vmName string) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] == "Running" {
			return fmt.Errorf("%s is already Running", vmName)
		}
	}

	err = exec.Command("powershell", "-NoProfile", "Start-VM -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed start %s: command execution error", vmName)
	}
	return nil
}

func StartVmById(vmid string) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState == "Running" {
		return fmt.Errorf("%s is already Running", vmid)
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' | Start-VM").Run()
	if err != nil {
		return fmt.Errorf("fialed start %s: command execution error", vmid)
	}
	return nil
}

func ResumeVm(vmName string) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] != "Pause" {
			return fmt.Errorf("%s is not pause", vmName)
		}
	}

	err = exec.Command("powershell", "-NoProfile", "Resume-VM -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed resume %s: command execution error", vmName)
	}
	return nil
}

func ResumeVmById(vmid string) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState != "Pause" {
		return fmt.Errorf("%s is not pause", vmid)
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' | Resume-VM").Run()
	if err != nil {
		return fmt.Errorf("fialed resume %s: command execution error", vmid)
	}
	return nil
}

func ShutdownVm(vmName string) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] != "Running" {
			return fmt.Errorf("%s is not running", vmName)
		}
	}

	err = exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("fialed shutdown %s: command execution error", vmName)
	}
	return nil
}

func ShutdownVmById(vmid string) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState != "Running" {
		return fmt.Errorf("%s is not running", vmid)
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' | Stop-VM").Run()
	if err != nil {
		return fmt.Errorf("fialed shutdown %s: command execution error", vmid)
	}
	return nil
}

func DestroyVm(vmName string) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] != "Running" {
			return fmt.Errorf("%s is not Running", vmName)
		}
	}

	err = exec.Command("powershell", "-NoProfile", "Stop-VM -Force -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed destroy %s: command execution error", vmName)
	}
	return nil
}

func DestroyVmById(vmid string) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState != "Running" {
		return fmt.Errorf("%s is not running", vmid)
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' | Stop-VM -Force").Run()
	if err != nil {
		return fmt.Errorf("failed destroy %s: command execution error", vmid)
	}
	return nil
}

func SaveVm(vmName string) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] != "Running" {
			return fmt.Errorf("%s is not Running", vmName)
		}
	}

	err = exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed save %s: command execution error", vmName)
	}
	return nil
}

func SaveVmById(vmid string) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState != "Running" {
		return fmt.Errorf("%s is not running", vmid)
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' | Save-VM").Run()
	if err != nil {
		return fmt.Errorf("failed save %s: command execution error", vmid)
	}
	return nil
}

func SuspendVm(vmName string) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] != "Running" {
			return fmt.Errorf("%s is not Running", vmName)
		}
	}

	err = exec.Command("powershell", "-NoProfile", "Suspend-VM -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed suspend %s: command execution error", vmName)
	}
	return nil
}

func SuspendVmById(vmid string) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState != "Running" {
		return fmt.Errorf("%s is not running", vmid)
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' | Suspend-VM").Run()
	if err != nil {
		return fmt.Errorf("failed suspend %s: command execution error", vmid)
	}
	return nil
}

func RebootVm(vmName string, force bool) error {
	vmState, err := GetVmState(vmName)
	if err != nil {
		return err
	}
	for i := range vmState {
		if vmState[i] != "Running" {
			return fmt.Errorf("%s is not Running", vmName)
		}
	}

	rebootCmd := "Restart-VM"
	if force {
		rebootCmd += " -Force"
	}
	err = exec.Command("powershell", "-NoProf", rebootCmd+" -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed reboot %s: command execution error", vmName)
	}
	return nil
}

func RebootVmById(vmid string, force bool) error {
	vmState, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}
	if vmState != "Running" {
		return fmt.Errorf("%s is not running", vmid)
	}

	rebootCmd := "Restart-VM"
	if force {
		rebootCmd += " -Force"
	}
	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -eq '"+vmid+"' |"+rebootCmd).Run()
	if err != nil {
		return fmt.Errorf("failed reboot %s: command execution error", vmid)
	}
	return nil
}
