package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(content), "\n\n")

	freshRanges := strings.Split(parts[0], "\n")

	availableIngredients := strings.Split(parts[1], "\n")

	var freshCount int = 0
	for _, ingredient := range availableIngredients {
		ing, _ := strconv.Atoi(ingredient)

		var fresh bool = false
		for _, rng := range freshRanges {
			r := strings.Split(rng, "-")
			left, _ := strconv.Atoi(r[0])
			right, _ := strconv.Atoi(r[1])

			if left <= ing && ing <= right {
				// println(left, right, ing)
				fresh = true
				break
			}
		}
		if fresh {
			freshCount++
		}
	}

	println(freshCount)

}
