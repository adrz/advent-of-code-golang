package main

import (
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("02-day/input.txt")
	println(part1(data))
}

var Up = map[int]int{
	1: 1,
	2: 2,
	3: 3,
	4: 1,
	5: 2,
	6: 3,
	7: 4,
	8: 5,
	9: 6,
}

var Down = map[int]int{
	1: 4,
	2: 5,
	3: 6,
	4: 7,
	5: 8,
	6: 9,
	7: 7,
	8: 8,
	9: 9,
}

var Left = map[int]int{
	1: 1,
	2: 1,
	3: 2,
	4: 4,
	5: 4,
	6: 5,
	7: 7,
	8: 7,
	9: 8,
}

var Right = map[int]int{
	1: 2,
	2: 3,
	3: 3,
	4: 5,
	5: 6,
	6: 6,
	7: 8,
	8: 9,
	9: 9,
}

func part1(data string) []int {
	code := make([]int, 0)
	for _, line := range strings.Split(data, "\n") {
		pos := 5
		for _, char := range line {
			switch char {
			case 'U':
				pos = Up[pos]
			case 'D':
				pos = Down[pos]
			case 'L':
				pos = Left[pos]
			case 'R':
				pos = Right[pos]
			}
		}
		code = append(code, pos)
	}
	for _, c := range code {
		print(c)
	}
	return code
}
