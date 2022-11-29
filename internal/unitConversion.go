package internal

import (
	"regexp"
	"strconv"
)

// ConversionBtoXB is conversion for unit [0-9]*B to [KMGT]B
// ex) input: 2000B output: 2KB
func ConversionBtoXB(num string) string {
	f, _ := strconv.ParseFloat(num, 64)
	unit := "B"
	for f >= 1000 {
		f = f / 1000
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

// ConversionBtoXB is conversion for unit [KMGT]B to [0-9]*B
// ex) input: 2KB output: 2000B
func ConversionXBtoB(value string) string {
	num := regexp.MustCompile("[0-9.]*").Split(value, -1)[0]
	unit := regexp.MustCompile("[KMGTP]B").Split(value, -1)[0]
	f, _ := strconv.ParseFloat(num, 64)

	for unit != "B" {
		f = f * 1000
		switch unit {
		case "TB":
			unit = "GB"
		case "GB":
			unit = "MB"
		case "MB":
			unit = "KB"
		case "KB":
			unit = "B"
		}
	}
	return strconv.FormatFloat(f, 'f', 0, 64)
}
