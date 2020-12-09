package day7

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	rules        []string
	expectedBags int
}

type generateBagToNumberCase struct {
	input          []string
	expectedOutput map[string]int
}

type createRulesMapCase struct {
	input          []string
	expectedOutput map[string]map[string]int
}

func TestCountPossibleContainerBags(t *testing.T) {
	testCases := []testCase{
		{
			rules: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			expectedBags: 4,
		},
	}

	for _, testCase := range testCases {
		actualBags := CountPossibleContainerBags(testCase.rules, "shiny gold", 1)

		if testCase.expectedBags != actualBags {
			t.Errorf("expected %v possible container bags, but got %v", testCase.expectedBags, actualBags)
		}
	}
}

func TestCountAllContainedBags(t *testing.T) {
	testCases := []testCase{
		{
			rules: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			expectedBags: 32,
		},
		{
			rules: []string{
				"shiny gold bags contain 2 dark red bags.",
				"dark red bags contain 2 dark orange bags.",
				"dark orange bags contain 2 dark yellow bags.",
				"dark yellow bags contain 2 dark green bags.",
				"dark green bags contain 2 dark blue bags.",
				"dark blue bags contain 2 dark violet bags.",
				"dark violet bags contain no other bags.",
			},
			expectedBags: 126,
		},
	}

	for _, testCase := range testCases {
		actualBags := CountAllContainedBags(testCase.rules, "shiny gold")

		if testCase.expectedBags != actualBags {
			t.Errorf("expected %v bags, but got %v", testCase.expectedBags, actualBags)
		}
	}
}

func TestGenerateBagToNumber(t *testing.T) {
	var testCases = []generateBagToNumberCase{
		{
			input:          []string{"1 bright white"},
			expectedOutput: map[string]int{"bright white": 1},
		},
		{
			input:          []string{"2 shiny gold", "9 faded blue"},
			expectedOutput: map[string]int{"shiny gold": 2, "faded blue": 9},
		},
		{
			input:          []string{"3 faded blue", "4 dotted black", "1 funky orange"},
			expectedOutput: map[string]int{"faded blue": 3, "dotted black": 4, "funky orange": 1},
		},
	}

	for _, testCase := range testCases {
		actualOutputMap := generateBagToNumber(testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutputMap) {
			t.Errorf("expected %#v, but got %#v", testCase.expectedOutput, actualOutputMap)
		}
	}
}

func TestCreateRulesMap(t *testing.T) {
	var testCases = []createRulesMapCase{
		{
			input: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			},
			expectedOutput: map[string]map[string]int{
				"light red":    {"bright white": 1, "muted yellow": 2},
				"dark orange":  {"bright white": 3, "muted yellow": 4},
				"bright white": {"shiny gold": 1},
				"muted yellow": {"shiny gold": 2, "faded blue": 9},
			},
		},
	}

	for _, testCase := range testCases {
		actualOutputMap := createRulesMap(testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutputMap) {
			t.Errorf("expected %#v, but got %#v", testCase.expectedOutput, actualOutputMap)
		}
	}
}
