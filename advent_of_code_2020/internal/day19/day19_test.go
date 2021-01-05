package day19

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestCountMatchingRule(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"42: 9 14 | 10 1",
				"9: 14 27 | 1 26",
				"10: 23 14 | 28 1",
				"1: \"a\"",
				"11: 42 31",
				"5: 1 14 | 15 1",
				"19: 14 1 | 14 14",
				"12: 24 14 | 19 1",
				"16: 15 1 | 14 14",
				"31: 14 17 | 1 13",
				"6: 14 14 | 1 14",
				"2: 1 24 | 14 4",
				"0: 8 11",
				"13: 14 3 | 1 12",
				"15: 1 | 14",
				"17: 14 2 | 1 7",
				"23: 25 1 | 22 14",
				"28: 16 1",
				"4: 1 1",
				"20: 14 14 | 1 15",
				"3: 5 14 | 16 1",
				"27: 1 6 | 14 18",
				"14: \"b\"",
				"21: 14 1 | 1 14",
				"25: 1 1 | 1 14",
				"22: 14 14",
				"8: 42",
				"26: 14 22 | 1 20",
				"18: 15 15",
				"7: 14 5 | 1 21",
				"24: 14 1",
				"",
				"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
				"bbabbbbaabaabba",
				"babbbbaabbbbbabbbbbbaabaaabaaa",
				"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
				"bbbbbbbaaaabbbbaaabbabaaa",
				"bbbababbbbaaaaaaaabbababaaababaabab",
				"ababaaaaaabaaab",
				"ababaaaaabbbaba",
				"baabbaaaabbaaaababbaababb",
				"abbbbabbbbaaaababbbbbbaaaababb",
				"aaaaabbaabaaaaababaa",
				"aaaabbaaaabbaaa",
				"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
				"babaaabbbaaabaababbaabababaaab",
				"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
			},
			expectedOutput:  3,
			expectedOutput2: 12,
		},
	}

	for _, testCase := range testCases {
		actualOutput := CountMatchingRule(testCase.input, false)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := CountMatchingRule(testCase.input, true)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		}
	}
}

func TestBuildRuleRegex(t *testing.T) {
	testCases := []struct {
		input          map[string][]string
		expectedOutput string
	}{
		{
			input: map[string][]string{
				"0": {"4", "1", "5"},
				"1": {"2", "3", "|", "3", "2"},
				"2": {"4", "4", "|", "5", "5"},
				"3": {"4", "5", "|", "5", "4"},
				"4": {"\"a\""},
				"5": {"\"b\""},
			},
			expectedOutput: "",
		},
	}

	for _, testCase := range testCases {
		actualOutput := buildRuleRegex("0", testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %#v, but got %#v\n", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestBuildRulesMap(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput map[string][]string
	}{
		{
			input: []string{
				"0: 4 1 5",
				"1: 2 3 | 3 2",
				"2: 4 4 | 5 5",
				"3: 4 5 | 5 4",
				"4: \"a\"",
				"5: \"b\"",
			},
			expectedOutput: map[string][]string{
				"0": {"4", "1", "5"},
				"1": {"2", "3", "|", "3", "2"},
				"2": {"4", "4", "|", "5", "5"},
				"3": {"4", "5", "|", "5", "4"},
				"4": {"\"a\""},
				"5": {"\"b\""},
			},
		},
	}

	for _, testCase := range testCases {
		actualOutput := buildRulesMap(testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %#v, but got %#v\n", testCase.expectedOutput, actualOutput)
		}
	}
}
