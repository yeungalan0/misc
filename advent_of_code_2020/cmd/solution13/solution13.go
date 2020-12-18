package main

import (
	"fmt"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day13"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	configDirectory := os.Getenv("AOC_GO_CONFIG_DIR") // TODO: Backport this
	configFilePath := fmt.Sprintf("%v/day13/input", configDirectory)
	inputLines := utils.ReadFileLinesToSlice(configFilePath)

	busID, departTime := day13.FindEarliestBus(inputLines)
	out := (departTime - 1008833) * busID

	fmt.Printf("Problem 1: output: %v\n", out)

	out2 := day13.FindSynchronousEarliestTime(inputLines)

	fmt.Printf("Problem 2: output: %v\n", out2)
}
