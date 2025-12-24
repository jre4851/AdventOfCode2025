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
		
    // log.Println("Part 1 placeholder")
}
