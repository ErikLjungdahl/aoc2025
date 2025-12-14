package main

import (
	"fmt"
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

	lines := strings.Split(string(content), "\n")

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
}
