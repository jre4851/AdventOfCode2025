package puzzles

import (
	"dec9/helpers"
	"log"
)

func Part1() {
	       redSquares := helpers.ValidateInput()
		       maxArea := 0
		       var bestA, bestB helpers.GridCoords
		       n := len(redSquares)
		       for i := 0; i < n; i++ {
			       for j := i + 1; j < n; j++ {
				       a, b := redSquares[i], redSquares[j]
				       if a.X == b.X || a.Y == b.Y {
					       continue
				       }
				       area := (abs(a.X-b.X)+1) * (abs(a.Y-b.Y)+1)
				       if area > maxArea {
					       maxArea = area
					       bestA, bestB = a, b
				       }
			       }
		       }
		       log.Printf("Largest rectangle area: %d, corners: %v and %v\n", maxArea, bestA, bestB)
}

func abs(x int) int {
       if x < 0 {
	       return -x
       }
       return x
}
