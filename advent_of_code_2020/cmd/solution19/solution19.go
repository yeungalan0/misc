package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day19"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day19/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day19.CountMatchingRule(inputLines, false)

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day19.CountMatchingRule(inputLines, true)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
