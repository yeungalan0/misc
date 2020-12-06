package day3

import (
	"testing"
)

type testCase struct {
	input            []string
	expectedTreesHit map[Slope]int
}

var startingCase = testCase{
	input: []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	},
	expectedTreesHit: map[Slope]int{
		Slope{Right: 1, Down: 1}: 2,
		Slope{Right: 3, Down: 1}: 7,
		Slope{Right: 5, Down: 1}: 3,
		Slope{Right: 7, Down: 1}: 4,
		Slope{Right: 1, Down: 2}: 2,
	},
}

func TestTreesHit(t *testing.T) {
	for testSlope, expectedTreesHit := range startingCase.expectedTreesHit {
		treesHit, error := TreesHit(startingCase.input, testSlope)
		if error != nil {
			t.Errorf("Encountered error getting treesHit! error: %v\n", error)
		}

		if expectedTreesHit != treesHit {
			t.Errorf("expected %v treesHit, but got %v! slope: %v\n", expectedTreesHit, treesHit, testSlope)
		}
	}
}
