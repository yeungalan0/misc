package main

import (
	"fmt"

	day "github.com/yeungalan0/misc/advent_of_code_2022/internal/day3"
	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"

	"log"
)

func main() {
	lines := utils.ReadFileLinesToSlice("config/day3/input")

	solution1, err := day.Solve1(lines)
	if err != nil {
		log.Fatalf("Error solving day2: %v\n", err)
	}

	fmt.Printf("Solution 1 to day2: %v\n", solution1)

	solution2, err := day.Solve2(lines)
	if err != nil {
		log.Fatalf("Error solving day2: %v\n", err)
	}

	fmt.Printf("Solution 2 to day2: %v\n", solution2)
}
