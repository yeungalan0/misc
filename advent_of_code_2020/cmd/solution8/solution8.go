package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day8"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	inputLines := utils.ReadFileLinesToSlice("config/day8/input")

	operationList := day8.ParseOperations(inputLines)

	out, _, _ := day8.GetAccumulatorAtLoop(operationList)
	fmt.Printf("Problem 1: accumulator value at loop: %v\n", out)

	out2 := day8.GetAccumulatorAtTermination(inputLines)
	fmt.Printf("Problem 2: accumulator value at termination: %v\n", out2)
}
