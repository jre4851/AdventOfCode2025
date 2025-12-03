package puzzles

import (
	"dec1/helpers"
	"log"
)

func Part2() {
	combos, err := helpers.ValidateInput()
	if err != nil {
		log.Fatal(err)
	}

	timesAtZero := 0
	currentValue := 50

	for _, combo := range combos {
		steps := combo.Value
		pos := currentValue % 100
		if pos < 0 {
				pos += 100
		}

		if combo.Direction == "L" {
			first := (100 - pos) % 100
			if first == 0 {
					first = 100
			}
			if steps-1 >= first {
					timesAtZero += 1 + ((steps - 1 - first) / 100)
			}

			// Landing at 0
			newPos := (pos + steps) % 100
			if newPos == 0 {
					timesAtZero++
			}
			currentValue = newPos

		} else { 
				first := pos
				if first == 0 {
						first = 100
				}
				if steps-1 >= first {
						timesAtZero += 1 + ((steps - 1 - first) / 100)
				}

				// Landing at 0
				newPos := pos - steps
				newPos %= 100
				if newPos < 0 {
						newPos += 100
				}
				if newPos == 0 {
						timesAtZero++
				}
				currentValue = newPos
		}
}
	
	log.Printf("Dec 1, Part 2 - Times at Zero: %d", timesAtZero)
}
