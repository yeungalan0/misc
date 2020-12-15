package day11

import "testing"

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestGetStabalizedSeatCount(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"L.LL.LL.LL",
				"LLLLLLL.LL",
				"L.L.L..L..",
				"LLLL.LL.LL",
				"L.LL.LL.LL",
				"L.LLLLL.LL",
				"..L.L.....",
				"LLLLLLLLLL",
				"L.LLLLLL.L",
				"L.LLLLL.LL",
			},
			expectedOutput:  37,
			expectedOutput2: 26,
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetStabalizedSeatCount(testCase.input, false)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := GetStabalizedSeatCount(testCase.input, true)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}
