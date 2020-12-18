package day12

import "testing"

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestDetermineManhattanDistance(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			expectedOutput:  25,
			expectedOutput2: 286,
		},
	}

	for _, testCase := range testCases {
		actualOutput := DetermineManhattanDistance(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := DetermineManhattanDistance2(testCase.input)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}
