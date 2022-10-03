package internal

import (
	"strconv"
)

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

func ConversionXBtoB(value string) string {
	num := regexp.MustCompile("[0-9.]*").Split(value, -1)
	unit := regexp.MustCompile("[KMGTP]B").Split(value, -1)
        f, _ := strconv.ParseFloat(num, 64)
	
        for unit == "B" {
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
