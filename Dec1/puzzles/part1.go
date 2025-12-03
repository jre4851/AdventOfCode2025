package puzzles

import (
	"dec1/helpers"
	"log"
)

func Part1() {
   combos, err := helpers.ValidateInput()
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
   log.Printf("Dec 1, Part 1 - Times at Zero: %d", timesAtZero)
}
