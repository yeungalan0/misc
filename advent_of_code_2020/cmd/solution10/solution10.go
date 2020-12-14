package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day10"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day10/input", configDirectory)

	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day10.EfficientAdapterOrder(inputLines)

	fmt.Printf("Problem 1: output: %v\n", out)

	// out2 := day10.PossibleAdapterCombinations(inputLines, 3)

	// fmt.Printf("Problem 2: output: %v\n", out2)

	out2 := day10.ValidCombinations(inputLines, 3)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
