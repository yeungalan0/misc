package day23

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetCupLabels(t *testing.T) {
	testCases := []struct {
		input           []string
		rounds			int
		cupNumber int
		expectedOutput  string
		expectedOutput2 string
	}{
		// {
		// 	input: []string{
		// 		"389125467",
		// 	},
		// 	rounds: 10,
		// 	cupNumber: 9,
		// 	expectedOutput:  "92658374",
		// },
		// {
		// 	input: []string{
		// 		"389125467",
		// 	},
		// 	rounds: 100,
		// 	cupNumber: 9,
		// 	expectedOutput:  "67384529",
		// },
		{
			input: []string{
				"389125467",
			},
			rounds: 1000,
			cupNumber: 1000000,
			expectedOutput:  "149245887792",
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetCupLabels(testCase.input, testCase.rounds, testCase.cupNumber)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		// actualOutput2 := GetWinningScore(testCase.input, true)

		// if testCase.expectedOutput2 != actualOutput2 {
		// 	t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		// }
	}
}

func TestSimulateCrabCups(t *testing.T) {
	testCases := []struct {
		cups        []int
		rounds int
		expectedOutput []int
	}{
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutput: []int{3, 2, 8, 9, 1, 5, 4, 6, 7},
			rounds: 1,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutput: []int{2, 5, 4, 6, 7, 8, 9, 1, 3},
			rounds: 2,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutput: []int{5, 8, 9, 1, 3, 4, 6, 7, 2},
			rounds: 3,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutput: []int{8, 4, 6, 7, 9, 1, 3, 2, 5},
			rounds: 4,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutput: []int{4, 1, 3, 6, 7, 9, 2, 5, 8},
			rounds: 5,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutput: []int{1, 9, 3, 6, 7, 2, 5, 8, 4},
			rounds: 6,
		},
	}

	for _, testCase := range testCases {
		actualOutput := simulateCrabCups(testCase.cups, testCase.rounds)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestGetNewCurrCupIndex(t *testing.T) {
	testCases := []struct {
		currCupVal     int
		newCups        []int
		expectedOutput int
	}{
		{
			currCupVal:     4,
			newCups:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedOutput: 4,
		},
		{
			currCupVal:     2,
			newCups:        []int{5, 8, 3, 7, 4, 1, 9, 2, 6},
			expectedOutput: 8,
		},
		{
			currCupVal:     6,
			newCups:        []int{8, 2, 5, 7, 4, 1, 9, 3, 6},
			expectedOutput: 0,
		},
	}

	for _, testCase := range testCases {
		actualOutput := getNewCurrCupIndex(testCase.currCupVal, testCase.newCups)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestGetDestinationCupIndex(t *testing.T) {
	testCases := []struct {
		currCupIndex    int
		closeCupIndexes []int
		cups            []int
		expectedOutput  int
	}{
		{
			currCupIndex:    4,
			closeCupIndexes: []int{5, 6, 7},
			cups:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedOutput:  3,
		},
		{
			currCupIndex:    2,
			closeCupIndexes: []int{3, 4, 5},
			cups:            []int{5, 8, 3, 7, 4, 1, 9, 2, 6},
			expectedOutput:  7,
		},
		{
			currCupIndex:    6,
			closeCupIndexes: []int{7, 8, 0},
			cups:            []int{8, 2, 5, 7, 4, 1, 9, 3, 6},
			expectedOutput:  3,
		},
		{
			currCupIndex:    1,
			closeCupIndexes: []int{2, 3, 4},
			cups:            []int{5, 8, 9, 1, 3, 4, 6, 7, 2},
			expectedOutput:  7,
		},
	}

	for _, testCase := range testCases {
		actualOutput := getDestinationCupIndex(testCase.currCupIndex, testCase.closeCupIndexes, testCase.cups)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestGetCircularIndexes(t *testing.T) {
	testCases := []struct {
		startIndex     int
		endIndex       int
		expectedOutput []int
	}{
		{
			startIndex:     0,
			endIndex:       4,
			expectedOutput: []int{0, 1, 2, 3},
		},
		{
			startIndex:     8,
			endIndex:       6,
			expectedOutput: []int{8, 0, 1, 2, 3, 4, 5},
		},
		{
			startIndex:     7,
			endIndex:       2,
			expectedOutput: []int{7, 8, 0, 1},
		},
	}

	for _, testCase := range testCases {
		actualOutput := getCupsWithCircularIndexes(testCase.startIndex, testCase.endIndex, 9)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}
