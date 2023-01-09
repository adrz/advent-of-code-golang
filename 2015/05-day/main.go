package main

import (
	"fmt"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("05-day/input.txt")
	fmt.Println(part1(data))
}

func checkIfNice(str string) bool {
	if len(str) < 3 {
		return false
	}
	if !strings.ContainsAny(str, "aeiou") {
		return false
	}
	for i := 0; i < len(str)-1; i++ {
		if str[i:i+1] == "ab" {
			return false
		} else if str[i:i+1] == "cd" {
			return false
		} else if str[i:i+1] == "pq" {
			return false
		} else if str[i:i+1] == "xy" {
			return false
		}
	}

	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func part1(data string) int {
	nbNice := 0
	for _, str := range strings.Split(data, "\n") {
		if checkIfNice(str) {
			nbNice++
		}
	}
	return nbNice
}
