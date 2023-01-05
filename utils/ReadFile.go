package utils

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(pathFromCaller string) string {
	// Get the path to the caller
	data, err := os.ReadFile(pathFromCaller)
	check(err)
	strContent := string(data)
	return strings.TrimRight(strContent, "\n")
}
