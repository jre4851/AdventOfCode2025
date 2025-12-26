package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type GridBlock struct {
    ID   int
    Grid [][]rune
}

func ValidateInput() ([]GridBlock, []string, error) {
    file, err := os.Open("input/aocInput.txt")
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    var (
        blocks      []GridBlock
        summaries   []string
        currentGrid [][]rune
        currentID   int
        inGrid      bool
        scanner     = bufio.NewScanner(file)
    )

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            if inGrid && len(currentGrid) > 0 {
                blocks = append(blocks, GridBlock{ID: currentID, Grid: currentGrid})
                currentGrid = nil
                inGrid = false
            }
            continue
        }
        if strings.HasSuffix(line, ":") && len(line) > 1 && unicode.IsDigit(rune(line[0])) {
            fmt.Sscanf(line, "%d:", &currentID)
            inGrid = true
            currentGrid = nil
        } else if inGrid {
            currentGrid = append(currentGrid, []rune(line))
        } else if strings.Contains(line, "x") && strings.Contains(line, ":") {
            summaries = append(summaries, line)
        }
    }
    // Add last block if file doesn't end with blank line
    if inGrid && len(currentGrid) > 0 {
        blocks = append(blocks, GridBlock{ID: currentID, Grid: currentGrid})
    }
    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }
    return blocks, summaries, nil
}

// --- Shape orientation helpers ---
func rotate(grid [][]rune) [][]rune {
    h, w := len(grid), len(grid[0])
    newGrid := make([][]rune, w)
    for i := range newGrid {
        newGrid[i] = make([]rune, h)
        for j := range newGrid[i] {
            newGrid[i][j] = grid[h-1-j][i]
        }
    }
    return newGrid
}

func flip(grid [][]rune) [][]rune {
    h := len(grid)
    newGrid := make([][]rune, h)
    for i := range grid {
        newGrid[i] = make([]rune, len(grid[i]))
        for j := range grid[i] {
            newGrid[i][j] = grid[i][len(grid[i])-1-j]
        }
    }
    return newGrid
}

func gridsEqual(a, b [][]rune) bool {
    if len(a) != len(b) || len(a[0]) != len(b[0]) {
        return false
    }
    for i := range a {
        for j := range a[i] {
            if a[i][j] != b[i][j] {
                return false
            }
        }
    }
    return true
}

func UniqueOrientations(grid [][]rune) [][][]rune {
    var result [][][]rune
    seen := [][][]rune{}
    for rot := 0; rot < 4; rot++ {
        for flipState := 0; flipState < 2; flipState++ {
            var g [][]rune
            if flipState == 0 {
                g = grid
            } else {
                g = flip(grid)
            }
            for r := 0; r < rot; r++ {
                g = rotate(g)
            }
            // Normalize: trim leading/trailing empty rows/cols
            g = trimGrid(g)
            // Check uniqueness
            unique := true
            for _, s := range seen {
                if gridsEqual(g, s) {
                    unique = false
                    break
                }
            }
            if unique {
                seen = append(seen, g)
                result = append(result, g)
            }
        }
    }
    return result
}

func trimGrid(grid [][]rune) [][]rune {
    // Remove empty rows/cols (all '.') from edges
    top, bottom := 0, len(grid)-1
    left, right := 0, len(grid[0])-1
    // Top
    for top <= bottom {
        empty := true
        for j := left; j <= right; j++ {
            if grid[top][j] == '#' {
                empty = false
                break
            }
        }
        if !empty { break }
        top++
    }
    // Bottom
    for bottom >= top {
        empty := true
        for j := left; j <= right; j++ {
            if grid[bottom][j] == '#' {
                empty = false
                break
            }
        }
        if !empty { break }
        bottom--
    }
    // Left
    for left <= right {
        empty := true
        for i := top; i <= bottom; i++ {
            if grid[i][left] == '#' {
                empty = false
                break
            }
        }
        if !empty { break }
        left++
    }
    // Right
    for right >= left {
        empty := true
        for i := top; i <= bottom; i++ {
            if grid[i][right] == '#' {
                empty = false
                break
            }
        }
        if !empty { break }
        right--
    }
    // Build trimmed grid
    h := bottom - top + 1
    w := right - left + 1
    trimmed := make([][]rune, h)
    for i := 0; i < h; i++ {
        trimmed[i] = make([]rune, w)
        for j := 0; j < w; j++ {
            trimmed[i][j] = grid[top+i][left+j]
        }
    }
    return trimmed
}
