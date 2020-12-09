package day6

type getAnswers func([]string) int

// GetAnswerSum takes all answers and returns the sum of unique answers for each group
func GetAnswerSum(allAnswers []string, fn getAnswers) int {
	uniqueGroupAnswersSum := 0
	groupAnswers := []string{}

	for _, answer := range allAnswers {
		if answer == "" {
			uniqueGroupAnswersSum += fn(groupAnswers)
			groupAnswers = []string{}
			continue
		}

		groupAnswers = append(groupAnswers, answer)
	}

	uniqueGroupAnswersSum += fn(groupAnswers)

	return uniqueGroupAnswersSum
}

// GetUnionAnswers returns the number of answers in a set union of answers
func GetUnionAnswers(groupAnswer []string) int {
	answersSet := map[rune]bool{}

	for _, personalAnswer := range groupAnswer {
		for _, answer := range personalAnswer {
			answersSet[answer] = true
		}
	}

	return len(answersSet)
}

// GetIntersectionAnswers returns the number of answers in a set intersection of answers
func GetIntersectionAnswers(groupAnswer []string) int {
	answerSet := generateAnswerSet(groupAnswer[0])

	newGroupAnswer := groupAnswer[1:]

	for _, personalAnswer := range newGroupAnswer {
		personalAnswerSet := generateAnswerSet(personalAnswer)
		for answer := range answerSet {
			if !personalAnswerSet[answer] {
				delete(answerSet, answer)
			}
		}
	}

	return len(answerSet)
}

func generateAnswerSet(personalAnswers string) map[rune]bool {
	set := map[rune]bool{}
	for _, answer := range personalAnswers {
		set[answer] = true
	}

	return set
}
