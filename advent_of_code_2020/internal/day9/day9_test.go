package day9

import (
	"testing"
)

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestFindWeakness(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"35",
				"20",
				"15",
				"25",
				"47",
				"40",
				"62",
				"55",
				"65",
				"95",
				"102",
				"117",
				"150",
				"182",
				"127",
				"219",
				"299",
				"277",
				"309",
				"576",
			},
			expectedOutput:  127,
			expectedOutput2: 62,
		},
	}

	for _, testCase := range testCases {
		actualOutput, error := FindWeakness(testCase.input, 5)
		if error != nil {
			t.Errorf("Got error: %v", error)
		}

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualOutput2, error := ExploitWeakness(testCase.input, actualOutput)
		if error != nil {
			t.Errorf("Got error: %v", error)
		}

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}
