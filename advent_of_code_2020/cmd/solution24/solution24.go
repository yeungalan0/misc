package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day24"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day24/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day24.BlackTileLayoutCount(inputLines, 0)

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day24.BlackTileLayoutCount(inputLines, 100)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
