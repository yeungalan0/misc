package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2022/internal/day1"
	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"

	"log"
)

func main() {
	lines := utils.ReadFileLinesToSlice("config/day1/input")

	solution1, err := day1.Solve1(lines)
	if err != nil {
		log.Fatalf("Error solving day1: %v\n", err)
	}

	fmt.Printf("Solution 1 to day1: %v\n", solution1)

	solution2, err := day1.Solve2(lines)
	if err != nil {
		log.Fatalf("Error solving day1: %v\n", err)
	}

	fmt.Printf("Solution 2 to day1: %v\n", solution2)
}
