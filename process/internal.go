package process

import (
	"regexp"
	"strings"
)

func listingOfExecuteResults(res []byte, flag string) (list []string) {
	split := regexp.MustCompile("\r\n|\n").Split(string(res), -1)
	for i, _ := range split {
		split[i] = strings.Trim(split[i], " ")
		if split[i] != flag && !regexp.MustCompile("^-*$").Match([]byte(split[i])) && split[i] != "" {
			list = append(list, split[i])
		}
	}
	return
}
