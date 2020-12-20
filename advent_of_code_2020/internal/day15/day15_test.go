package day15

import "testing"

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestGetRecitationAtTurn(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"0,3,6",
			},
			expectedOutput:  436,
			expectedOutput2: 175594,
		},
		{
			input: []string{
				"1,3,2",
			},
			expectedOutput:  1,
			expectedOutput2: 2578,
		},
		{
			input: []string{
				"2,1,3",
			},
			expectedOutput:  10,
			expectedOutput2: 3544142,
		},
		{
			input: []string{
				"1,2,3",
			},
			expectedOutput:  27,
			expectedOutput2: 261214,
		},
		{
			input: []string{
				"2,3,1",
			},
			expectedOutput:  78,
			expectedOutput2: 6895259,
		},
		{
			input: []string{
				"3,2,1",
			},
			expectedOutput:  438,
			expectedOutput2: 18,
		},
		{
			input: []string{
				"3,1,2",
			},
			expectedOutput:  1836,
			expectedOutput2: 362,
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetRecitationAtTurn(testCase.input, 2020)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := GetRecitationAtTurn(testCase.input, 30000000)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}
