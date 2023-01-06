package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("12-day/input.txt")
	fmt.Println(part1(data))
}

func part1(data string) int {
	totalSum := 0
	re := regexp.MustCompile(`-?\d+`)
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		num, _ := strconv.Atoi(match[0])
		totalSum += num
	}
	return totalSum
}
