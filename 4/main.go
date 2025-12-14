package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		lines[i] = "." + line + "."
	}
	length := len(lines[0])

	lines = append(lines, strings.Repeat(".", length))
	lines = append([]string{strings.Repeat(".", length)}, lines...)

	for _, line := range lines {
		println(line)
	}

	var accessible int
	for y := 1; y < len(lines)-1; y++ {

		for x := 1; x < len(lines[0])-1; x++ {
			if lines[y][x] == '@' {
				count := strings.Count(lines[y-1][x-1:x+2]+lines[y][x-1:x+2]+lines[y+1][x-1:x+2], "@")
				println(y, x, count, lines[y-1][x-1:x+2]+lines[y][x-1:x+2]+lines[y+1][x-1:x+2])
				if count < 5 {
					accessible++
					println(y, x)
				}
			}
		}
	}

	println(accessible)
}
