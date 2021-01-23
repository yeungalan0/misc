package day23

import (
	"fmt"
	"log"
	"strings"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

// GetCupLabels returns the labels (in clockwise order) except 1 of the cups after
// N rounds of crab cups
func GetCupLabels(input []string, rounds int, numberOfCups int) string {
	cups := utils.ConvertStringListToIntList(strings.Split(input[0], ""))

	if numberOfCups > len(cups) {
		newCups := make([]int, numberOfCups)
		
		for i := 1; i <= numberOfCups; i++ {
			if i <= len(cups) {
				newCups[i-1] = cups[i-1]
			}
			newCups[i-1] = i
		}

		cups = newCups
	}

	cupsAfterPlay := simulateCrabCups(cups, rounds)
	indexOf1 := utils.SliceIndex(len(cupsAfterPlay), func(i int) bool { return cupsAfterPlay[i] == 1})

	if numberOfCups > len(cups) {
		return fmt.Sprint(cupsAfterPlay[indexOf1 + 1] * cupsAfterPlay[indexOf1 + 2])
	}

	cupIndexesAfter1 := getCupsWithCircularIndexes(indexOf1 + 1, indexOf1 + numberOfCups, numberOfCups)
	cupValsAfter1 := getValuesFromIndexes(cupsAfterPlay, cupIndexesAfter1)

	labels := ""

	for _, val := range cupValsAfter1 {
		labels += fmt.Sprint(val)
	}

	return labels
}

func simulateCrabCups(cups []int, rounds int) []int {
	numberOfCups := len(cups)
	newCups := make([]int, numberOfCups)
	currCupIndex := 0

	for rounds > 0 {
		offset := 0
		closeCupIndexes := getCupsWithCircularIndexes(currCupIndex + 1, currCupIndex + 4, numberOfCups)
		destinationCupIndex := getDestinationCupIndex(currCupIndex, closeCupIndexes, cups)

		newCups[0] = cups[currCupIndex]
		offset++
		newCups = fillSliceIndexes(offset, closeCupIndexes[2] + 1, destinationCupIndex + 1, cups, newCups)
		offset += circularIndexLength(closeCupIndexes[2] + 1, destinationCupIndex + 1, numberOfCups)
		newCups = fillSliceIndexes(offset, currCupIndex + 1, currCupIndex + 4, cups, newCups)
		offset += circularIndexLength(currCupIndex + 1, currCupIndex + 4, numberOfCups)
		newCups = fillSliceIndexes(offset, destinationCupIndex + 1, currCupIndex, cups, newCups)

		currCupIndex = getNewCurrCupIndex(cups[currCupIndex], newCups)
		cups, newCups = newCups, cups
		rounds--
	}

	return cups
}

func circularIndexLength(startIndex, endIndex, numberOfCups int) int {
	startIndex = startIndex % numberOfCups
	endIndex = endIndex % numberOfCups

	if endIndex < startIndex {
		endIndex += numberOfCups
	}

	return endIndex - startIndex
}

func fillSliceIndexes(offset, oldSliceStartIndex, oldSliceEndIndex int, oldSlice, newSlice []int) []int {
	sliceLen := len(oldSlice)
	oldSliceStartIndex = oldSliceStartIndex % sliceLen
	oldSliceEndIndex = oldSliceEndIndex % sliceLen

	if oldSliceEndIndex < oldSliceStartIndex {
		oldSliceEndIndex += sliceLen
	}

	for i := oldSliceStartIndex; i < oldSliceEndIndex; i++ {
		newSlice[(offset + i - oldSliceStartIndex) % sliceLen] = oldSlice[i % sliceLen]
	}

	return newSlice
}

func getValuesFromIndexes(slice, indexes []int) []int {
	values := []int{}
	for _, i := range indexes {
		values = append(values, slice[i])
	}

	return values
}

func getNewCurrCupIndex(currCupVal int, newCups []int) int {
	numberOfCups := len(newCups)
	currCupIndex := utils.SliceIndex(numberOfCups, func(i int) bool {return newCups[i] == currCupVal})
	newCurrCupIndex := (currCupIndex + 1) % numberOfCups
	return newCurrCupIndex
}

func getDestinationCupIndex(currCupIndex int, closeCupIndexes, cups []int) int {
	currCupVal := cups[currCupIndex]
	destinationIndex := -1

    for destinationIndex == -1 {
		currCupVal--
		if currCupVal < 0 {
			currCupVal = 9
		}

		possibleIndex := utils.SliceIndex(len(cups), func(i int) bool {return cups[i] == currCupVal})
		isInCloseCupIndexes := utils.Contains2(len(closeCupIndexes), func(i int) bool {return closeCupIndexes[i] == possibleIndex})

		if possibleIndex == currCupIndex {
			log.Fatalf("Looping detected! possbileIndex: %v, currCupIndex: %v", possibleIndex, currCupIndex)
		}

		if !isInCloseCupIndexes {
			destinationIndex = possibleIndex
		}
	}

	return destinationIndex
}

// getCupsWithCircularIndexes returns the indexes in cirular fashion inclusive of start and exclusive of end
func getCupsWithCircularIndexes(startIndex, endIndex, numberOfCups int) []int {
	indexes := []int{}
	startIndex = startIndex % numberOfCups
	endIndex = endIndex % numberOfCups

	if endIndex < startIndex {
		endIndex += 9
	}

	for i := startIndex; i < endIndex; i++ {
		indexes = append(indexes, i % numberOfCups)
	}

	return indexes
}