package utils

import (
	"runtime"
	"strings"
)

func GetOSInfo() (string, string) {
	osArch := runtime.GOARCH
	osName := runtime.GOOS
	if osName == "windows" {
		osName = "win"
		if strings.Contains(osArch, "64") {
			osArch = "x64"
		} else {
			osArch = "x86"
		}
	}
	return osName, osArch
}
