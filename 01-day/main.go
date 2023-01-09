package main

import (
	"strconv"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("01-day/input.txt")
	println(part1(data))
	println(part2(data))
}

var Right = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}

var Left = map[string]string{
	"N": "W",
	"W": "S",
	"S": "E",
	"E": "N",
}

var dirs = map[string][2]int{
	"N": {0, 1},
	"E": {1, 0},
	"S": {0, -1},
	"W": {-1, 0},
}

var visited = map[[2]int]bool{}

func part1(data string) int {
	direction := "N"
	posX, posY := 0, 0
	for _, char := range strings.Split(data, ", ") {
		if string(char[0]) == "R" {
			direction = Right[direction]
		} else {
			direction = Left[direction]
		}
		distance, _ := strconv.Atoi(char[1:])
		posX += dirs[direction][0] * distance
		posY += dirs[direction][1] * distance
	}
	return utils.Abs(posX) + utils.Abs(posY)
}

func part2(data string) int {
	direction := "N"
	posX, posY := 0, 0
	for _, char := range strings.Split(data, ", ") {
		if string(char[0]) == "R" {
			direction = Right[direction]
		} else {
			direction = Left[direction]
		}
		distance, _ := strconv.Atoi(char[1:])
		for i := 0; i < distance; i++ {
			posX += dirs[direction][0]
			posY += dirs[direction][1]
			if visited[[2]int{posX, posY}] {
				return utils.Abs(posX) + utils.Abs(posY)
			}
			visited[[2]int{posX, posY}] = true
		}
	}
	return utils.Abs(posX) + utils.Abs(posY)
}
