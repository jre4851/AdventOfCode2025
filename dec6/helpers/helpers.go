package helpers

import (
	"unicode"
)

func IsSpaceColumn(grid [][]rune, col int) bool {
    for row := range grid {
        if !unicode.IsSpace(grid[row][col]) {
            return false
        }
    }
    return true
}

func EvaluateProblem(grid [][]rune, left, right int) int {
    height := len(grid)
    // Operator is in the bottom row, rightmost column of the block
    op := grid[height-1][right]

    var numbers []int
    // Process columns right-to-left
    for c := right; c >= left; c-- {
        val := 0
        for r := 0; r < height-1; r++ {
            ch := grid[r][c]
            if unicode.IsDigit(ch) {
                val = val*10 + int(ch-'0')
            }
        }
        numbers = append(numbers, val)
    }

    result := numbers[0]
    for _, n := range numbers[1:] {
        if op == '+' {
            result += n
        } else {
            result *= n
        }
    }
    return result
}