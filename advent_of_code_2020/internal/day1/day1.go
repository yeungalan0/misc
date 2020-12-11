package day1

import (
	"errors"
)

// FindTripleSum find the three integers in the given list that add up to sum
func FindTripleSum(list []int, sum int) (int, int, int, error) {
	for index, element := range list {
		remainder := sum - element
		set := GenerateIntSet(list[index+1:])
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

// GenerateIntSet generates a map from a list of ints
func GenerateIntSet(list []int) map[int]bool {
	set := make(map[int]bool)
	for _, integer := range list {
		set[integer] = true
	}
	return set
}
