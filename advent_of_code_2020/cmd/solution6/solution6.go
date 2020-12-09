package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day6"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	allAnswers := utils.ReadFileLinesToSlice("config/day6/input")

	unionGroupAnswersSum := day6.GetAnswerSum(allAnswers, day6.GetUnionAnswers)

	fmt.Printf("Problem 1: union group answers detected: %v\n", unionGroupAnswersSum)

	intersectionGroupAnswerSum := day6.GetAnswerSum(allAnswers, day6.GetIntersectionAnswers)

	fmt.Printf("Problem 2: intersection group answers detected: %v\n", intersectionGroupAnswerSum)
}
