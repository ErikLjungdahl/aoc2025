package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	// part one
	var sum int = 0
	for _, line := range lines {
		var highest string = "00"

		for i := range line {
			highest = max(highest, highest[0:1]+line[i:i+1], highest[1:2]+line[i:i+1])
		}

		n, _ := strconv.Atoi(highest)
		sum += n

	}
	fmt.Println(sum)

	// part two
	var sum2 int = 0
	for _, line := range lines {
		var highest string = line[0:12]

		for i := 1; i < len(line); i++ {
			var options []string
			// 0, if len(line) - i >= 12
			// 1-11 if len(line) - i <= 12
			// max(0, len(line)-12+i)
			for j := max(0, 12-(len(line)-i)); j < min(12, i); j++ {
				options = append(options, highest[0:j]+line[i:i+1]+strings.Repeat("0", max(11-j, 0)))
			}
			options = append(options, highest)
			highest = slices.Max(options)
		}

		fmt.Println(highest)
		n, _ := strconv.Atoi(highest)
		sum2 += n

	}
	fmt.Println(sum2)
}
