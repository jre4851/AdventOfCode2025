package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ManualInfo struct {
	LightDiagram string
	WiringSchematic [][]int
	JoltageRequirements []int
}

func ValidateInput() []ManualInfo {
	file, _ := os.Open("./input/aocInput.txt")
	defer file.Close()

	var manual []ManualInfo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		segments := splitSegments(line)
		if len(segments) == 0 {
			continue
		}

		lightDiagram := segments[0][1 : len(segments[0])-1]

		wiringSchematic := make([][]int, 0)
		joltageRequirements := make([]int, 0)

		for i := 1; i < len(segments); i++ {
			seg := segments[i]
			if len(seg) == 0 {
				continue
			}
			switch seg[0] {
			case '(': 
				nums := parseInts(seg[1 : len(seg)-1])
				wiringSchematic = append(wiringSchematic, nums)
			case '{': 
				joltageRequirements = parseInts(seg[1 : len(seg)-1])
			}
		}

		manual = append(manual, ManualInfo{
			LightDiagram:        lightDiagram,
			WiringSchematic:     wiringSchematic,
			JoltageRequirements: joltageRequirements,
		})
	}
	return manual
}

func parseInts(s string) []int {
	parts := strings.Split(s, ",")
	var res []int
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			res = append(res, n)
		}
	}
	return res
}

func splitSegments(line string) []string {
	var segments []string
	var start int
	var inSegment bool
	var open, close byte

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '(', '[', '{':
			inSegment = true
			open = line[i]
			switch open {
			case '(':
				close = ')'
			case '[':
				close = ']'
			case '{':
				close = '}'
			}
			start = i
		case ')', ']', '}':
			if inSegment && line[i] == close {
				segments = append(segments, line[start:i+1])
				inSegment = false
			}
		}
	}
	return segments
}
