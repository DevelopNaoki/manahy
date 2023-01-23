package hyperv

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func GetVmswitchList() (vmswitchList []Vmswitch, err error) {
	res, err := exec.Command("powershell", "-NoProfile", "(Get-Vmswitch | Format-List Id, Name, SwitchType | Out-String).Trim()").Output()
	if err != nil {
		return vmswitchList, fmt.Errorf("failed get vm list: command execution error")
	}

	vms := regexp.MustCompile("\r\n\r\n|\n\n").Split(string(res), -1)
	for i := range vms {
		var vmswitch Vmswitch
		split := regexp.MustCompile("\r\n|\n").Split(vms[i], -1)
		for j := range split {
			switch {
			case strings.Contains(split[j], "Name"):
				SwitchName := regexp.MustCompile(":").Split(split[j], -1)
				vmswitch.VmswitchName = strings.TrimSpace(SwitchName[1])
			case strings.Contains(split[j], "Id"):
				SwitchId := regexp.MustCompile(":").Split(split[j], -1)
				vmswitch.VmswitchId = strings.TrimSpace(SwitchId[1])
			case strings.Contains(split[j], "SwitchType"):
				SwitchType := regexp.MustCompile(":").Split(split[j], -1)
				vmswitch.VmswitchType = strings.TrimSpace(SwitchType[1])
			}
		}
		vmswitchList = append(vmswitchList, vmswitch)
	}
	return vmswitchList, nil
}
