package day18

import "testing"

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestSumEvaluatedExpressions(t *testing.T) {
	testCases := []testCase{
		{
			input:           []string{"1 + (2 * 3) + (4 * (5 + 6))"},
			expectedOutput:  51,
			expectedOutput2: 51,
		},
		{
			input:           []string{"2 * 3 + (4 * 5)"},
			expectedOutput:  26,
			expectedOutput2: 46,
		},
		{
			input:           []string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"},
			expectedOutput:  437,
			expectedOutput2: 1445,
		},
		{
			input:           []string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"},
			expectedOutput:  12240,
			expectedOutput2: 669060,
		},
		{
			input:           []string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"},
			expectedOutput:  13632,
			expectedOutput2: 23340,
		},
	}

	for _, testCase := range testCases {
		actualOutput := SumEvaluatedExpressions(testCase.input, false)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := SumEvaluatedExpressions(testCase.input, true)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		}
	}
}
