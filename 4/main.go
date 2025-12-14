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

	// for _, line := range lines {
	// 	println(line)
	// }

	var accessible int
	var last int
	for i := 0; ; i++ {

		for y := 1; y < len(lines)-1; y++ {

			for x := 1; x < len(lines[0])-1; x++ {
				if lines[y][x] == '@' {
					countAt := strings.Count(lines[y-1][x-1:x+2]+lines[y][x-1:x+2]+lines[y+1][x-1:x+2], "@")
					countX := strings.Count(lines[y-1][x-1:x+2]+lines[y][x-1:x+2]+lines[y+1][x-1:x+2], "x")
					count := countAt + countX
					// println(y, x, count, lines[y-1][x-1:x+2]+lines[y][x-1:x+2]+lines[y+1][x-1:x+2])
					if count < 5 {
						accessible++
						// println(y, x)
						lines[y] = lines[y][0:x] + "x" + lines[y][x+1:len(lines[0])]
					}
				}
			}
		}

		if i == 0 {
			println(accessible)
		}

		if accessible == last {
			println(accessible)
			break
		}
		last = accessible

		for y := 1; y < len(lines)-1; y++ {

			for x := 1; x < len(lines[0])-1; x++ {
				if lines[y][x] == 'x' {
					lines[y] = lines[y][0:x] + "." + lines[y][x+1:len(lines[0])]
				}
			}
		}
	}

}
