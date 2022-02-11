package cmd

import (
	"fmt"
)

func Error(code int) {
	switch code {
	case 1:
		fmt.Print("error: need valid command\n")
	case 2:
		fmt.Print("error: need valid options\n")
	}
}
