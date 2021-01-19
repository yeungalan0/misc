package day22

import (
	"testing"
)

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestGetWinningScore(t *testing.T) {
	testCases := []struct {
		input           []string
		expectedOutput  int
		expectedOutput2 int
	}{
		{
			input: []string{
				"Player 1:",
				"9",
				"2",
				"6",
				"3",
				"1",
				"",
				"Player 2:",
				"5",
				"8",
				"4",
				"7",
				"10",
			},
			expectedOutput:  306,
			expectedOutput2: 291,
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetWinningScore(testCase.input, false)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := GetWinningScore(testCase.input, true)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		}
	}
}