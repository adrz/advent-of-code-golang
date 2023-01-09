package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

type Reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
	distance int
}

func main() {
	data := utils.ReadFile("14-day/input.txt")
	fmt.Println(part1(data))
}

func parseReindeer(data string) Reindeer {
	info := strings.Split(data, " ")
	name := info[0]
	speed, _ := strconv.Atoi(info[3])
	flyTime, _ := strconv.Atoi(info[6])
	restTime, _ := strconv.Atoi(info[13])
	return Reindeer{name: name, speed: speed, flyTime: flyTime, restTime: restTime, distance: 0}
}

func (reinder *Reindeer) calcDistance(time int) {
	remainingFlyTime := reinder.flyTime
	remainingRestTime := reinder.restTime
	for t := 0; t < time; t++ {
		if remainingFlyTime > 0 {
			reinder.distance += reinder.speed
			remainingFlyTime--
		} else {
			remainingRestTime--
			if remainingRestTime == 0 {
				remainingFlyTime = reinder.flyTime
				remainingRestTime = reinder.restTime
			}
		}
	}
}

func part1(data string) (int, string) {
	reinders := []Reindeer{}
	for _, line := range strings.Split(data, "\n") {
		reinders = append(reinders, parseReindeer(line))
	}
	maxDistance := 0
	nameMaxDistance := ""

	for _, reinder := range reinders {
		reinder.calcDistance(2503)
		fmt.Println(reinder.name, reinder.distance)
		if reinder.distance > maxDistance {
			maxDistance = reinder.distance
			nameMaxDistance = reinder.name
		}

	}
	fmt.Println(nameMaxDistance)
	return maxDistance, nameMaxDistance
}
