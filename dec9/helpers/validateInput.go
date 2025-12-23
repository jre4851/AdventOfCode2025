package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type GridCoords struct {
	X int
	Y int
}

func ValidateInput() []GridCoords {
       file, err := os.Open("./input/aocInput.txt")
       if err != nil {
	       return nil
       }
       defer file.Close()

       var redSquares []GridCoords
       scanner := bufio.NewScanner(file)
       for scanner.Scan() {
	       line := scanner.Text()
	       parts := strings.Split(line, ",")
	       if len(parts) < 2 {
		       continue
	       }
	       xVal, err1 := strconv.Atoi(parts[0])
	       yVal, err2 := strconv.Atoi(parts[1])
	       if err1 != nil || err2 != nil {
		       continue
	       }
	       redSquares = append(redSquares, GridCoords{X: xVal, Y: yVal})
       }
       return redSquares
}
