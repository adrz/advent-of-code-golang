package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("07-day/input.txt")
	fmt.Println(part1(data))
}

func part1(data string) int {
	myDict := make(map[string]int)
	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "NOT") {
			keys := strings.Split(line, " ")
			key1 := keys[1]
			key2 := keys[3]
			myDict[key2] = ^myDict[key1]
			if myDict[key2] < 0 {
				myDict[key2] = 65536 + myDict[key2]
			}
		} else if strings.Contains(line, "AND") {
			keys := strings.Split(line, " ")
			key1 := keys[0]
			key2 := keys[2]
			key3 := keys[4]
			myDict[key3] = myDict[key1] & myDict[key2]
			if myDict[key3] < 0 {
				myDict[key3] = 65536 + myDict[key3]
			}
		} else if strings.Contains(line, "OR") {
			keys := strings.Split(line, " ")
			key1 := keys[0]
			key2 := keys[2]
			key3 := keys[4]
			myDict[key3] = myDict[key1] | myDict[key2]
			if myDict[key3] < 0 {
				myDict[key3] = 65536 + myDict[key3]
			}

		} else if strings.Contains(line, "LSHIFT") {
			keys := strings.Split(line, " ")
			key1 := keys[0]
			key2 := keys[4]
			key3, _ := strconv.Atoi(keys[2])
			myDict[key2] = myDict[key1] << key3
			if myDict[key2] < 0 {
				myDict[key2] = 65536 + myDict[key2]
			}

		} else if strings.Contains(line, "RSHIFT") {
			keys := strings.Split(line, " ")
			key1 := keys[0]
			key2 := keys[4]
			key3, _ := strconv.Atoi(keys[2])
			myDict[key2] = myDict[key1] >> key3
			if myDict[key2] < 0 {
				myDict[key2] = 65536 + myDict[key2]
			}
		} else {
			keys := strings.Split(line, " ")
			value, _ := strconv.Atoi(keys[0])
			key := keys[2]
			myDict[key] = value
			if myDict[key] < 0 {
				myDict[key] = 65536 + myDict[key]
			}
		}
	}
	return myDict["a"]
}
