// Hyper-v operation and management package
package hyperv

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/DevelopNaoki/manahy/internal"
)

func GetVmList() (vmList []Vm, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Format-List VMId, VMName, State, ProcessorCount, MemoryStartup | Out-String).Trim()").Output()
	if err != nil {
		return vmList, fmt.Errorf("failed get vm list: command execution error")
	}

	vms := regexp.MustCompile("\r\n\r\n|\n\n").Split(string(res), -1)
	for i := range vms {
		var vm Vm
		split := regexp.MustCompile("\r\n|\n").Split(vms[i], -1)
		for j := range split {
			switch {
			case strings.Contains(split[j], "VMName"):
				vmName := regexp.MustCompile(":").Split(split[j], -1)
				vm.VmName = strings.TrimSpace(vmName[1])
			case strings.Contains(split[j], "VMId"):
				vmId := regexp.MustCompile(":").Split(split[j], -1)
				vm.VmId = strings.TrimSpace(vmId[1])
			case strings.Contains(split[j], "State"):
				state := regexp.MustCompile(":").Split(split[j], -1)
				vm.State = strings.TrimSpace(state[1])
			case strings.Contains(split[j], "ProcessorCount"):
				processor := regexp.MustCompile(":").Split(split[j], -1)
				vm.Processor = strings.TrimSpace(processor[1])
			case strings.Contains(split[j], "MemoryStartup"):
				memory := regexp.MustCompile(":").Split(split[j], -1)
				vm.Memory, err = internal.ConversionBtoXB(strings.TrimSpace(memory[1]))
				if err != nil {
					return nil, err
				}
			}
		}
		vmList = append(vmList, vm)
	}
	return vmList, nil
}

func GetVmState(vmName string) (status []string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Where-Object VMName -eq '"+vmName+"').State").Output()
	if err != nil {
		return status, fmt.Errorf("failed get vm status: command execution error")
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

func GetVmStateById(vmId string) (status string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Where-Object VMId -eq '"+vmId+"').State").Output()
	if err != nil {
		return status, fmt.Errorf("failed get vm status: command execution error")
	}

	status = strings.TrimSpace(string(res))
	if status == "" {
		return status, fmt.Errorf("%s does not existed", vmId)
	}
	return status, nil
}

func GetVmIdByPath(vmPath string) (vmId string, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-VM | Where-Object ConfigurationLocation -eq '"+vmPath+"').VMId").Output()
	if err != nil {
		return vmId, fmt.Errorf("failed get vmId: command execution error")
	}

	vmId = strings.TrimSpace(string(res))
	if vmId == "" {
		return vmId, fmt.Errorf("vm does not existed in %s", vmPath)
	}
	return vmId, nil
}

func ShowVm(vmName string) (err error) {
	return nil
}

func ShowVmById(vmId string) (err error) {
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
		return fmt.Errorf("failed remove %s: command execution error", vmName)
	}
	return nil
}

func RemoveVmById(vmId string) error {
	_, err := GetVmStateById(vmId)
	if err != nil {
		return err
	}

	err = exec.Command("powershell", "-NoProfile", "Get-VM | Where-Object VMId -wq '"+vmId+"' | Remove-VM -Force").Run()
	if err != nil {
		return fmt.Errorf("failed remove %s: command execution error", vmId)
	}
	return nil
}
