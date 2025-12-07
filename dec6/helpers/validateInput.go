package helpers

import (
	"bufio"
	"os"
)

func ValidateInput() ([]string, error) {
	file, err := os.Open("./input/aocInput.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var homework []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		homework = append(homework, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return homework, nil
}

func ValidatePart2Input() ([][]rune, error) {
	file, err := os.Open("./input/aocInput.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}
	return grid, nil
}

