package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"
	"log"

	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"
)

//go:embed input
var s string

type section struct {
	lower int
	upper int
}

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
	count := 0

	for _, pair := range lines {
		// Get sections for each pair
		splitPairs := strings.Split(pair, ",")
		if len(splitPairs) != 2 {
			fmt.Printf("SPLIT: %#v, PAIR: %v\n", splitPairs, pair)
		}
		s1 := splitPairs[0]
		s2 := splitPairs[1]

		splitS1 := strings.Split(s1, "-")
		splitS2 := strings.Split(s2, "-")

		sections := []section{
			{
				lower: utils.HandleErr(strconv.Atoi(splitS1[0])),
				upper: utils.HandleErr(strconv.Atoi(splitS1[1])),
			},
			{
				lower: utils.HandleErr(strconv.Atoi(splitS2[0])),
				upper: utils.HandleErr(strconv.Atoi(splitS2[1])),
			},
		}

		sort.Slice(sections, func(i, j int) bool {
			if sections[i].lower != sections[j].lower {
				return sections[i].lower < sections[j].lower
			}

			return sections[i].upper > sections[j].upper
		})

		if sections[0].lower <= sections[1].lower && sections[0].upper >= sections[1].upper {
			count += 1
		}
	}

	return count, nil
}

func solve2(lines []string) (int, error) {
	count := 0

	for _, pair := range lines {
		// Get sections for each pair
		splitPairs := strings.Split(pair, ",")
		if len(splitPairs) != 2 {
			fmt.Printf("SPLIT: %#v, PAIR: %v\n", splitPairs, pair)
		}
		s1 := splitPairs[0]
		s2 := splitPairs[1]

		splitS1 := strings.Split(s1, "-")
		splitS2 := strings.Split(s2, "-")

		sections := []section{
			{
				lower: utils.HandleErr(strconv.Atoi(splitS1[0])),
				upper: utils.HandleErr(strconv.Atoi(splitS1[1])),
			},
			{
				lower: utils.HandleErr(strconv.Atoi(splitS2[0])),
				upper: utils.HandleErr(strconv.Atoi(splitS2[1])),
			},
		}

		sort.Slice(sections, func(i, j int) bool {
			if sections[i].lower != sections[j].lower {
				return sections[i].lower < sections[j].lower
			}

			return sections[i].upper > sections[j].upper
		})

		if sections[1].lower <= sections[0].upper {
			count += 1
		}
	}

	return count, nil
}
