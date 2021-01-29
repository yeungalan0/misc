package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day25"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day25/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	out := day25.GetEncryptionKey(inputLines, 7)

	fmt.Printf("Problem 1: output: %v\n", out)

	fmt.Printf("HOORAY! Merry Christmas/Happy holidays, and a very very happy new year to all! :)\n")
}