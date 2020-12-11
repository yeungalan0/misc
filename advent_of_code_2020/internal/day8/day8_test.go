package day8

import (
	"testing"
)

type testCase struct {
	input                    []string
	expectedOutput           int
	expectedTerminatedOutput int
}

func TestGetAccumulatorAtLoop(t *testing.T) {
	var testCases = []testCase{
		{
			input: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			expectedOutput:           5,
			expectedTerminatedOutput: 8,
		},
	}

	for _, testCase := range testCases {
		operationList := ParseOperations(testCase.input)
		actualOutput, looped, _ := GetAccumulatorAtLoop(operationList)

		if testCase.expectedOutput != actualOutput && looped == false {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualTerminatedOutput := GetAccumulatorAtTermination(testCase.input)

		if testCase.expectedTerminatedOutput != actualTerminatedOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedTerminatedOutput, actualTerminatedOutput)
		}
	}
}
