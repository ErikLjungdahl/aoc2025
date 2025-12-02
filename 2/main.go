package main

import (
	"fmt"
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

	ranges := strings.Split(string(content), ",")

	var invalidIds []int
	var invalidSum int

	for _, r := range ranges {
		rsplit := strings.Split(string(r), "-")
		// fmt.Println("rsplit", rsplit)

		n1, err1 := strconv.Atoi(rsplit[0])
		n2, err2 := strconv.Atoi(rsplit[1])

		if err1 != nil || err2 != nil {
			log.Fatal(err)
		}

		for id := n1; id <= n2; id++ {
			s := strconv.Itoa(id)

			len := len(s)
			// fmt.Println(s, len, s[:len/2], s[len/2:])
			if s[:len/2] == s[len/2:] {
				invalidIds = append(invalidIds, id)
				invalidSum += id
			}
		}
	}

	// fmt.Println(invalidIds)
	fmt.Println(invalidSum)

}
