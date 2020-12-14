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

	out := day11.GetStabalizedSeatCount(inputLines)

	fmt.Printf("Problem 1: output: %v\n", out)
}
