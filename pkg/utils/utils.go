package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/schollz/progressbar/v3"
)

var client = &http.Client{}

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

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetRemoteTextFile(url string) string {
	response, httperr := client.Get(url)
	if httperr != nil {
		fmt.Println("\nCould not retrieve " + url + ".\n\n")
		fmt.Printf("%s", httperr)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, readerr := ioutil.ReadAll(response.Body)
		if readerr != nil {
			fmt.Printf("%s", readerr)
			os.Exit(1)
		}
		return string(contents)
	}
	os.Exit(1)
	return ""
}

func DownloadFile(url, target string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	f, _ := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}

func ExtractZipFile(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
