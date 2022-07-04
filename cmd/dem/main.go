package main

import (
	"dem/pkg/node"
	"fmt"
	"os"
	"strings"
)

const (
	demVersion = "0.0.1"
)

const helpText = `
Usage: dem [command] [flags]
	   dem [-h | --help | -v | --version]

Manage your development environment:
	i, install Install development environment. eg: dem i node or dem i node:8.x or dem i node:8.10.0
`

func main() {
	args := os.Args[1:]
	command := ""
	if len(args) > 1 {
		command = args[0]
	}
	if len(args) == 0 {
		help()
		return
	}
	switch command {
	case "-v", "--version":
		version()
		return
	case "i", "install":
		install()
		return
	default:
		help()
		return
	}
}

func help() {
	fmt.Printf("%s\n", helpText)
}

func version() {
	fmt.Printf("v%s\n", demVersion)
}

func install() {
	args := os.Args[2:]
	fmt.Printf("args: %v\n", args)
	mirror := ""
	if len(args) == 2 && strings.HasPrefix(args[1], "--mirror=") {
		// verify mirror eg: https://npm.taobao.org/mirrors/node
		mirror = strings.TrimPrefix(args[1], "--mirror=")
		fmt.Printf("mirror: %s\n", mirror)
	}
	arr := strings.Split(args[0], ":")
	var ver string
	if len(arr) == 1 {
		ver = ""
	} else {
		ver = arr[1]
	}
	switch arr[0] {
	case "node":
		node.NewNodeEnv(mirror, ver).Install()
	default:
		help()
		return
	}
}
