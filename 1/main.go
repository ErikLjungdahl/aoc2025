package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var dial int = 50
	var newDial int = 50

	var atZero int

	var subSteps []int

	for _, line := range lines {

		n, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		// part 2
		for i := 1; i <= n; i++ {
			if line[0] == 'R' {
				newDial = newDial + 1
			} else {
				newDial = newDial - 1
			}

			subSteps = append(subSteps, newDial)

		}

		// part 1
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
	atZero = 0

	for _, line := range subSteps {
		if mod100(line) == 0 {
			atZero += 1
		}
	}

	println(atZero)

}

func mod100(a int) int {
	m := a % 100
	if m < 0 {
		return -m
	}
	return m
}
