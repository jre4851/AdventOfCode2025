package puzzles

import (
	"dec2/helpers"
	"log"
)

func Part1() {
	ranges, err := helpers.ValidateInput()
	if err != nil {
		log.Fatal(err)
	}

	var totalCount int
	for _, r := range ranges {
		count := r.EndingValue - r.StartingValue + 1
		totalCount += count
	}
	
	log.Printf("Dec 2, Part 1 - Total Count: %d", totalCount)
}