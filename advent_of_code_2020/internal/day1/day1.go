package day1

import (
	"errors"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

// FindTripleSum find the three integers in the given list that add up to sum
func FindTripleSum(list []int, sum int) (int, int, int, error) {
	for index, element := range list {
		remainder := sum - element
		set := utils.GenerateIntSet(list[index+1:])
		int1, int2, err := FindPairSum(set, remainder)
		if err == nil {
			return element, int1, int2, nil
		}
	}
	return -1, -1, -1, errors.New("No pair found that adds to sum :(")
}

// FindPairSum find the two integers in the given list that add up to sum
func FindPairSum(set map[int]bool, sum int) (int, int, error) {
	for element := range set {
		compliment := sum - element
		if set[compliment] {
			return element, compliment, nil
		}
	}

	return -1, -1, errors.New("No pair found that adds to sum :(")
}
