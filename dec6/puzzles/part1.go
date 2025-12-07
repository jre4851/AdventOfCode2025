package puzzles

import (
	"dec6/helpers"
	"log"
	"strconv"
	"strings"
)

func Part1() {
		inputs, err := helpers.ValidateInput()
		if err != nil {
			log.Fatal(err)
		}

		var parts []string
		parts = append(parts, inputs...)

		firstValues := strings.Fields(parts[0])
		secondValues := strings.Fields(parts[1])
		thirdValues := strings.Fields(parts[2])
		fourthValues := strings.Fields(parts[3])
		operators := strings.Fields(parts[4])
		grandTotal := 0

	  for i := 0; i < len(operators); i++ {
			if operators[i] == "+" {
				firstVal, _ := strconv.Atoi(firstValues[i])
				secondVal, _ := strconv.Atoi(secondValues[i])
				thirdVal, _ := strconv.Atoi(thirdValues[i])
				fourthVal, _ := strconv.Atoi(fourthValues[i])
				log.Printf("Adding %d + %d + %d + %d", firstVal, secondVal, thirdVal, fourthVal)
				grandTotal += firstVal + secondVal + thirdVal + fourthVal
			} else {
				firstVal, _ := strconv.Atoi(firstValues[i])
				secondVal, _ := strconv.Atoi(secondValues[i])
				thirdVal, _ := strconv.Atoi(thirdValues[i])
				fourthVal, _ := strconv.Atoi(fourthValues[i])
				log.Printf("Multiplying %d * %d * %d * %d", firstVal, secondVal, thirdVal, fourthVal)
				grandTotal += firstVal * secondVal * thirdVal * fourthVal
			} 
		}
    log.Println("Part 1 Grand Total:", grandTotal)
}
