package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/blang/semver"
)

func main() {
	version := "1.0.21"
	v, _ := semver.Make(version)
	fmt.Println(v.Validate())
	sv := strings.Split(version, ".")
	fmt.Printf("%s\n", sv)
	version = cleanVersion(version)
	fmt.Printf("%s\n", version)
}

func cleanVersion(version string) string {
	re := regexp.MustCompile("\\d+.\\d+.\\d+")
	matched := re.FindString(version)

	if len(matched) == 0 {
		re = regexp.MustCompile("\\d+.\\d+")
		matched = re.FindString(version)
		if len(matched) == 0 {
			matched = version + ".0.0"
		} else {
			matched = matched + ".0"
		}
		fmt.Println(matched)
	}

	return matched
}
