package main

import (
	"fmt"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("03-day/input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

var dirs = map[string][2]int{
	"^": [2]int{-1, 0}, // row then col
	"v": [2]int{1, 0},
	"<": [2]int{0, -1},
	">": [2]int{0, 1},
}

func part1(data string) int {
	myDict := make(map[[2]int]int)
	position := [2]int{0, 0}
	for _, char := range data {
		dir := dirs[string(char)]
		position[0] += dir[0]
		position[1] += dir[1]
		myDict[position]++
	}
	return len(myDict)
}

func part2(data string) int {
	myDict := make(map[[2]int]int)
	positionSanta := [2]int{0, 0}
	positionRobot := [2]int{0, 0}
	for i, char := range data {
		dir := dirs[string(char)]
		if i%2 == 0 {
			positionSanta[0] += dir[0]
			positionSanta[1] += dir[1]
			myDict[positionSanta]++
			continue
		}
		positionRobot[0] += dir[0]
		positionRobot[1] += dir[1]
		myDict[positionRobot]++
	}
	return len(myDict)
}
