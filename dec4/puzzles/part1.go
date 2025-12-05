package puzzles

import (
	"dec4/helpers"
	"log"
)

func Part1() {
	floorPlan, err := helpers.ValidateInput()
	if err != nil {
		log.Fatal(err)
	}

	neighbors := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	var optimizedRolls []helpers.PaperRoll
	for _, current := range floorPlan {
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
			if (atCount < 4)  {
					optimizedRolls = append(optimizedRolls, current)
				}
		}
	}

	log.Println("Part 1 - Optimized Rolls:", len(optimizedRolls))
}
