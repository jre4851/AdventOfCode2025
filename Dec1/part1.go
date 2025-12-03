package main

import (
	"log"
)

func part1() {
	combos, err := validateInput()
	if err != nil {
		log.Fatal(err)
	}

	var currentValue int = 50
	var timesAtZero int = 0
	for _, combo := range combos {
		if combo.Direction == "L" {
			currentValue -= combo.Value
			for currentValue < 0 {
				currentValue += 100
			}
			if currentValue == 0 {
				timesAtZero++
			}
		}
		if combo.Direction == "R" {
			currentValue += combo.Value
			for currentValue >= 100 {
				currentValue -= 100
			}
			if currentValue == 0 {
				timesAtZero++
			}
		}
	}
	log.Printf("Part 1 - Times at Zero: %d", timesAtZero)
}
