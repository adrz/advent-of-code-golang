package main

import (
	"fmt"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("01-day/input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data string) int {
	floor := 0
	for _, char := range data {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	return floor
}

func part2(data string) int {
	floor := 0
	for i, char := range data {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return 0
}
