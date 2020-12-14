package day9

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day1"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

// FindWeakness finds the first number that isn't the sum of two others in the window size
func FindWeakness(input []string, windowSize int) (int, error) {
	intList := utils.ConvertStringListToIntList(input)
	windowSet := utils.GenerateIntSet(intList[:windowSize])

	for index, intValue := range intList[windowSize:] {
		if _, _, error := day1.FindPairSum(windowSet, intValue); error != nil {
			return intValue, nil
		}
		// Little hacky, using this new lists index to refer to the original list
		delete(windowSet, intList[index])
		windowSet[intValue] = true
	}

	return -1, fmt.Errorf("Didn't find weakness in input. :(")
}

// ExploitWeakness finds a continuous set of ints in the input that add up to the input weakness int
// and sums the max and min in the set
func ExploitWeakness(input []string, weakness int) (int, error) {
	intList := utils.ConvertStringListToIntList(input)

	startIndex := 0
	endIndex := 0
	sum := intList[0]

	for endIndex < len(intList) {
		if sum == weakness && endIndex > startIndex {
			return sumMinAndMax(intList[startIndex : endIndex+1]), nil
		} else if sum > weakness {
			if startIndex < endIndex {
				sum = sum - intList[startIndex]
				startIndex++
			} else if startIndex == endIndex {
				startIndex++
				endIndex++
				sum = intList[startIndex]
			}
		} else {
			endIndex++
			sum += intList[endIndex]
		}
	}

	return -1, fmt.Errorf("Couldn't exploit weakness in input. :(")
}

func sumMinAndMax(intList []int) int {
	min := intList[0]
	max := intList[0]

	for _, intValue := range intList[1:] {
		if intValue > max {
			max = intValue
		} else if intValue < min {
			min = intValue
		}
	}

	return min + max
}
