package internal

import (
	"strconv"
)

func ConversionBtoXB(num string) string {
        f, _ := strconv.ParseFloat(num, 64)
        unit := "B"
        for f >= 1024 {
                f = f / 1024
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
