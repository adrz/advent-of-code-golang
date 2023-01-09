package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("06-day/input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func getCoord(line string) (int, int) {
	coord := strings.Split(line, ",")
	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])
	return x, y
}

func part1(data string) int {
	grid := [1000][1000]int{}
	for _, line := range strings.Split(data, "\n") {
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "turn" {
			x_start, y_start := getCoord(lineSplit[2])
			x_end, y_end := getCoord(lineSplit[4])
			if lineSplit[1] == "on" {
				for x := x_start; x <= x_end; x++ {
					for y := y_start; y <= y_end; y++ {
						grid[x][y] = 1
					}
				}
			}
			if lineSplit[1] == "off" {
				for x := x_start; x <= x_end; x++ {
					for y := y_start; y <= y_end; y++ {
						grid[x][y] = 0
					}
				}
			}
		} else if lineSplit[0] == "toggle" {
			x_start, y_start := getCoord(lineSplit[1])
			x_end, y_end := getCoord(lineSplit[3])
			for x := x_start; x <= x_end; x++ {
				for y := y_start; y <= y_end; y++ {
					if grid[x][y] == 0 {
						grid[x][y] = 1
					} else {
						grid[x][y] = 0
					}
				}
			}
		}
	}
	nbLights := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] == 1 {
				nbLights++
			}
		}
	}
	return nbLights
}

func part2(data string) int {
	grid := [1000][1000]int{}
	for _, line := range strings.Split(data, "\n") {
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "turn" {
			x_start, y_start := getCoord(lineSplit[2])
			x_end, y_end := getCoord(lineSplit[4])
			if lineSplit[1] == "on" {
				for x := x_start; x <= x_end; x++ {
					for y := y_start; y <= y_end; y++ {
						grid[x][y] += 1
					}
				}
			}
			if lineSplit[1] == "off" {
				for x := x_start; x <= x_end; x++ {
					for y := y_start; y <= y_end; y++ {
						if grid[x][y] > 0 {
							grid[x][y] -= 1
						}
					}
				}
			}
		} else if lineSplit[0] == "toggle" {
			x_start, y_start := getCoord(lineSplit[1])
			x_end, y_end := getCoord(lineSplit[3])
			for x := x_start; x <= x_end; x++ {
				for y := y_start; y <= y_end; y++ {
					grid[x][y] += 2
				}
			}
		}
	}
	brighness := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			brighness += grid[x][y]
		}
	}
	return brighness
}
