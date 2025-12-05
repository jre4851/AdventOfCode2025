package helpers

import (
	"bufio"
	"fmt"
	"os"
)

type PaperRoll struct {
	PositionX int
	PositionY int
	Value string
}

func ValidateInput() ([]PaperRoll, error) {
    file, err := os.Open("./input/aocInput.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}
		defer file.Close()

		var floorPlan []PaperRoll
		scanner := bufio.NewScanner(file)
		yIndex := 0
		for scanner.Scan() {
			line := scanner.Text()
			for xIndex, char := range line {
				floorPlan = append(floorPlan, PaperRoll{
					PositionX: xIndex,
					PositionY: yIndex,
					Value: string(char),
				})
			}
			yIndex++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return nil, err
		}

		return floorPlan, nil
}
