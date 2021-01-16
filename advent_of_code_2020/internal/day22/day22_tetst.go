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
		expectedOutput2 string
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
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetWinningScore(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}