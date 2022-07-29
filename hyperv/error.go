// hyperv package is manage Hyper-V
package hyperv

import (
	"fmt"
)

func PrintError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s  : [\x1b[31mFalse\x1b[0m]\n  %s\n", msg, err)
	} else {
		fmt.Printf("%s  : [\x1b[32mPass\x1b[0m]\n", msg)
	}
}
