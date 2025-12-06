package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeSet struct {
	StartingValue int
	EndingValue   int
}

func ValidateInput() ([]RangeSet, []int, error) {
	file, err := os.Open("./input/aocInput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()

	var ranges []RangeSet
	var ingredientList []int
	ingredientsSection := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ingredientsSection = true
			continue
		}
		if ingredientsSection {
			value, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, err
			}
			ingredientList = append(ingredientList, value)
		} else {
			parts := strings.Split(line, "-")
			if err != nil {
				return nil, nil, err
			}
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			ranges = append(ranges, RangeSet{
				StartingValue: start,
				EndingValue:   end,
			})			
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	return ranges, ingredientList, nil
}
