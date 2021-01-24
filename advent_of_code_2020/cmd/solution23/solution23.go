package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day23"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day23/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day23.GetCupLabels(inputLines, 100, 9)

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day23.GetCupLabels(inputLines, 10000000, 1000000)

	fmt.Printf("Problem 2: output: %v\n", out2)
}