package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rang struct {
	start int
	end   int
}

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

	println("part 2")
	var intRanges []rang

	for _, r := range freshRanges {
		r := strings.Split(r, "-")
		left, _ := strconv.Atoi(r[0])
		right, _ := strconv.Atoi(r[1])

		rang := rang{start: left, end: right}
		intRanges = append(intRanges, rang)
	}

	sort.Slice(intRanges, func(i, j int) bool { return intRanges[i].start < intRanges[j].start })

	var newRanges []rang

	var currStart int = intRanges[0].start
	var currEnd int = intRanges[0].end

	for _, rng := range intRanges[1:] {
		// println(rng.start)

		if rng.end < currEnd {
			continue
		}

		if rng.start > currEnd {
			rang := rang{start: currStart, end: currEnd}
			newRanges = append(newRanges, rang)
			currStart = rng.start
		}
		currEnd = rng.end

	}
	rang := rang{start: currStart, end: currEnd}
	newRanges = append(newRanges, rang)

	println("newRanges", len(newRanges))

	var sum int
	for _, rng := range newRanges {
		// println(rng.start, rng.end)
		sum += rng.end + 1 - rng.start
	}

	println(sum)

}
