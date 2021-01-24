package day23

import (
	"log"
	"testing"
)

func TestGetCupLabels(t *testing.T) {
	testCases := []struct {
		input           []string
		rounds          int
		cupNumber       int
		expectedOutput  string
		expectedOutput2 string
	}{
		{
			input: []string{
				"389125467",
			},
			rounds:         10,
			cupNumber:      9,
			expectedOutput: "92658374",
		},
		{
			input: []string{
				"389125467",
			},
			rounds:         100,
			cupNumber:      9,
			expectedOutput: "67384529",
		},
		{
			input: []string{
				"389125467",
			},
			rounds: 10000000,
			cupNumber: 1000000,
			expectedOutput:  "149245887792",
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetCupLabels(testCase.input, testCase.rounds, testCase.cupNumber)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		// actualOutput2 := GetWinningScore(testCase.input, true)

		// if testCase.expectedOutput2 != actualOutput2 {
		// 	t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		// }
	}
}

func TestSimulateCrabCups(t *testing.T) {
	testCases := []struct {
		cups        []int
		rounds int
		expectedOrder []int
	}{
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOrder: []int{3, 2, 8, 9, 1, 5, 4, 6, 7},
			rounds: 1,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOrder: []int{2, 5, 4, 6, 7, 8, 9, 1, 3},
			rounds: 2,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOrder: []int{5, 8, 9, 1, 3, 4, 6, 7, 2},
			rounds: 3,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOrder: []int{8, 4, 6, 7, 9, 1, 3, 2, 5},
			rounds: 4,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOrder: []int{4, 1, 3, 6, 7, 9, 2, 5, 8},
			rounds: 5,
		},
		{
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOrder: []int{1, 9, 3, 6, 7, 2, 5, 8, 4},
			rounds: 6,
		},
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	for _, testCase := range testCases {
		cupCircle := getCircularLinkedList(testCase.cups, len(testCase.cups))
		outputCupCircle := simulateCrabCups(testCase.rounds, testCase.cups[0], cupCircle)

		currNode, _ := getNodeByVal(testCase.expectedOrder[0], outputCupCircle)

		for _, cupVal := range testCase.expectedOrder {
			if cupVal != currNode.val {
				t.Errorf("expected %v, but got %v\n", cupVal, currNode.val)
			}
			currNode = currNode.next
		}
	}
}

func TestGetDestinationNode(t *testing.T) {
	testCases := []struct {
		currCupVal      int
		cups            []int
		expectedOutputVal  int
	}{
		{
			currCupVal:      5,
			cups:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedOutputVal:  4,
		},
		{
			currCupVal:    3,
			cups:            []int{5, 8, 3, 7, 4, 1, 9, 2, 6},
			expectedOutputVal:  2,
		},
		{
			currCupVal:    9,
			cups:            []int{8, 2, 5, 7, 4, 1, 9, 3, 6},
			expectedOutputVal:  7,
		},
		{
			currCupVal:    8,
			cups:            []int{5, 8, 9, 1, 3, 4, 6, 7, 2},
			expectedOutputVal:  7,
		},
		{
			currCupVal:    2,
			cups:            []int{5, 2, 9, 1, 3, 4, 6, 7, 8},
			expectedOutputVal:  8,
		},
		{
			currCupVal: 3,
			cups: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			expectedOutputVal: 2,
		},
	}

	for _, testCase := range testCases {
		cupCircle := getCircularLinkedList(testCase.cups, len(testCase.cups))
		currNode, _ := getNodeByVal(testCase.currCupVal, cupCircle)
		closeNodeList := getCloseNodeList(currNode)
		actualOutput := getDestinationNode(currNode, closeNodeList, cupCircle)

		if testCase.expectedOutputVal != actualOutput.val {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutputVal, actualOutput.val)
		}
	}
}
