package helpers

import (
	"runtime"
	"strings"
)

// get os name
func GetOsName() string {
	osName := runtime.GOOS
	if osName == "windows" {
		return "win"
	}
	return osName
}

// get os arch
func GetOsArch() string {
	osArch := runtime.GOARCH
	if strings.Contains(osArch, "64") {
		return "x64"
	} else {
		return "x86"
	}
}

func GetFileExt() string {
	if runtime.GOOS == "windows" {
		return "zip"
	} else {
		return "tar.gz"
	}
}
