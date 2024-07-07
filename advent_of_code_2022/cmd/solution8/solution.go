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

type coord struct {
	row int
	col int
}

func solve1(lines []string) (int, error) {
	grid := getGrid(lines)
	count := 0

	// Add all visible trees from top/bottom
	for r := 0; r < len(grid); r++ {
		// Left and right visibility
		for c := 0; c < len(grid[0]); c++ {
			if isVisible(grid, r, c) {
				// fmt.Printf("VISIBLE: %#v\n", coord{r, c})
				count += 1
			}
		}

	}

	return count, nil
}

func isVisible(grid [][]int, r, c int) bool {
	if r == 0 || c == 0 {
		return true
	}

	co := coord{r, c}

	return isVisibleDir(grid, co, -1, 0) ||
		isVisibleDir(grid, co, 1, 0) ||
		isVisibleDir(grid, co, 0, -1) ||
		isVisibleDir(grid, co, 0, 1)
}

func isVisibleDir(grid [][]int, co coord, rowChange, colChange int) bool {
	height := grid[co.row][co.col]

	for {
		co.row += rowChange
		co.col += colChange

		if co.row < 0 || co.col < 0 || co.row > len(grid)-1 || co.col > len(grid[0])-1 {
			return true
		}

		currHeight := grid[co.row][co.col]

		if currHeight >= height {
			return false
		}
	}
}

func getGrid(lines []string) [][]int {
	grid := make([][]int, 0)

	for _, line := range lines {
		layer := make([]int, 0)

		for _, c := range line {
			val := utils.HandleErr(strconv.Atoi(string(c)))
			layer = append(layer, val)
		}

		grid = append(grid, layer)
	}

	return grid
}

func solve2(lines []string) (int, error) {
	grid := getGrid(lines)
	maxScore := 0

	// Add all visible trees from top/bottom
	for r := 0; r < len(grid); r++ {
		// Left and right visibility
		for c := 0; c < len(grid[0]); c++ {
			score := scenicScore(grid, r, c)
			if score > maxScore {
				// fmt.Printf("VISIBLE: %#v\n", coord{r, c})
				maxScore = score
			}
		}

	}

	return maxScore, nil

}

func scenicScore(grid [][]int, r, c int) int {
	if r == 0 || c == 0 {
		return 0
	}

	co := coord{r, c}

	return scenicScoreDir(grid, co, -1, 0) *
		scenicScoreDir(grid, co, 1, 0) *
		scenicScoreDir(grid, co, 0, -1) *
		scenicScoreDir(grid, co, 0, 1)
}

func scenicScoreDir(grid [][]int, co coord, rowChange, colChange int) int {
	height := grid[co.row][co.col]

	score := 0

	for {
		co.row += rowChange
		co.col += colChange

		if co.row < 0 || co.col < 0 || co.row > len(grid)-1 || co.col > len(grid[0])-1 {
			return score
		}

		score += 1
		currHeight := grid[co.row][co.col]

		if currHeight >= height {
			return score
		}
	}
}
