package process

import (
	"fmt"
	"os"
)

func ExitCodeList () {
	fmt.Print("Exit code | status\n")
	fmt.Print("        0 | - It ran successfully\n")
	fmt.Print("        1 | - The command does not exist. Run \"manahy help\" or \"manahy [command] help\", Run the appropriate command again\n")
}

func Error (code int) {
	switch code {
	case 1:
		fmt.Print("error: need valid command\n")
	}
	fmt.Print("If you want to know more detailed information, please execute \"manahy exitcode\"\n")
	os.Exit(code)
}
