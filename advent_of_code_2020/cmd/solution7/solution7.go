package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day7"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	inputLines := utils.ReadFileLinesToSlice("config/day7/input")

	out := day7.CountPossibleContainerBags(inputLines, "shiny gold", 1)

	fmt.Printf("Problem 1: container bags detected: %v\n", out)

	out2 := day7.CountAllContainedBags(inputLines, "shiny gold")

	fmt.Printf("Problem 2: contained bags detected: %v\n", out2)
}
