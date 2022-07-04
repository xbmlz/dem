package node

import (
	"dem/internal/helpers"
	"fmt"
	"os"
	"path"
	"regexp"
)

type NodeEnvOptions struct {
	root      string
	mirror    string // eg: https://nodejs.org/dist/
	proxy     string // eg: http://
	osArch    string // eg: x64
	osName    string // eg: win
	verifyssl bool   // eg: true
	version   string // eg: 8.10.0
	fileExt   string // eg: zip
	// args      []string // eg: ["node:8.10.0"]
}

func NewNodeEnv(m string, args []string) *NodeEnvOptions {
	mirror := "https://nodejs.org/dist/"
	version := "latest"
	if m != "" {
		mirror = m
	}
	rootPath := helpers.GetCurrentDirectory() + "/dist/node/"
	helpers.MakeDir(rootPath)
	opts := &NodeEnvOptions{
		root:      rootPath,
		mirror:    mirror,
		proxy:     "none",
		osArch:    helpers.GetOsArch(),
		osName:    helpers.GetOsName(),
		verifyssl: true,
		version:   version,
		fileExt:   helpers.GetFileExt(),
	}
	return opts
}

func (env *NodeEnvOptions) Install() {
	if env.version == "latest" {
		env.version = env.GetLastestVersion()
	}
	if env.IsInstalled(env.version) {
		fmt.Printf("node %s is already installed !\n", env.version)
		return
	}
	fmt.Printf("Installing node %s\n", env.version)
	packageName := fmt.Sprintf("node-v%s-%s-%s.%s", env.version, env.osName, env.osArch, env.fileExt)
	downloadUrl := fmt.Sprintf("%s/v%s/%s", env.mirror, env.version, packageName)
	savePath := path.Join(env.root, packageName)
	// fmt.Printf("downloading %s, save path is %s\n", downloadUrl, savePath)
	// download
	helpers.DownloadFile(downloadUrl, savePath)
	// extract
	fmt.Printf("extracting %s\n", packageName)
	if err := helpers.ExtractFile(savePath, env.root); err != nil {
		fmt.Printf("extract file error: %s\n", err.Error())
	} else {
		os.Remove(savePath)
	}
	fmt.Printf("node %s installed\n", env.version)
}

func (env *NodeEnvOptions) GetLastestVersion() string {
	url := fmt.Sprintf("%s/latest/SHASUMS256.txt", env.mirror)
	content := helpers.GetRemoteTextFile(url)
	re := regexp.MustCompile("node-v(.+)+msi")
	reg := regexp.MustCompile("node-v|-x.+")
	return reg.ReplaceAllString(re.FindString(content), "")
}

func (env *NodeEnvOptions) IsInstalled(src string) bool {
	nodeDir := fmt.Sprintf("node-v%s-%s-%s", env.version, env.osName, env.osArch)
	nodeExcutable := fmt.Sprintf("%s/%s/node.exe", env.root, nodeDir)
	return helpers.IsFileExists(nodeExcutable)
}

func (env *NodeEnvOptions) GetVersion(version string) string {

	return env.version
}
