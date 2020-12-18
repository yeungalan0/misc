package day13

import "testing"

type testCase struct {
	input           []string
	inputInt1       int
	inputInt2       int
	expectedOutput  int
	expectedOutput2 int
	expectedOutput3 int
}

func TestFindEarliestBus(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			expectedOutput:  295,
			expectedOutput2: 1068781,
		},
	}

	for _, testCase := range testCases {
		busID, departureTime := FindEarliestBus(testCase.input)

		actualOutput := (departureTime - 939) * busID
		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestFindSynchronousEarliestTime(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			expectedOutput2: 1068781,
		},
		{
			input: []string{
				"-1",
				"17,x,13,19",
			},
			expectedOutput2: 3417,
		},
		{
			input: []string{
				"-1",
				"67,7,59,61",
			},
			expectedOutput2: 754018,
		},
		{
			input: []string{
				"-1",
				"67,x,7,59,61",
			},
			expectedOutput2: 779210,
		},
		{
			input: []string{
				"-1",
				"67,7,x,59,61",
			},
			expectedOutput2: 1261476,
		},
		{
			input: []string{
				"-1",
				"1789,37,47,1889",
			},
			expectedOutput2: 1202161486,
		},
	}

	for _, testCase := range testCases {
		actualOutput2 := int(FindSynchronousEarliestTime(testCase.input))

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}

func TestExtendedEuclid(t *testing.T) {
	testCases := []testCase{
		{
			inputInt1:       102,
			inputInt2:       38,
			expectedOutput:  2,  // GCD
			expectedOutput2: 3,  // a - first coefficient
			expectedOutput3: -8, // b - second coefficient
		},
		{
			inputInt1:       17,
			inputInt2:       43,
			expectedOutput:  1,  // GCD
			expectedOutput2: -5, // a - first coefficient
			expectedOutput3: 2,  // b - second coefficient
		},
	}

	for _, testCase := range testCases {
		actualGCD, actualCoefa, actualCoefb := extendedEuclid(int64(testCase.inputInt1), int64(testCase.inputInt2))

		if testCase.expectedOutput != int(actualGCD) ||
			testCase.expectedOutput2 != int(actualCoefa) ||
			testCase.expectedOutput3 != int(actualCoefb) {
			t.Errorf("(expected: actual) - %v:%v, %v:%v, %v:%v", testCase.expectedOutput, actualGCD, testCase.expectedOutput2, actualCoefa, testCase.expectedOutput3, actualCoefb)
		}
	}
}

func TestFindInverseMod(t *testing.T) {
	testCases := []testCase{
		{
			inputInt1:      17,
			inputInt2:      43,
			expectedOutput: 38,
		},
	}

	for _, testCase := range testCases {
		actualOutput := findInverseMod(int64(testCase.inputInt1), int64(testCase.inputInt2))

		if testCase.expectedOutput != int(actualOutput) {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}
