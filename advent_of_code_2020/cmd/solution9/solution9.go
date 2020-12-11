package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day9"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"

	"log"
)

func main() {
	inputLines := utils.ReadFileLinesToSlice("config/day9/input")

	out, error := day9.FindWeakness(inputLines, 25)
	if error != nil {
		log.Fatalf("Error: %v", error)
	}

	fmt.Printf("Problem 1: weakness: %v\n", out)

	out2, error := day9.ExploitWeakness(inputLines, out)
	if error != nil {
		log.Fatalf("Error: %v", error)
	}

	fmt.Printf("Problem 2: exploited weakness: %v\n", out2)
}
