package day17

import "testing"

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestGetActiveCubes(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				".#.",
				"..#",
				"###",
			},
			expectedOutput:  112,
			expectedOutput2: 848,
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetActiveCubes(testCase.input, 6, false)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := GetActiveCubes(testCase.input, 6, true)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}
