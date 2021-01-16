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
	testCases := []struct {
		input           []string
		expectedOutput  int
		expectedOutput2 string
	}{
		{
			input: []string{
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			expectedOutput:  5,
			expectedOutput2: "mxmxvkd,sqjhc,fvjkl",
		},
	}

	for _, testCase := range testCases {
		actualOutput := CountNoAllergens(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := GetDangerousIngredientList(testCase.input)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		}
	}
}