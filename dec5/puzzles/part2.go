package puzzles

import (
	"dec5/helpers"
	"log"
	"sort"
)

func Part2() {
	ranges, _, err := helpers.ValidateInput()
	if err != nil {
		log.Fatalf("Failed to validate input: %v", err)
	}

	// Sort intervals by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].StartingValue < ranges[j].StartingValue
	})

	// Merge intervals
	var merged []helpers.RangeSet
	for _, curr := range ranges {
		n := len(merged)
		if n == 0 || curr.StartingValue > merged[n-1].EndingValue+1 {
			merged = append(merged, curr)
		} else {
			if curr.EndingValue > merged[n-1].EndingValue {
				merged[n-1].EndingValue = curr.EndingValue
			}
		}
	}

	// Count total unique values
	var total int
	for _, m := range merged {
		total += m.EndingValue - m.StartingValue + 1
	}

	log.Println("Part 2 - Total Unique Values in Ranges:", total)
}
