package day10

import (
	"log"
	"sort"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

// EfficientAdapterOrder returns the 1 and 3 jolt differences multiplied when all adapters are used
func EfficientAdapterOrder(input []string) int {
	adapters := generateSortedAdapterListWithOutletAndDevice(input)

	joltageDifferenceMap := generateJoltageDifferenceMap(adapters)

	return joltageDifferenceMap[1] * joltageDifferenceMap[3]
}

// generateJoltageDifferenceMap generates a map of differences in joltage as a key and the count as the value
func generateJoltageDifferenceMap(adapters []int) map[int]int {
	joltageDifferenceMap := map[int]int{}

	for index, adapter := range adapters[1:] {
		// Hacky, using index in range list to refer to original list to skip the first element in for loop
		joltageDifference := adapter - adapters[index]
		if joltageDifference <= 3 {
			joltageDifferenceMap[joltageDifference]++
		} else {
			log.Fatalf("Got unexpected joltage difference of %v! currrent joltage: %v, previous joltage: %v\n",
				joltageDifference, adapter, adapters[index])
		}
	}

	return joltageDifferenceMap
}

// PossibleAdapterCombinations returns the number of possible adapter combinations from the outlet to your device
// Assumes input list has a valid adapter order
func PossibleAdapterCombinations(input []string, maxJoltageDifference int) int {
	adapters := generateSortedAdapterListWithOutletAndDevice(input)
	combinations := 1
	currentIndex := 0

	// Basic idea is an independent subset is surrounded by maxJoltage difference on both sides, and multiplying
	// the solutions for each independent subset yields the solution for the whole set. For instance when
	// maxJoltage difference is 3, in the joltage set: 0, 3, 4, 5, 7, 8, 11, 12, 14, 15, 18..., the possible solutions
	// for 0, 3, 4, 5, 7, 8, 11 can be independently determined and multiplied with the possible combinations of
	// for 11, 12, 14, 15, 18, to yield a solution for both sets, and this process continues on for all subsets
	for currentIndex < len(adapters)-1 {
		groupEndIndex := findNextMaxJoltageDifference(adapters, currentIndex, maxJoltageDifference)
		combinations *= getValidCombinations(adapters[currentIndex:groupEndIndex+1], maxJoltageDifference)
		currentIndex = groupEndIndex
	}

	return combinations
}

func findNextMaxJoltageDifference(adapters []int, index int, maxJoltageDifference int) int {
	for currentIndex := index + 1; currentIndex < len(adapters); currentIndex++ {
		if adapters[currentIndex]-adapters[currentIndex-1] == maxJoltageDifference {
			return currentIndex
		}
	}

	return -1
}

// ValidCombinations Credit for algorithm: https://www.reddit.com/r/adventofcode/comments/ka8z8x/2020_day_10_solutions/gfcxuxf?utm_source=share&utm_medium=web2x&context=3
// I developed my own solution initially (PossibleAdapterCombinations), but the one at this
// source is far more elegant, and runs in about the same time as my own solution,
// so I thought it'd be fun to implement! This is basically the inverse of my getValidCombinationsHelper
// function, but since we're building all possible solutions from a set of possible next Joltages is
// FAR more efficient than using getValidCombinationsHelper to solve the whole problem
func ValidCombinations(input []string, maxJoltageDifference int) int {
	adapters := generateSortedAdapterListWithOutletAndDevice(input)
	adaptersSet := utils.GenerateIntSet(adapters)

	// Adapter to number of solutions up to that point map
	solutionsMap := map[int]int{}
	solutionsMap[0] = 1

	for _, adapter := range adapters {
		// Going from lowest to highest joltage here is vital to the algorithm
		for difference := 1; difference <= maxJoltageDifference; difference++ {
			nextPossibleAdapter := adapter + difference
			if adaptersSet[nextPossibleAdapter] {
				solutionsMap[nextPossibleAdapter] += solutionsMap[adapter]
			}
		}
	}

	return solutionsMap[adapters[len(adapters)-1]]
}

func getValidCombinations(adapters []int, maxJoltageDifference int) int {
	startingIndex := 1

	return getValidCombinationsHelper(adapters, startingIndex, maxJoltageDifference)
}

func getValidCombinationsHelper(adapters []int, index int, maxJoltageDifference int) int {
	if index >= len(adapters)-2 {
		return 1
	}

	validCombinations := getValidCombinationsHelper(adapters, index+1, maxJoltageDifference)
	adaptersWithIndexRemoved := removeIndex(adapters, index)
	if isValidCombination(adaptersWithIndexRemoved, maxJoltageDifference) {
		validCombinations += getValidCombinationsHelper(adaptersWithIndexRemoved, index, maxJoltageDifference)
	}

	return validCombinations
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isValidCombination(adapters []int, maxJoltageDifference int) bool {
	for i := len(adapters) - 1; i > 0; i-- {
		if adapters[i]-adapters[i-1] > maxJoltageDifference {
			return false
		}
	}
	return true
}

func generateSortedAdapterListWithOutletAndDevice(input []string) []int {
	adapters := utils.ConvertStringListToIntList(input)
	// Add charing addapter effective joltage
	adapters = append(adapters, 0)

	sort.Ints(adapters)
	// Add laptops built in adapter's effective joltage
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	return adapters
}
