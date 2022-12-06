package main

import (
	"fmt"
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

func solve1(lines []string) (string, error) {
	stacks := [][]string{
		{"Z", "T", "F", "R", "W", "J", "G"},
		{"G", "W", "M"},
		{"J", "N", "H", "G"},
		{"J", "R", "C", "N", "W"},
		{"W", "F", "S", "B", "G", "Q", "V", "M"},
		{"S", "R", "T", "D", "V", "W", "C"},
		{"H", "B", "N", "C", "D", "Z", "G", "V"},
		{"S", "J", "N", "M", "G", "C"},
		{"G", "P", "N", "W", "C", "J", "D", "L"},
	}
	// stacks := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	for _, instruction := range lines {
		number, from, to := parseInstruction(instruction)

		// Account for 0 based indexing
		from = from - 1
		to = to - 1

		for i := 0; i < number; i++ {
			newFrom, val := pop(stacks[from])
			stacks[to] = append(stacks[to], val)
			stacks[from] = newFrom
		}
	}

	solution := ""

	for _, stack := range stacks {
		solution = solution + stack[len(stack)-1]
	}

	return solution, nil

}

func parseInstruction(instruction string) (int, int, int) {
	s := strings.Split(instruction, " ")

	return utils.HandleErr(strconv.Atoi(s[1])), utils.HandleErr(strconv.Atoi(s[3])), utils.HandleErr(strconv.Atoi(s[5]))
}

func pop(stack []string) ([]string, string) {
	val := stack[len(stack)-1]

	return stack[:len(stack)-1], val
}

func solve2(lines []string) (string, error) {
	stacks := [][]string{
		{"Z", "T", "F", "R", "W", "J", "G"},
		{"G", "W", "M"},
		{"J", "N", "H", "G"},
		{"J", "R", "C", "N", "W"},
		{"W", "F", "S", "B", "G", "Q", "V", "M"},
		{"S", "R", "T", "D", "V", "W", "C"},
		{"H", "B", "N", "C", "D", "Z", "G", "V"},
		{"S", "J", "N", "M", "G", "C"},
		{"G", "P", "N", "W", "C", "J", "D", "L"},
	}
	// stacks := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	for _, instruction := range lines {
		number, from, to := parseInstruction(instruction)

		// Account for 0 based indexing
		from = from - 1
		to = to - 1

		crates := stacks[from][len(stacks[from])-number:]
		newFrom := stacks[from][:len(stacks[from])-number]
		stacks[to] = append(stacks[to], crates...)
		stacks[from] = newFrom
	}

	solution := ""

	for _, stack := range stacks {
		solution = solution + stack[len(stack)-1]
	}

	return solution, nil
}
