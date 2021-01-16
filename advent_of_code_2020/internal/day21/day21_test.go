package day21

import (
	"testing"
)

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestCountNoAllergens(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			expectedOutput: 5,
		},
	}

	for _, testCase := range testCases {
		actualOutput := CountNoAllergens(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}