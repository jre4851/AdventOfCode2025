package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ranges struct {
	StartingValue int
	EndingValue   int
}

func ValidateInput() ([]Ranges, error) {
	file, err := os.Open("./input/testInput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var ranges []Ranges
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, Ranges{
			StartingValue: start,
			EndingValue:   end, 
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ranges, nil
}