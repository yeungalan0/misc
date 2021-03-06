package day21

import (
	"log"
	"sort"
	"strings"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

type food struct {
	ingredients []string
	allergens []string
}

// GetDangerousIngredientList returns an list of ingredients arranged alphabetically according to the allergen
func GetDangerousIngredientList(input []string) string {
	dangerousIngredientList := ""
	foods := parseInput(input)

	allergensToIngredients := determineIngredientAllergens(foods)
	// Sanity check
	for allergen, ingredients := range allergensToIngredients {
		if len(ingredients) != 1 {
			log.Fatalf("Allergen still has multiple ingredients! %v: %v", allergen, ingredients)
		}
	}

	i := 0
	keys := make([]string, len(allergensToIngredients))
	for key := range allergensToIngredients {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	for _, key := range keys {
		dangerousIngredientList += allergensToIngredients[key][0] + ","
	}

	return dangerousIngredientList[:len(dangerousIngredientList)-1]
}

// CountNoAllergens counts the number of incgredients listed (with repeats) that don't have an allergen
func CountNoAllergens(input []string) int {
	foods := parseInput(input)

	allergensToIngredients := determineIngredientAllergens(foods)
	ingredientsSet := convertToIngredientsSet(allergensToIngredients)

	// convert map to set of ingredients with possible allergens
	// for all ingredients of foods, if ingredient not in set count
	noAllergenCount := 0

	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if !ingredientsSet[ingredient] {
				noAllergenCount++
			}
		}
	}

	return noAllergenCount
}

func convertToIngredientsSet(allergensToIngredients map[string][]string) map[string]bool {
	ingredientsSet := map[string]bool{}

	for _, ingredients := range allergensToIngredients {
		for _, ingredient := range ingredients {
			ingredientsSet[ingredient] = true
		}
	}

	return ingredientsSet
}

func determineIngredientAllergens(foods []food) map[string][]string {
	allergensToIngredients := map[string][]string{}

	for _, food := range foods {
		for _, allergen := range food.allergens {
			if _, isPresent := allergensToIngredients[allergen]; isPresent {
				allergensToIngredients[allergen] = intersection(allergensToIngredients[allergen], food.ingredients)
			} else {
				ingredientsCopy := make([]string, len(food.ingredients))
				copy(ingredientsCopy, food.ingredients)
				allergensToIngredients[allergen] = ingredientsCopy
			}

			allergensToIngredients = narrowPossibleIngredients(allergen, allergensToIngredients)
		}
	}

	return allergensToIngredients
}

func narrowPossibleIngredients(allergen string, allergensToIngredients map[string][]string) map[string][]string {
	if len(allergensToIngredients[allergen]) > 1 {
		return allergensToIngredients
	}

	queue := []string{allergen}
	processed := []string{}

	for len(queue) > 0 {
		determinedAllergen := queue[0]
		ingredient := allergensToIngredients[determinedAllergen][0]
		for currAllergen, ingredients := range allergensToIngredients {
			if currAllergen != determinedAllergen {
				allergensToIngredients[currAllergen] = utils.DeleteElement(ingredient, ingredients)
			}
		}

		for currAllergen := range allergensToIngredients {
			if len(allergensToIngredients[currAllergen]) == 1 && !utils.Contains(len(processed), func(i int) bool {return processed[i] == currAllergen}) {
				queue = append(queue, currAllergen)
			}
		}
		processed = append(processed, determinedAllergen)
		queue = queue[1:]
	}

	return allergensToIngredients
}

func intersection(stringSlice1 []string, stringSlice2 []string) []string {
	intersection := []string{}
	for _, s := range stringSlice1 {
		if utils.Contains(len(stringSlice2), func(i int) bool {return stringSlice2[i] == s}) {
			intersection = append(intersection, s)
		}
	}

	return intersection
}

func parseInput(input []string) []food {
	foods := []food{}

	for _, line := range input {
		ingredientsAndAllergens := strings.Split(line, " (contains ")
		ingredients := strings.Split(ingredientsAndAllergens[0], " ")
		allergensString := ingredientsAndAllergens[1]
		allergensString = allergensString[:len(allergensString)-1]
		allergens := strings.Split(allergensString, ", ")

		foods = append(foods, food{ingredients: ingredients, allergens: allergens})
	}

	return foods
}