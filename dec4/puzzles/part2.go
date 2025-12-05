package puzzles

import (
	"dec4/helpers"
	"log"
)

func Part2() {
	floorPlan, err := helpers.ValidateInput()
	if err != nil {
		log.Fatal(err)
	}

	neighbors := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	var optimizedCount int
	var moarPaper bool = true
	for moarPaper {
		moarPaper = false
		for i, current := range floorPlan {
			if current.Value == "@" {
				atCount := 0
				for _, n := range neighbors {
					nx, ny := current.PositionX+n[0], current.PositionY+n[1]
					for _, check := range floorPlan {
						if check.PositionX == nx && check.PositionY == ny && check.Value == "@" {
							atCount++
							break
						}
					}				
				}
				if atCount < 4 {						
						floorPlan[i].Value = "."
						moarPaper = true
						optimizedCount++
				}
			}
		}
	}

	log.Printf("Part 2 - Optimized Rolls: %d", optimizedCount)
}
