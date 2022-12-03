package day3

import (
	"fmt"
)

func Solve1(lines []string) (int, error) {
	pointsMap := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	sum := 0

	for _, line := range lines {
		points, ok := pointsMap[line]
		if !ok {
			return 0, fmt.Errorf("Line not in pointsMap: %v", line)
		}

		sum += points
	}

	return sum, nil
}

func Solve2(lines []string) (int, error) {

	pointsMap := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	sum := 0

	for _, line := range lines {
		points, ok := pointsMap[line]
		if !ok {
			return 0, fmt.Errorf("Line not in pointsMap: %v", line)
		}

		sum += points
	}

	return sum, nil
}
