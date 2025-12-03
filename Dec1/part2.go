package main

import (
	"log"
)

func part2() {
	combos, err := validateInput()
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
            // Count hits at 0 during rotation (exclude final landing).
            // We need k in [1..steps-1] such that (pos + k) % 100 == 0.
            first := (100 - pos) % 100
            if first == 0 {
                // first "hit" would be at k=0 (current position), not during a click, so shift to 100
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

        } else { // "R"
            // Count hits at 0 during rotation (exclude final landing).
            // We need k in [1..steps-1] such that (pos - k) % 100 == 0  => k ≡ pos (mod 100).
            // Smallest positive k is pos if pos > 0; if pos == 0, smallest positive is 100.
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
	// for _, combo := range combos {
		// if combo.Direction == "L" {
		// 	start := currentValue - combo.Value + 1
		// 	end := currentValue
		// 	for i := start; i <= end; i++ {
		// 		if ((i % 100) + 100) % 100 == 0 {
		// 			timesAtZero++
		// 		}
		// 	}
		// 	currentValue -= combo.Value
		// 	for currentValue < 0 {
		// 		currentValue += 100
		// 	}
		// }
		// if combo.Direction == "R" {
		// 	start := currentValue + 1
		// 	end := currentValue + combo.Value
		// 	for i := start; i <= end; i++ {
		// 		if (i % 100) == 0 {
		// 			timesAtZero++
		// 		}
		// 	}
		// 	currentValue += combo.Value
		// 	for currentValue >= 100 {
		// 		currentValue -= 100
		// 	}
		// }
	
	log.Printf("Part 2 - Times at Zero: %d", timesAtZero)
}
