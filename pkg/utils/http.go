package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

var client = &http.Client{}

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
