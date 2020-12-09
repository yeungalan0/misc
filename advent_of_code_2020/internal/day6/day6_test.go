package day6

import (
	"testing"
)

type testCase struct {
	answers               []string
	answerSum             int
	intersectionAnswerSum int
}

var testCases = []testCase{
	{
		answers: []string{
			"abc",
			"",
			"a",
			"b",
			"c",
			"",
			"ab",
			"ac",
			"",
			"a",
			"a",
			"a",
			"a",
			"",
			"b",
		},
		answerSum:             11,
		intersectionAnswerSum: 6,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testCases {
		actualAnswerSum := GetAnswerSum(testCase.answers, GetUnionAnswers)
		if testCase.answerSum != actualAnswerSum {
			t.Errorf("Expected sum of %v differed from actual sum of %v\n", testCase.answerSum, actualAnswerSum)
		}

		intersectionAnswerSum := GetAnswerSum(testCase.answers, GetIntersectionAnswers)
		if testCase.intersectionAnswerSum != intersectionAnswerSum {
			t.Errorf("Expected sum of %v differed from actual sum of %v\n", testCase.intersectionAnswerSum, intersectionAnswerSum)
		}
	}
}
