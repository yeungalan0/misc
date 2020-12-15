package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day11"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR")
	configFilePath := fmt.Sprintf("%v/day11/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day11.GetStabalizedSeatCount(inputLines, false)

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day11.GetStabalizedSeatCount(inputLines, true)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
