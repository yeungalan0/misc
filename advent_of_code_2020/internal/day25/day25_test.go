package day25

import (
	"testing"
)

func TestGetEncryptionKey(t *testing.T) {
	testCases := []struct {
		publicKeys     []string
		subjectNumber  int
		expectedOutput int
	}{
		{
			publicKeys:      []string{"5764801", "17807724"},
			subjectNumber: 7,
			expectedOutput:  14897079,
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetEncryptionKey(testCase.publicKeys, 7)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestRunTransformation(t *testing.T) {
	testCases := []struct {
		loopSize       int
		subjectNumber  int
		expectedOutput int
	}{
		{
			loopSize: 1,
			subjectNumber: 7,
			expectedOutput: 7,
		},
		{
			loopSize: 2,
			subjectNumber: 7,
			expectedOutput: 49,
		},
		{
			loopSize: 3,
			subjectNumber: 7,
			expectedOutput: 343,
		}, 
		{
			loopSize: 11,
			subjectNumber: 7,
			expectedOutput: 17807724,
		},
	}

	for _, testCase := range testCases {
		actualOutput, _ := runTransformation(testCase.loopSize, testCase.subjectNumber, map[int]int{})

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}