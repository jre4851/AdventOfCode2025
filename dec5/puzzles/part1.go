package puzzles

import (
	"dec5/helpers"
	"log"
)

func Part1() {
		ranges, ingredients, err := helpers.ValidateInput()
		if err != nil {
			log.Fatalf("Failed to validate input: %v", err)
		}

		var validIngredients []int

		for _, ingredient := range ingredients {
			for _, r := range ranges {
				if ingredient >= r.StartingValue && ingredient <= r.EndingValue {
					if !helpers.Contains(validIngredients, ingredient) {
						log.Printf("Ingredient %d found in range %+v", ingredient, r)
						validIngredients = append(validIngredients, ingredient)
					}
				}
			}
		}

		log.Printf("Part 1 Valid Ingredients: %d", len(validIngredients))
}
