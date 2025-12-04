package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ValidateInput() ([]string, error) {
    file, err := os.Open("./input/aocInput.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()		
		inputs = append(inputs, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return inputs, nil
}
