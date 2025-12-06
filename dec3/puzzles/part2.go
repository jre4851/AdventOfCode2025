package puzzles

import (
	"dec3/helpers"
	"log"
	"math/big"
)

func Part2() {
    inputs, err := helpers.ValidateInput()
    if err != nil {
        log.Fatal(err)
    }

    grandTotal := big.NewInt(0)
    for _, inputStr := range inputs {
        k := 12
        n := len(inputStr)
        if n < k {
            k = n
        }
        // Greedy algorithm: build the largest number by picking k digits in order
        result := make([]byte, 0, k)
        start := 0
        for i := 0; i < k; i++ {
            maxDigit := byte('0')
            maxPos := start
            // The range to search is from start to n - (k - i)
            end := n - (k - i)
            for j := start; j <= end; j++ {
                if inputStr[j] > maxDigit {
                    maxDigit = inputStr[j]
                    maxPos = j
                }
            }
            result = append(result, maxDigit)
            start = maxPos + 1
        }
        bankJoltage := big.NewInt(0)
        bankJoltage.SetString(string(result), 10)
        grandTotal.Add(grandTotal, bankJoltage)
    }
    log.Printf("Part 2 Total: %s", grandTotal.String())
	}
