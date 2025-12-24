package helpers

import (
	"bufio"
	"os"
	"strings"
)

func ValidateInput() (map[string][]string) {
    file, _ := os.Open("./input/testInput.txt")
		defer file.Close()

		deviceMap := make(map[string][]string)
		scanner := bufio.NewScanner(file)
		
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue //malformed line
			} 

			device := strings.TrimSpace(parts[0])
			outputs := strings.Fields(parts[1])

			deviceMap[device] = outputs
		}

		return deviceMap
}
