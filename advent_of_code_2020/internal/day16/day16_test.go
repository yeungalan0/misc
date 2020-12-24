package day16

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestDetermineTicketFields(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput []string
	}{
		{
			input: []string{
				"class: 0-1 or 4-19",
				"row: 0-5 or 8-19",
				"seat: 0-13 or 16-19",
				"",
				"your ticket:",
				"11,12,13",
				"",
				"nearby tickets:",
				"3,9,18",
				"15,1,5",
				"5,14,9",
			},
			expectedOutput: []string{
				"row", "class", "seat",
			},
		},
	}

	for _, testCase := range testCases {
		actualOutput := DetermineTicketFields(testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestGetTicketErrorRate(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"class: 1-3 or 5-7",
				"row: 6-11 or 33-44",
				"seat: 13-40 or 45-50",
				"",
				"your ticket:",
				"7,1,14",
				"",
				"nearby tickets:",
				"7,3,47",
				"40,4,50",
				"55,2,20",
				"38,6,12",
			},
			expectedOutput: 71,
		},
	}

	for _, testCase := range testCases {
		actualOutput := GetTicketErrorRate(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestParseInput(t *testing.T) {
	testCases := []struct {
		input                 []string
		expectedFieldRange    []fieldRange
		expectedTicket        []int
		expectedNearbyTickets [][]int
	}{
		{
			input: []string{
				"class: 1-3 or 5-7",
				"row: 6-11 or 33-44",
				"seat: 13-40 or 45-50",
				"",
				"your ticket:",
				"7,1,14",
				"",
				"nearby tickets:",
				"7,3,47",
				"40,4,50",
				"55,2,20",
				"38,6,12",
			},
			expectedFieldRange: []fieldRange{
				{field: "class", valueRanges: []valueRange{
					{min: 1, max: 3}, {min: 5, max: 7},
				}},
				{field: "row", valueRanges: []valueRange{
					{min: 6, max: 11}, {min: 33, max: 44},
				}},
				{field: "seat", valueRanges: []valueRange{
					{min: 13, max: 40}, {min: 45, max: 50},
				}},
			},
			expectedTicket: []int{7, 1, 14},
			expectedNearbyTickets: [][]int{
				{7, 3, 47},
				{40, 4, 50},
				{55, 2, 20},
				{38, 6, 12},
			},
		},
	}

	for _, testCase := range testCases {
		actualValueRange, actualTicket, actualNearbyTickets := parseInput(testCase.input)

		if !cmp.Equal(testCase.expectedFieldRange, actualValueRange) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedFieldRange, actualValueRange)
		}

		if !cmp.Equal(testCase.expectedTicket, actualTicket) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedTicket, actualTicket)
		}

		if !cmp.Equal(testCase.expectedNearbyTickets, actualNearbyTickets) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedNearbyTickets, actualNearbyTickets)
		}
	}
}
