package day10

import "testing"

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

type testCaseInt struct {
	input          []int
	expectedOutput int
}

func TestEfficientAdapterOrder(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"16", "10", "15", "5", "1", "11", "7", "19", "6", "12", "4",
			},
			expectedOutput:  35,
			expectedOutput2: 8,
		},
		{
			input: []string{
				"28", "33", "18", "42", "31", "14", "46", "20", "48", "47", "24", "23",
				"49", "45", "19", "38", "39", "11", "1", "32", "25", "35", "8", "17", "7",
				"9", "4", "2", "34", "10", "3",
			},
			expectedOutput:  220,
			expectedOutput2: 19208,
		},
		{
			input:           []string{"1", "2", "3", "4"},
			expectedOutput:  4,
			expectedOutput2: 7,
		},
	}

	for _, testCase := range testCases {
		actualOutput := EfficientAdapterOrder(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := PossibleAdapterCombinations(testCase.input, 3)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}

func TestGetValidCombinations(t *testing.T) {
	testCases := []testCaseInt{
		{
			input:          []int{0, 1, 2, 3, 4, 7},
			expectedOutput: 7,
		},
		{
			input:          []int{0, 2, 3, 6},
			expectedOutput: 2,
		},
		{
			input:          []int{0, 1, 2, 3, 5, 8},
			expectedOutput: 6,
		},
		{
			input:          []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22},
			expectedOutput: 8,
		},
		{
			input: []int{
				0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24,
				25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52,
			},
			expectedOutput: 19208,
		},
	}

	for _, testCase := range testCases {
		actualOutput := getValidCombinations(testCase.input, 3)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}
