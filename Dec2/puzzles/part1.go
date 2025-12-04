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

	var totalCount int
	var maxInputRange int = ranges[len(ranges)-1].EndingValue
	var possibleCombinations []int
		for i := 1; i <= maxInputRange/2; i++ {
			var candidate string = strconv.Itoa(i) + strconv.Itoa(i)
			val, err := strconv.Atoi(candidate)
			if err == nil {
				possibleCombinations = append(possibleCombinations, val)
			}
		}
	for _, r := range ranges {
		
	}
	
	log.Printf("Dec 2, Part 1 - Total Count: %d", totalCount)
}