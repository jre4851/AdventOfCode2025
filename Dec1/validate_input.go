package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Combos struct {
	Direction string
	Value int
}

func validateInput() ([]Combos, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var combos []Combos
	scanner := bufio.NewScanner(file)
	lineNum := 1
	validCount := 0
	invalidCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			fmt.Printf("Line %d too short: '%s'\n", lineNum, line)
			invalidCount++
			lineNum++
			continue
		}
		dir := line[0]
		if dir != 'L' && dir != 'R' {
			fmt.Printf("Line %d invalid direction: '%s'\n", lineNum, line)
			invalidCount++
			lineNum++
			continue
		}
		valStr := line[1:]
		if _, err := strconv.Atoi(valStr); err != nil {
			fmt.Printf("Line %d invalid number: '%s'\n", lineNum, line)
			invalidCount++
			lineNum++
			continue
		}
		validCount++
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Invalid number in line: %s", line)
		}
		combos = append(combos, Combos{
			Direction: string(line[0]),
			Value: value,
		})
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	
	return combos, nil
}
