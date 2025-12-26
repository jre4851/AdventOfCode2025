package puzzles

import (
	"dec11/helpers"
	"encoding/json"
	"fmt"
	// "log"
)

func Part1() {
		deviceMap := helpers.ValidateInput()
		jsonBytes, _ := json.MarshalIndent(deviceMap, "", "  ")
		fmt.Println("Dec 11 - Part 1 - Device Map:")
		fmt.Println(string(jsonBytes))
		
    paths := CountPaths(deviceMap, "you", "out")
		fmt.Printf("Number of unique paths from 'you' to 'out': %d\n", paths)
}

func CountPaths(deviceMap map[string][]string, current, target string) int {
    if current == target {
        return 1
    }
    count := 0
    for _, next := range deviceMap[current] {
        count += CountPaths(deviceMap, next, target)
    }
    return count
}