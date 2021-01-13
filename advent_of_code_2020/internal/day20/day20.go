package day20

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

type tileComponents struct {
	top    string
	bottom string
	right  string
	left   string
	body   []string
}

type tile struct {
	id          int
	orientation int
}

type tileEdge struct {
	side    string
	pattern string
}

func (myTC tileComponents) Equal(otherTC tileComponents) bool {
	return myTC.top == otherTC.top && myTC.bottom == otherTC.bottom &&
		myTC.right == otherTC.right && myTC.left == otherTC.left && cmp.Equal(myTC.body, otherTC.body)
}

func (myTE tileEdge) Equal(otherTE tileEdge) bool {
	return myTE.side == otherTE.side && myTE.pattern == otherTE.pattern
}

var orientations []int = []int{
	0, // input tile
	1, // 1 right quarter turn
	2, // 2 right quarter turns
	3, // 3 right quarter turns
	4, // input tile flipped
	5, // flipped 1 right quarter turn
	6, // flipped 2 right quarter turns
	7, // flipped 3 right quarter turns
}

// CalculateRoughWaters returns a number indicating how rough the waters are by finding sea monsters
// func CalculateRoughWaters(input []string) int {

// }

// FitTiles fits all given tiles together in a square and returns the corner IDs multiplied
func FitTiles(input []string) int {
	tileMap, patternMap := parseInput(input)

	sideLength := int(math.Sqrt(float64(len(tileMap))))

	solution := [][]tile{}
	for i := 0; i < sideLength; i++ {
		solution = append(solution, []tile{})
	}

	tiles2DSlice, e := fitTilesHelper(solution, tileMap, patternMap, sideLength, sideLength)
	if e != nil {
		log.Fatalf("%v\n", e)
	}

	return tiles2DSlice[0][0].id * tiles2DSlice[0][sideLength-1].id *
		tiles2DSlice[sideLength-1][0].id * tiles2DSlice[sideLength-1][sideLength-1].id
}

func parseInput(input []string) (map[int]tileComponents, map[tileEdge][]tile) {
	tileMap := map[int]tileComponents{}
	tileStrings := []string{}

	input = append(input, "")
	for _, line := range input {
		if line == "" {
			id, tileEdges := parseTileString(tileStrings)
			tileMap[id] = tileEdges
			tileStrings = []string{}
		} else {
			tileStrings = append(tileStrings, line)
		}
	}

	patternMap := generatePatterMap(tileMap)

	return tileMap, patternMap
}

func generatePatterMap(tileMap map[int]tileComponents) map[tileEdge][]tile {
	patternMap := map[tileEdge][]tile{}

	for id, inputTileEdges := range tileMap {
		for _, orientation := range orientations {
			newTile := tile{id: id, orientation: orientation}
			orientedTile := orientTileEdges(orientation, inputTileEdges)
			topTE := tileEdge{side: "top", pattern: orientedTile.top}
			rightTE := tileEdge{side: "right", pattern: orientedTile.right}
			bottomTE := tileEdge{side: "bottom", pattern: orientedTile.bottom}
			leftTE := tileEdge{side: "left", pattern: orientedTile.left}

			patternMap[topTE] = append(patternMap[topTE], newTile)
			patternMap[rightTE] = append(patternMap[rightTE], newTile)
			patternMap[bottomTE] = append(patternMap[bottomTE], newTile)
			patternMap[leftTE] = append(patternMap[leftTE], newTile)
		}
	}

	return patternMap
}

func parseTileString(tileStrings []string) (int, tileComponents) {
	top := ""
	right := ""
	bottom := ""
	left := ""
	body := []string{}

	idString := strings.Split(tileStrings[0], " ")[1]
	idString = strings.ReplaceAll(idString, ":", "")
	id, e := strconv.Atoi(idString)
	if e != nil {
		log.Fatalf("Couldn't convert %v to string!\n", idString)
	}

	tileStrings = tileStrings[1:]
	for i, tileRow := range tileStrings {
		if i == 0 {
			top = tileRow
		} else if i == len(tileStrings)-1 {
			bottom = tileRow
		} else {
			body = append(body, tileRow[1:len(tileRow)-1])
		}

		left += string(tileRow[0])
		right += string(tileRow[len(tileRow)-1])
	}

	return id, tileComponents{top: top, right: right, bottom: bottom, left: left, body: body}
}

func fitTilesHelper(currentSolution [][]tile, tileMap map[int]tileComponents, patternMap map[tileEdge][]tile, rows int, cols int) ([][]tile, error) {
	// If working row one, only need to worry about 1 edge connection
	// else need to worry about two edge connections at most

	// check location of next tile
	// get edge(s) that need matching
	// place piece
	// get edge(s) that need matching
	// find possible pieces
	// place piece and recurse

	// printSolution(currentSolution, tileMap)

	row, col := getNextTileLocation(currentSolution, rows, cols)

	if row == -1 && col == -1 {
		return currentSolution, nil
	}

	if row == 0 && col == 0 {
		for id := range tileMap {
			for orientation := range orientations {
				currentSolution[0] = []tile{{id: id, orientation: orientation}}
				solution, e := fitTilesHelper(currentSolution, tileMap, patternMap, rows, cols)
				if e == nil {
					return solution, nil
				}
			}
		}
	} else {
		edges := getNeededEdges(currentSolution, tileMap, row, col)
		for _, perfectTile := range getTilesThatFit(edges, patternMap) {
			if containsInGrid(perfectTile, currentSolution) {
				continue
			}

			guessSolution := copySolution(currentSolution)
			guessSolution[row] = append(guessSolution[row], perfectTile)

			solution, e := fitTilesHelper(guessSolution, tileMap, patternMap, rows, cols)
			if e == nil {
				return solution, nil
			}
		}
	}

	return [][]tile{}, fmt.Errorf("No solution found! :(")
}

func getTilesThatFit(edges []tileEdge, patternMap map[tileEdge][]tile) []tile {
	matchedTiles := [][]tile{}

	for _, edge := range edges {
		matchedTiles = append(matchedTiles, patternMap[edge])
	}

	if len(matchedTiles) == 0 {
		return matchedTiles[0]
	}

	intersection := []tile{}
	for _, matchedTile := range matchedTiles[0] {
		inIntersection := true
		for _, matchedSlice := range matchedTiles[1:] {
			if !containsInSlice(matchedTile, matchedSlice) {
				inIntersection = false
				break
			}
		}

		if inIntersection {
			intersection = append(intersection, matchedTile)
		}
	}

	return intersection
}

func containsInGrid(aTile tile, tileGrid [][]tile) bool {
	for _, tileRow := range tileGrid {
		if containsInSlice(aTile, tileRow) {
			return true
		}
	}

	return false
}

func containsInSlice(aTile tile, tileSlice []tile) bool {
	for _, sliceTile := range tileSlice {
		if aTile.id == sliceTile.id {
			return true
		}
	}

	return false
}

func getNeededEdges(currentSolution [][]tile, tileMap map[int]tileComponents, row, col int) []tileEdge {
	neededEdges := []tileEdge{}
	if col > 0 {
		leftTile := currentSolution[row][col-1]
		orientedTile := orientTileEdges(leftTile.orientation, tileMap[leftTile.id])
		neededEdges = append(neededEdges, tileEdge{side: "left", pattern: orientedTile.right})
	}

	if row > 0 {
		topTile := currentSolution[row-1][col]
		orientedTile := orientTileEdges(topTile.orientation, tileMap[topTile.id])
		neededEdges = append(neededEdges, tileEdge{side: "top", pattern: orientedTile.bottom})
	}

	return neededEdges
}

func orientTileEdges(orientation int, te tileComponents) tileComponents {
	edges := []string{te.top, te.right, te.bottom, te.left}
	// Flip the tile vertically
	if orientation > 3 {
		edges[1], edges[3] = edges[3], edges[1]
		edges[0] = utils.ReverseString(edges[0])
		edges[2] = utils.ReverseString(edges[2])
	}

	rotations := orientation % 4
	edges[(0+rotations)%4], edges[(1+rotations)%4], edges[(2+rotations)%4], edges[(3+rotations)%4] =
		edges[0], edges[1], edges[2], edges[3]

	if rotations == 1 {
		edges[0] = utils.ReverseString(edges[0])
		edges[2] = utils.ReverseString(edges[2])
	} else if rotations == 2 {
		edges[0] = utils.ReverseString(edges[0])
		edges[1] = utils.ReverseString(edges[1])
		edges[2] = utils.ReverseString(edges[2])
		edges[3] = utils.ReverseString(edges[3])
	} else if rotations == 3 {
		edges[1] = utils.ReverseString(edges[1])
		edges[3] = utils.ReverseString(edges[3])
	}

	return tileComponents{top: edges[0], right: edges[1], bottom: edges[2], left: edges[3]}
}

func getNextTileLocation(currentSolution [][]tile, rows, cols int) (int, int) {
	if len(currentSolution) == 0 {
		return 0, 0
	}

	for r := 0; r < rows; r++ {
		if len(currentSolution[r]) < cols {
			return r, len(currentSolution[r])
		}
	}

	return -1, -1
}

func copySolution(currentSolution [][]tile) [][]tile {
	duplicate := make([][]tile, len(currentSolution))

	for row := range currentSolution {
		duplicate[row] = make([]tile, len(currentSolution[row]))
		copy(duplicate[row], currentSolution[row])
	}

	return duplicate
}

func printSolution(currentSolution [][]tile, tileMap map[int]tileComponents) {
	if len(currentSolution) == 0 || len(currentSolution[0]) == 0 {
		fmt.Println()
		return
	}

	tileLength := len(tileMap[currentSolution[0][0].id].top)
	filler := strings.Repeat("O", tileLength-2)

	for _, tileRow := range currentSolution {
		printStrings := [][]string{}

		for i := 0; i < tileLength; i++ {
			printStrings = append(printStrings, []string{})
		}

		for _, solutionTile := range tileRow {
			orientedTile := orientTileEdges(solutionTile.orientation, tileMap[solutionTile.id])
			printStrings[0] = append(printStrings[0], orientedTile.top)

			for i := 1; i < len(orientedTile.top)-1; i++ {
				printStrings[i] = append(printStrings[i], fmt.Sprintf("%v%v%v", string(orientedTile.left[i]), filler, string(orientedTile.right[i])))
			}
			printStrings[tileLength-1] = append(printStrings[tileLength-1], orientedTile.bottom)
		}

		for _, row := range printStrings {
			for _, tileString := range row {
				fmt.Printf("%v ", tileString)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	for _, solutionTiles := range currentSolution {
		for _, solutionTile := range solutionTiles {
			fmt.Printf("%v (%v)     ", solutionTile.id, solutionTile.orientation)
		}
		fmt.Println()
	}
}
