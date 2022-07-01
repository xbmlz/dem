package main

import (
	"dem/pkg/utils"
	"fmt"
	"os"
	"regexp"
)

const demVersion = "0.0.1"

const helpTip = `
Usage: dem [command] [flags]
	   dem [-h | --help | -v | --version]

Manage your development environment:
	i, install Install development environment. eg: dem i node or dem i node:8.x or dem i node:8.10.0
`

// type GlobalOptions struct {
// 	osName string // os name eg: linux or darwin or windows
// 	osArch string // os arch eg: aarch64 or x86_64
// }

// var options = &GlobalOptions{
// 	osName: runtime.GOOS,
// 	osArch: utils.GetOSArch(),
// }

func main() {
	args := os.Args
	if len(args) == 1 {
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
	fmt.Println(helpTip)
}

func version() {
	fmt.Printf("v%s", demVersion)
}

func install(args []string) {
	switch args[0] {
	case "node":
		osName, osArch := utils.GetOSInfo()
		fmt.Println("os name is:", osName)
		fmt.Println("os arch is:", osArch)
		fmt.Println("install node")
		content := utils.GetRemoteTextFile("https://cdn.npmmirror.com/binaries/node/latest/SHASUMS256.txt")
		re := regexp.MustCompile("node-v(.+)+msi")
		reg := regexp.MustCompile("node-v|-x.+")
		version := reg.ReplaceAllString(re.FindString(content), "")
		fmt.Printf("lastest version: %s\n", version)
		// downloadUrl := fmt.Sprintf("https://nodejs.org/dist/v%s/node-v%s-%s-%s.zip",
		// 	version, version, osName, osArch)
		downloadUrl := fmt.Sprintf("https://npm.taobao.org/mirrors/node/v%s/node-v%s-%s-%s.zip",
			version, version, osName, osArch)
		// https://registry.npmmirror.com/-/binary/node/v18.4.0/node-v18.4.0-win-x64.zip
		target := utils.GetCurrentDirectory() + "/node-v18.4.0-win-x64.zip"
		fmt.Printf("download url: %s, target path: %s", downloadUrl, target)
		utils.DownloadFile(downloadUrl, target)
		fmt.Println("extract zip file")
		utils.ExtractZipFile(target, utils.GetCurrentDirectory())
	default:
		fmt.Println("unknown command")
	}
	// if len(args) == 1 {
	// 	// get latest version
	// 	fmt.Println("get latest version")
	// }
	// if args
}
