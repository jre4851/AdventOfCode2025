package puzzles

import (
	"dec6/helpers"
	"fmt"
)

// Part2 loads the worksheet from a hardcoded file and prints the grand total.
func Part2() {
    grid, err := helpers.ValidatePart2Input()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    height := len(grid)
    if height == 0 {
        fmt.Println(0)
        return
    }
    width := len(grid[0])

    grandTotal := 0
    col := width - 1
    for col >= 0 {
        if helpers.IsSpaceColumn(grid, col) {
            col--
            continue
        }
        // group contiguous non-space columns into one problem
        right := col
        for col >= 0 && !helpers.IsSpaceColumn(grid, col) {
            col--
        }
        left := col + 1

        // Debug: print block info
        fmt.Printf("Problem block: columns %d to %d\n", left, right)
        op := grid[len(grid)-1][left]
        fmt.Printf("Operator: %v\n", string(op))
        var numbers []int
        for c := right; c >= left; c-- {
            val := 0
            for r := 0; r < len(grid)-1; r++ {
                ch := grid[r][c]
                if ch >= '0' && ch <= '9' {
                    val = val*10 + int(ch-'0')
                }
            }
            numbers = append(numbers, val)
        }
        fmt.Printf("Numbers: %v\n", numbers)
        // Apply operator
        result := numbers[0]
        for _, n := range numbers[1:] {
            if op == '+' {
                result += n
            } else if op == '*' {
                result *= n
            }
        }
        fmt.Printf("Result for block: %d\n\n", result)
        grandTotal += result
    }
    fmt.Printf("Grand Total: %d\n", grandTotal)
}



