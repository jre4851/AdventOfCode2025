package puzzles

import (
	"dec8/helpers"
	"log"
)

func Part1() {
		boxes := helpers.ValidateInput()
		
    log.Println("Part 1 - Junction Boxes Loaded:", len(boxes))
		for _, box := range boxes {
			log.Printf("Box - x: %d, y: %d, z: %d", box.XAxis, box.YAxis, box.ZAxis)
		}
}
