package main

import (
	"fmt"
	"os"

	"github.com/DevelopNaoki/manahy/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "")
		os.Exit(-1)
	}
}
