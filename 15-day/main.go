package main

import (
	"fmt"
	"strings"

	"github.com/adrz/advent-of-code-golang/utils"
)

func main() {
	data := utils.ReadFile("15-day/input.txt")
	fmt.Println(part1(data))
}

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseIngredient(data string) Ingredient {
	var name string
	var capacity, durability, flavor, texture, calories int
	fmt.Println("ok")
	fmt.Print(data)
	_, err := fmt.Sscanf(data, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)

	if err != nil {
		panic(err)
	}
	return Ingredient{
		name:       name,
		capacity:   capacity,
		durability: durability,
		flavor:     flavor,
		texture:    texture,
		calories:   calories,
	}
}

func part1(data string) int {
	ingredients := []Ingredient{}
	for _, line := range strings.Split(data, "\n") {
		ingredients = append(ingredients, parseIngredient(line))
	}
	maxScore := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				score := calculateScore(ingredients, i, j, k, l)
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}
	return maxScore
}

func calculateScore(ingredients []Ingredient, i, j, k, l int) int {
	capacity := ingredients[0].capacity*i + ingredients[1].capacity*j + ingredients[2].capacity*k + ingredients[3].capacity*l
	durability := ingredients[0].durability*i + ingredients[1].durability*j + ingredients[2].durability*k + ingredients[3].durability*l
	flavor := ingredients[0].flavor*i + ingredients[1].flavor*j + ingredients[2].flavor*k + ingredients[3].flavor*l
	texture := ingredients[0].texture*i + ingredients[1].texture*j + ingredients[2].texture*k + ingredients[3].texture*l
	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0
	}
	return capacity * durability * flavor * texture
}
