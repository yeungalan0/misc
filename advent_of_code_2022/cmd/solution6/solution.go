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
	queue := make([]string, 0)
	index := 0

	if len(lines) > 1 {
		return 0, fmt.Errorf("unexepected number of lines (%d): %v", len(lines), lines)
	}

	for i, c := range lines[0] {
		if len(queue) < 4 {
			queue = append(queue, string(c))
			continue
		}

		if unique(queue) {
			index = i
			break
		}

		queue = append(queue, string(c))
		queue = queue[1:]
	}

	return index, nil
}

func unique(queue []string) bool {
	seen := make(map[string]bool, 0)

	for _, v := range queue {
		if _, ok := seen[v]; ok {
			return false
		}

		seen[v] = true
	}

	return true
}

func solve2(lines []string) (int, error) {
	queue := make([]string, 0)
	index := 0

	if len(lines) > 1 {
		return 0, fmt.Errorf("unexepected number of lines (%d): %v", len(lines), lines)
	}

	for i, c := range lines[0] {
		if len(queue) < 14 {
			queue = append(queue, string(c))
			continue
		}

		if unique(queue) {
			index = i
			break
		}

		queue = append(queue, string(c))
		queue = queue[1:]
	}

	return index, nil
}
