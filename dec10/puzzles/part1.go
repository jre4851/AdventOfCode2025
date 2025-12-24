package puzzles

import (
	"dec10/helpers"
)

func minButtonPresses(target []int, buttons [][]int, nLights int) int {
	m := len(buttons)
	n := nLights
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	for j, btn := range buttons {
		for _, light := range btn {
			matrix[light][j] = 1
		}
	}
	b := make([]int, n)
	copy(b, target)
	row := 0
	pivotCol := make([]int, n)
	for i := range pivotCol {
		pivotCol[i] = -1
	}
	for col := 0; col < m && row < n; col++ {
		pivot := -1
		for i := row; i < n; i++ {
			if matrix[i][col] == 1 {
				pivot = i
				break
			}
		}
		if pivot == -1 {
			continue
		}
		matrix[row], matrix[pivot] = matrix[pivot], matrix[row]
		b[row], b[pivot] = b[pivot], b[row]
		pivotCol[row] = col
		for i := 0; i < n; i++ {
			if i != row && matrix[i][col] == 1 {
				for j := col; j < m; j++ {
					matrix[i][j] ^= matrix[row][j]
				}
				b[i] ^= b[row]
			}
		}
		row++
	}
	for i := row; i < n; i++ {
		if b[i] != 0 {
			return -1
		}
	}
	freeCols := []int{}
	used := make([]bool, m)
	for i := 0; i < row; i++ {
		if pivotCol[i] != -1 {
			used[pivotCol[i]] = true
		}
	}
	for i := 0; i < m; i++ {
		if !used[i] {
			freeCols = append(freeCols, i)
		}
	}
	minPresses := -1
	total := 1 << len(freeCols)
	for mask := 0; mask < total; mask++ {
		x := make([]int, m)
		for i, col := range freeCols {
			if (mask>>i)&1 == 1 {
				x[col] = 1
			}
		}
		for i := row - 1; i >= 0; i-- {
			sum := b[i]
			for j := pivotCol[i] + 1; j < m; j++ {
				if matrix[i][j] == 1 {
					sum ^= x[j]
				}
			}
			x[pivotCol[i]] = sum
		}
		presses := 0
		for i := 0; i < m; i++ {
			presses += x[i]
		}
		if minPresses == -1 || presses < minPresses {
			minPresses = presses
		}
	}
	return minPresses
}

func Part1() {
	machines := helpers.ValidateInput()
	totalPresses := 0
	for _, m := range machines {
		// Convert light diagram to target []int
		target := make([]int, len(m.LightDiagram))
		for i, c := range m.LightDiagram {
			if c == '#' {
				target[i] = 1
			} else {
				target[i] = 0
			}
		}
		presses := minButtonPresses(target, m.WiringSchematic, len(m.LightDiagram))
		totalPresses += presses
	}
	println("Fewest total button presses:", totalPresses)
}
