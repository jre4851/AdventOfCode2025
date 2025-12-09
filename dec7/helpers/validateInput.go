package helpers

import (
	"bufio"
	"os"
)

func ValidateInput() [][]rune {
    file, _ := os.Open("./input/aocInput.txt")
    defer file.Close()

    var grid [][]rune
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := []rune(scanner.Text())
        grid = append(grid, line)
    }
    return grid
}

func FindBeamStart(grid [][]rune) (int, int) {
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 'S' {
                return r, c
            }
        }
    }
    return -1, -1
}