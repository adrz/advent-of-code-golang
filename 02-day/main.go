package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("02-day/input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func min(a, b, c int) int {
	minimum := a
	if b < minimum {
		minimum = b
	}
	if c < minimum {
		minimum = c
	}
	return minimum
}

func part1(data string) int {
	totalArea := 0
	for _, line := range strings.Split(data, "\n") {
		lwh := strings.Split(line, "x")
		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])
		totalArea += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}
	return totalArea
}

func part2(data string) int {
	totalArea := 0
	for _, line := range strings.Split(data, "\n") {
		lwh := strings.Split(line, "x")
		lwh_int := make([]int, len(lwh))
		for i, v := range lwh {
			lwh_int[i], _ = strconv.Atoi(v)
		}
		sort.Ints(lwh_int)
		totalArea += 2*lwh_int[0] + 2*lwh_int[1] + lwh_int[0]*lwh_int[1]*lwh_int[2]
	}
	return totalArea
}
