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
	var invalidSum2 int

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

			for divider := 1; divider <= len/2; divider++ {
				if isInvalid(s, len, divider) {
					invalidSum2 += id
					break
				}

			}
		}
	}

	// fmt.Println(invalidIds)
	fmt.Println(invalidSum)
	fmt.Println(invalidSum2)

}

func isInvalid(s string, len int, divider int) bool {
	var prevSubset string = s[0:divider]
	for i := divider; i < len; i = i + divider {
		newSubset := s[i:min(i+divider, len)]
		// fmt.Println("inner", prevSubset, newSubset)
		if prevSubset == newSubset {
			prevSubset = newSubset
		} else {
			return false
		}
	}
	// fmt.Println(s)
	return true
}
