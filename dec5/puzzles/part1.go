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

		// for _, r := range ranges {
		// 	for i := r.StartingValue; i <= r.EndingValue; i++ {
		// 		if helpers.Contains(ingredients, i) {
		// 			if !helpers.Contains(validIngredients, i) {
		// 				log.Printf("Ingredient %d found in range %+v", i, r)
		// 				validIngredients = append(validIngredients, i)
		// 			}
		// 		}
		// 	}
		// }

		log.Printf("Part 1 Valid Ingredients: %d", len(validIngredients))
}
