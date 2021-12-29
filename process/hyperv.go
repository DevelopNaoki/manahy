package process

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func VmList(option string) {
	if option == "running" || option == "all" {
		fmt.Print("Running VM's\n")
		arg := "Get-VM | where {$_.State -eq 'running'} | Select-Object Name"
		execCmdGetVMList(arg)
		fmt.Print("\n")
	}
	if option == "save" || option == "all" {
		fmt.Print("Saved VM's\n")
		arg := "Get-VM | where {$_.State -eq 'saved'} | Select-Object Name"
		execCmdGetVMList(arg)
		fmt.Print("\n")
	}
	if option == "off" || option == "all" {
		fmt.Print("Stop VM's\n")
		arg := "Get-VM | where {$_.State -eq 'off'} | Select-Object Name"
		execCmdGetVMList(arg)
		fmt.Print("\n")
	}
}

func VmState(name string) {
	fmt.Print(execCmdGetVmState(name) + "\n")
}

//--------------- internal function ---------------
func execCmdGetVMList(arg string) {
	cmd := exec.Command("powershell", "-NoProfile", arg)
	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	splitRes := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for index, _ := range splitRes {
		splitRes[index] = strings.Trim(splitRes[index], " ")
		if splitRes[index] != "Name" && splitRes[index] != "----" && splitRes[index] != "" {
			fmt.Printf("- %s\n", splitRes[index])
		}
	}
	return
}

func execCmdGetVmState(name string) (state string) {
	c := "Get-VM '" + name + "' | Select-Object State"
	cmd := exec.Command("powershell", "-NoProfile", c)
	res, err := cmd.Output()
	if err != nil {
		state = "NotFound"
	} else {
		splitRes := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
		for index, _ := range splitRes {
			splitRes[index] = strings.Trim(splitRes[index], " ")
			if splitRes[index] != "Name" && splitRes[index] != "----" && splitRes[index] != "" {
				state = splitRes[index]
			}
		}
	}
	return
}
