package main

import (
	"dem/pkg/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const demVersion = "0.0.1"

const helpTip = `
Usage: dem [command] [flags]
	   dem [-h | --help | -v | --version]

Manage your development environment:
	i, install Install development environment. eg: dem i node or dem i node:8.x or dem i node:8.10.0
`

type Options struct {
	Language string
	Version  string
	Mirror   string
	OsName   string
	OsArch   string
}

var opts Options

func main() {
	args := os.Args
	opts.OsName, opts.OsArch = utils.GetOSInfo()
	if len(args) <= 1 {
		help()
		return
	}

	switch args[1] {
	case "-h", "--help":
		help()
	case "-v", "--version":
		version()
	case "i", "install":
		install(args[2:])
	default:
		help()
	}
}

func help() {
	fmt.Print(helpTip)
}

func version() {
	fmt.Printf("v%s", demVersion)
}

func install(args []string) {
	arr := strings.Split(args[0], ":")

	if len(arr) == 2 {
		if arr[1] == "latest" {
			opts.Version = GetLastestVersion(arr[0])
		} else {
			opts.Version = arr[1]
		}
	} else {
		opts.Version = GetLastestVersion(arr[0])
	}

	fmt.Printf("Installing %s %s...\n", arr, opts.Version)

	switch args[0] {
	case "node":

		// osName, osArch := utils.GetOSInfo()
		// fmt.Println("os name is:", osName)
		// fmt.Println("os arch is:", osArch)
		// fmt.Println("install node")
		// content := utils.GetRemoteTextFile("https://cdn.npmmirror.com/binaries/node/latest/SHASUMS256.txt")
		// re := regexp.MustCompile("node-v(.+)+msi")
		// reg := regexp.MustCompile("node-v|-x.+")
		// version := reg.ReplaceAllString(re.FindString(content), "")
		// fmt.Printf("lastest version: %s\n", version)
		// // downloadUrl := fmt.Sprintf("https://nodejs.org/dist/v%s/node-v%s-%s-%s.zip",
		// // 	version, version, osName, osArch)
		// downloadUrl := fmt.Sprintf("https://npm.taobao.org/mirrors/node/v%s/node-v%s-%s-%s.zip",
		// 	version, version, osName, osArch)
		// // https://registry.npmmirror.com/-/binary/node/v18.4.0/node-v18.4.0-win-x64.zip
		// target := utils.GetCurrentDirectory() + "/node-v18.4.0-win-x64.zip"
		// fmt.Printf("download url: %s, target path: %s", downloadUrl, target)
		// utils.DownloadFile(downloadUrl, target)
		// fmt.Println("extract zip file")
		// utils.ExtractZipFile(target, utils.GetCurrentDirectory())
	default:
		fmt.Println("unknown command")
	}
	// if len(args) == 1 {
	// 	// get latest version
	// 	fmt.Println("get latest version")
	// }
	// if args
}

func GetLastestVersion(language string) string {
	var lastVersion string
	switch language {
	case "node":
		opts.Mirror = "https://nodejs.org/dist/"
		lastVersion = GetNodeLatestVersion()
	}
	return lastVersion
}

// Node
func GetNodeLatestVersion() string {
	url := fmt.Sprintf("%s/latest/SHASUMS256.txt", opts.Mirror)
	fmt.Println(url)
	content := utils.GetRemoteTextFile(url)
	re := regexp.MustCompile("node-v(.+)+msi")
	reg := regexp.MustCompile("node-v|-x.+")
	latestVersion := reg.ReplaceAllString(re.FindString(content), "")
	return latestVersion
}

func InstallNode() string {
	// downloadUrl := fmt.Sprintf("%s/v%s/node-v%s-%s-%s.zip", opts.Mirror, opts.Version, opts.Version, opts.OsName, opts.OsArch)
	// target := utils.GetCurrentDirectory() + "/node/" + opts.Version + "-win-x64.zip"
	return ""
}
