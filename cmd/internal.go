package cmd

import "fmt"

func displayVmList(vmList []string, message string) {
	fmt.Println(message)
	for i := range vmList {
		fmt.Println("-", vmList[i])
	}
	fmt.Println()
}
