package puzzles

import (
	"dec2/helpers"
	"log"
	"strconv"
)

func Part2() {
	       ranges, err := helpers.ValidateInput()
	       if err != nil {
		       log.Fatal(err)
	       }

	       var totalSum int
	       for _, r := range ranges {
		       for id := r.StartingValue; id <= r.EndingValue; id++ {
			       s := strconv.Itoa(id)
			       l := len(s)
			       // Try all possible substring lengths
			       for subLen := 1; subLen <= l/2; subLen++ {
				       if l%subLen != 0 {
					       continue
				       }
				       repeatCount := l / subLen
				       if repeatCount < 2 {
					       continue
				       }
				       sub := s[:subLen]
				       if sub[0] == '0' {
					       continue // no leading zeros
				       }
				       valid := true
				       for i := 1; i < repeatCount; i++ {
					       if s[i*subLen:(i+1)*subLen] != sub {
						       valid = false
						       break
					       }
				       }
				       if valid {
					       totalSum += id
					       break // only count once per id
				       }
			       }
		       }
	       }
	       log.Printf("Dec 2, Part 2 - Total Sum: %d", totalSum)
}