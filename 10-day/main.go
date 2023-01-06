package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "3113322113"
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(data string) int {
	for i := 0; i < 40; i++ {
		data = lookAndSay(data)
	}
	return len(data)
}

func part2(data string) int {
	for i := 0; i < 50; i++ {
		data = lookAndSay(data)
	}
	return len(data)
}

func lookAndSay(data string) string {
	var result strings.Builder
	for i := 0; i < len(data); i++ {
		count := 1
		for i+1 < len(data) && data[i] == data[i+1] {
			count++
			i++
		}
		result.WriteString(fmt.Sprintf("%d%c", count, data[i]))
	}
	return result.String()
}
