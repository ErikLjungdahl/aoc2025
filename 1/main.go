package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("inputs/1.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var dial int = 50

	var atZero int

	for _, line := range lines {
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if line[0] == 'R' {
			dial += n
		} else {
			dial -= n
		}

		if dial%100 == 0 {
			atZero += 1
		}
	}

	println(atZero)
}
