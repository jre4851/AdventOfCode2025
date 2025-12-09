package puzzles

import (
	"dec7/helpers"
	"log"
)

type Beam struct {
    row, col   int
    dirRow, dirCol int
}

type state struct {
    r, c, dr, dc int
}

func Part1() {
    grid := helpers.ValidateInput()
    startRow, startCol := helpers.FindBeamStart(grid)
    log.Printf("Beam starts at row %d, column %d", startRow, startCol)

    startBeam := Beam{row: startRow, col: startCol, dirRow: 1, dirCol: 0}
    splits := simulate(grid, startBeam)

    log.Printf("Total beam splits: %d", splits)
}

func simulate(grid [][]rune, start Beam) (int) {
    visited := make(map[state]bool)
    splitCount := 0
    queue := []Beam{start}

    for len(queue) > 0 {
        b := queue[0]
        queue = queue[1:]

        nr, nc := b.row+b.dirRow, b.col+b.dirCol
        if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
            continue
        }

        cell := grid[nr][nc]
        s := state{nr, nc, b.dirRow, b.dirCol}
        if visited[s] {
            continue
        }
        visited[s] = true

        if cell == '^' {
            // Split: stop downward beam, emit left and right beams downward
            splitCount++
            log.Printf("Split at (%d,%d), cell=%c", nr, nc, cell)
            // Left beam
            if nc-1 >= 0 {
                queue = append(queue, Beam{nr, nc-1, 1, 0})
            }
            // Right beam
            if nc+1 < len(grid[0]) {
                queue = append(queue, Beam{nr, nc+1, 1, 0})
            }
            continue
        }

        // Continue downward if not a splitter
        queue = append(queue, Beam{nr, nc, b.dirRow, b.dirCol})
    }

    return splitCount
}
