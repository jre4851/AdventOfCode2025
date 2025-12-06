package puzzles

import (
	"dec3/helpers"
	"log"
	"sort"
	"strconv"
)

func Part2() {
    inputs, err := helpers.ValidateInput()
    if err != nil {
        log.Fatal(err)
    }

    grandTotal := 0
    for _, inputStr := range inputs {
        digits := make([]int, len(inputStr))
        for i, ch := range inputStr {
            digits[i] = int(ch - '0')
        }
        sort.Slice(digits, func(i, j int) bool {
            return digits[i] > digits[j]
        })
        // Take up to 12 largest digits
        n := 12
        if len(digits) < 12 {
            n = len(digits)
        }
        largestDigits := digits[:n]
        // Concatenate digits into a string
        digitStr := ""
        for _, d := range largestDigits {
            digitStr += strconv.Itoa(d)
        }
        totalSum, err := strconv.Atoi(digitStr)
        if err != nil {
            log.Fatal(err)
        }
        grandTotal += totalSum
    }
    log.Printf("Part 2 Total: %d", grandTotal)
}
