package day14

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	inputStringList []string
	inputInt        int
	expectedOutput  int
	expectedOutput2 int
}

func TestCalculateBitmaskSum(t *testing.T) {
	testCases := []testCase{
		{
			inputStringList: []string{
				"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			expectedOutput: 165,
		},
	}

	for _, testCase := range testCases {
		actualOutput := int(CalculateBitmaskSum(testCase.inputStringList))

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestCalculateBitmaskSum2(t *testing.T) {
	testCases := []testCase{
		{
			inputStringList: []string{
				"mask = 000000000000000000000000000000X1001X",
				"mem[42] = 100",
				"mask = 00000000000000000000000000000000X0XX",
				"mem[26] = 1",
			},
			expectedOutput: 208,
		},
		{
			inputStringList: []string{
				"mask = 0000000000000000000000000000000000X1",
				"mem[1] = 4",
				"mem[2] = 3",
				"mem[3] = 2",
				"mem[4] = 1",
				"mask = 00000000000000000000000000000000X100",
				"mem[1] = 1",
				"mem[2] = 2",
				"mem[3] = 3",
			},
			expectedOutput: 16,
		},
	}

	for _, testCase := range testCases {
		actualOutput := int(CalculateBitmaskSum2(testCase.inputStringList))

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestParseInputToBitOperations(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput []bitmaskOperations
	}{
		{
			input: []string{
				"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
				"mask = XXXXXXXXXXXXX1XXXXXXXXX1XXXXXXXXXXX0",
				"mem[1] = 24",
				"mem[2] = 102",
				"mem[3] = 1004",
			},
			expectedOutput: []bitmaskOperations{
				{
					bitmask: [36]int{
						2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
						2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 0, 2,
					},
					addressesToValue: []addressToValue{
						{address: 8, value: 11},
						{address: 7, value: 101},
						{address: 8, value: 0},
					},
				},
				{
					bitmask: [36]int{
						2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2,
						2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0,
					},
					addressesToValue: []addressToValue{
						{address: 1, value: 24},
						{address: 2, value: 102},
						{address: 3, value: 1004},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		actualOutput := parseInputToBitOperations(testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("\nexpected: %v\n  actual: %v", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestSetBitValue(t *testing.T) {
	testCases := []testCase{
		{
			inputInt:        0,
			expectedOutput:  1024,
			expectedOutput2: 0,
		},
		{
			inputInt:        1200,
			expectedOutput:  1200,
			expectedOutput2: 176,
		},
	}

	for _, testCase := range testCases {
		actualOutput := int(setBitValue(int64(testCase.inputInt), true, 10))
		actualOutput2 := int(setBitValue(int64(testCase.inputInt), false, 10))

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput2, actualOutput2)
		}
	}
}

func TestMaskAddress(t *testing.T) {
	testCases := []struct {
		inputAddress   int
		inputMask      [36]int
		expectedOutput []int64
	}{
		{
			inputAddress: 42,
			inputMask: [36]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 1, 2,
			},
			expectedOutput: []int64{26, 27, 58, 59},
		},
		{
			inputAddress: 26,
			inputMask: [36]int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 2, 2,
			},
			expectedOutput: []int64{16, 17, 18, 19, 24, 25, 26, 27},
		},
	}

	for _, testCase := range testCases {
		actualOutput := maskAddress(testCase.inputAddress, testCase.inputMask)
		less := func(i, j int) bool { return actualOutput[i] < actualOutput[j] }
		sort.Slice(actualOutput, less)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}
