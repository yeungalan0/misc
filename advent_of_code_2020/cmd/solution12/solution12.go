package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day12"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day12/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day12.DetermineManhattanDistance(inputLines)

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day12.DetermineManhattanDistance2(inputLines)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
