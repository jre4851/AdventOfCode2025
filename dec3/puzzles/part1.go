package puzzles

import (
	"dec3/helpers"
	"log"
	"math/big"
	"strconv"
)

func Part1() {
	inputs, err := helpers.ValidateInput()
	if err != nil {
		log.Fatal(err)
	}
	
	var maxDigit int
	var maxDigitIndex int
	var secondMaxDigit int
	var grandTotal int
	// Find the largest digit and its index in the array
	for _, inputStr := range inputs {
		// Use math/big to handle arbitrarily large numbers
		var bigNum big.Int
		bigNum.SetString(inputStr, 10)
		numStr := bigNum.String()
		maxDigit = -1
		maxDigitIndex = -1
		for i := 0; i < len(numStr)-1; i++ {
			digit := int(numStr[i] - '0')
			if digit > maxDigit {
				maxDigit = digit
				maxDigitIndex = i
			}
		}
		// Find the largest digit among numbers after maxDigitIndex	
		secondMaxDigit = -1
		for i := maxDigitIndex + 1; i < len(numStr); i++ {
			digit := int(numStr[i] - '0')
			if digit > secondMaxDigit {
				secondMaxDigit = digit
			}
		}
		
		var maxDigitStr = strconv.Itoa(maxDigit)
		var secondMaxDigitStr = strconv.Itoa(secondMaxDigit)
		totalSum, err := strconv.Atoi(maxDigitStr + secondMaxDigitStr)
		if err != nil {
			log.Fatal(err)
		}
		grandTotal += totalSum
	}

	log.Printf("Part 1 Total: %d", grandTotal)
}
