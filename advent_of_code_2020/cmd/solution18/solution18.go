package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day18"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day18/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day18.SumEvaluatedExpressions(inputLines, false)

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day18.SumEvaluatedExpressions(inputLines, true)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
