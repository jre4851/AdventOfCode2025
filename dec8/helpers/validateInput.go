package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type JunctionBox struct {
	XAxis, YAxis, ZAxis int
}

func ValidateInput() ([]JunctionBox){
    file, _ := os.Open("./input/aocInput.txt")
		defer file.Close()

		var boxes []JunctionBox
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			boxes = append(boxes, JunctionBox{
				XAxis: x,
				YAxis: y,
				ZAxis: z,
			})
		}
		return boxes
}
