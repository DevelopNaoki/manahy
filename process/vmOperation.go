package process

import "os/exec"

// ConnectVm connect the VM
func ConnectVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "vmconnect localhost '"+name+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

// StartVm start the VM
func StartVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "Start-VM '"+name+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

// StopVm stop the VM
func StopVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "Stop-VM -Name '"+name+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

// DestroyVm force stop VM
func DestroyVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "Stop-VM -Force -Name '"+name+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

// SaveVm save VM
func SaveVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "Save-VM -Name '"+name+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

func SuspendVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "Suspend-VM -Name '"+name+"'").Run()
	if err != nil {
		return err
	}
	return nil
}

func RestartVm(name string) error {
	err := exec.Command("powershell", "-NoProfile", "Restart-VM -Name '"+name+"' -Force").Run()
	if err != nil {
		return err
	}
	return nil
}
