package puzzles

import (
	"dec2/helpers"
	"log"

	"strconv"
)

func Part1() {
	ranges, err := helpers.ValidateInput()
	if err != nil {
		log.Fatal(err)
	}

	var totalSum int
	for _, r := range ranges {
		for id := r.StartingValue; id <= r.EndingValue; id++ {
			s := strconv.Itoa(id)
			l := len(s)
			if l%2 != 0 {
				continue // must be even length
			}
			half := l / 2
			if s[:half] == s[half:] && s[0] != '0' {
				totalSum += id
			}
		}
	}
	log.Printf("Dec 2, Part 1 - Total Sum: %d", totalSum)
}