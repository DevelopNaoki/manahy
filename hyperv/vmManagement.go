// Hyper-v operation and management package
package hyperv

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetVmList() (vmList []Vm, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Format-List VMId, VMName, State, ProcessorCount, MemoryStartup | Out-String).Trim()").Output()
	if err != nil {
		return vmList, fmt.Errorf("failed get vm list")
	}

	vms := regexp.MustCompile("\r\n\r\n|\n\n").Split(string(res), -1)
	for i := range vms {
		var vm Vm
		split := regexp.MustCompile("\r\n|\n").Split(vms[i], -1)
		for j := range split {
			if strings.Contains(split[j], "VMName") {
				vmName := regexp.MustCompile(":").Split(split[j], -1)
				vm.VmName = strings.TrimSpace(vmName[1])
			} else if strings.Contains(split[j], "VMId") {
				vmId := regexp.MustCompile(":").Split(split[j], -1)
				vm.VmId = strings.TrimSpace(vmId[1])
			} else if strings.Contains(split[j], "State") {
				state := regexp.MustCompile(":").Split(split[j], -1)
				vm.State = strings.TrimSpace(state[1])
			} else if strings.Contains(split[j], "ProcessorCount") {
				processor := regexp.MustCompile(":").Split(split[j], -1)
				vm.Processor = strings.TrimSpace(processor[1])
			} else if strings.Contains(split[j], "MemoryStartup") {
				memory := regexp.MustCompile(":").Split(split[j], -1)
				vm.Memory = unitConversion(strings.TrimSpace(memory[1]))
			}
		}
		vmList = append(vmList, vm)
	}
	return vmList, nil
}

func GetVmState(vmName string) (status []string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Where-Object VMName -eq '"+vmName+"').State").Output()
	if err != nil {
		return status, fmt.Errorf("failed get vm status")
	}

	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if split[i] != "" {
			status = append(status, split[i])
		}
	}

	if len(status) == 0 {
		return status, fmt.Errorf("%s is does not existed", vmName)
	}
	return status, nil
}

func GetVmStateById(vmid string) (status string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Where-Object VMId -eq '"+vmid+"').State").Output()
	if err != nil {
		return status, fmt.Errorf("failed get vm status")
	}

	status = strings.TrimSpace(string(res))
	if status == "" {
		return status, fmt.Errorf("%s does not existed", vmid)
	}
	return status, nil
}

func GetVmIdByPath(vmPath string) (vmid string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Where-Object ConfigurationLocation -eq '"+vmPath+"').VMId").Output()
	if err != nil {
		return vmid, fmt.Errorf("failed get vmid")
	}

	vmid = strings.TrimSpace(string(res))
	if vmid == "" {
		return vmid, fmt.Errorf("vm does not existed in %s", vmPath)
	}
	return vmid, nil
}

func ShowVm(vmName string) (err error) {
	return nil
}

func ShowVmById(vmid string) (err error) {
	return nil
}

func CreateVm() error {
	return nil
}

func RemoveVm(vmName string) error {
	_, err := GetVmState(vmName)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "Remove-VM -Force -Name '"+vmName+"'").Run()
	if err != nil {
		return fmt.Errorf("failed remove vm")
	}
	return nil
}

func RemoveVmById(vmid string) error {
	_, err := GetVmStateById(vmid)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -wq '"+vmid+"' | Remove-VM -Force").Run()
	if err != nil {
		return fmt.Errorf("failed remove vm")
	}
	return nil
}

func unitConversion(num string) string {
	f, _ := strconv.ParseFloat(num, 64)
	unit := "B"
	for f >= 1024 {
		f = f / 1024
		switch unit {
		case "B":
			unit = "KB"
		case "KB":
			unit = "MB"
		case "MB":
			unit = "GB"
		case "GB":
			unit = "TB"
		}
	}
	return strconv.FormatFloat(f, 'f', 0, 64) + unit
}
