package utils

import "github.com/fatih/color"

func ColorPrint(c int, format string, a ...interface{}) string {
	switch c {
	case -1:
		return color.HiRedString("[-] "+format, a...)
	case 0:
		return color.HiGreenString("[+] "+format, a...)
	case 1:

		return color.HiBlueString(format, a...)
	default:

		return color.HiWhiteString("[+] "+format, a...)
	}
}
