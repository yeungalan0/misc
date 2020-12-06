package day1

import (
	"errors"
)

// FindTripleSum find the three integers in the given list that add up to sum
func FindTripleSum(list []int, sum int) (int, int, int, error) {
	for index, element := range list {
		remainder := sum - element
		int1, int2, err := FindPairSum(list[index+1:], remainder)
		if err == nil {
			return element, int1, int2, nil
		}
	}
	return -1, -1, -1, errors.New("No pair found that adds to sum :(")
}

// FindPairSum find the two integers in the given list that add up to sum
func FindPairSum(list []int, sum int) (int, int, error) {
	set := generateMap(list)
	for _, element := range list {
		compliment := sum - element
		if set[compliment] {
			return element, compliment, nil
		}
	}

	return -1, -1, errors.New("No pair found that adds to sum :(")
}

func generateMap(list []int) map[int]bool {
	set := make(map[int]bool)
	for _, integer := range list {
		set[integer] = true
	}
	return set
}
