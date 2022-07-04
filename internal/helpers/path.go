package helpers

import (
	"log"
	"os"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
