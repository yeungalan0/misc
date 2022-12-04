package main

import (
	"fmt"
	"strings"

	_ "embed"
	"log"
)

//go:embed input
var s string

func main() {
	lines := strings.Split(s, "\n")

	solution1, err := solve1(lines)
	if err != nil {
		log.Fatalf("Error solving day: %v\n", err)
	}

	fmt.Printf("Solution 1 to day: %v\n", solution1)

	solution2, err := solve2(lines)
	if err != nil {
		log.Fatalf("Error solving day: %v\n", err)
	}

	fmt.Printf("Solution 2 to day: %v\n", solution2)
}

func solve1(lines []string) (int, error) {
	return 0, nil
}

func solve2(lines []string) (int, error) {
	return 0, nil
}
