package main

import "testing"

var testList = []int{1, 2, 3, 4, 5, 6}

func TestFindPairSum(t *testing.T)  {
	expectedSum := 10
	expectedProduct := 24
	int1, int2, err := FindPairSum(testList, expectedSum)
	if err != nil {
		t.Errorf("Failed finding sum: %v", err)
	}

	if int1 + int2 != expectedSum {
		t.Errorf("Integers don't sum up to goal - int1: %v, int2 %v, expectedSum: %v", int1, int2, expectedSum)
	}

	if int1 * int2 != expectedProduct {
		t.Errorf("Integers don't multiply to expected - int1: %v, int2 %v, expectedWhenMultiplied: %v", int1, int2, expectedProduct)
	}
}

func TestFindPairSumErrorsCorrectly(t *testing.T)  {
	tooLargeSum := 1000
	int1, int2, err := FindPairSum(testList, tooLargeSum)

	if int1 != -1 || int2 != -1 || err == nil {
		t.Errorf("Expected a proper error but it was not! - int1: %v, int2 %v, expectedWhenMultiplied: %v", int1, int2, err)
	}
}