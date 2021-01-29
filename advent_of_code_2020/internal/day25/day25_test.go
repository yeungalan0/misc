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
		actualOutput := GetEncryptionKey(testCase.publicKeys)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}
