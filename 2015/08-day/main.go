package main

import (
	"fmt"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("08-day/input.txt")
	diff := part1(data)
	fmt.Println(diff)
	fmt.Println(part2(data))
}

func part1(data string) int {
	totalChar := 0
	totalMem := 0
	for _, str := range strings.Split(data, "\n") {
		totalChar += len(str)
		i := 1
		for i < len(str)-1 {
			char := string(str[i])
			nextChar := string(str[i+1])
			if char == "\\" {
				if nextChar == "x" {
					i += 4
					totalMem++
				} else {
					i += 2
					totalMem++
				}
			} else {
				i++
				totalMem++
			}
		}
	}
	return totalChar - totalMem
}

func part2(data string) int {
	totalChar := 0
	totalAdded := 0
	for _, str := range strings.Split(data, "\n") {
		totalChar += len(str)
		for _, char := range str {
			if char == '"' || char == '\\' {
				totalAdded++
			}
		}
		totalAdded += 2
	}
	return totalAdded
}
