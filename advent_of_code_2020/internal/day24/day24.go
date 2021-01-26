package day24

import "log"

// Plans of attack:

// Looked at the reddit and used this idea of a coordinates system to inspire my approach:
// https://www.redblobgames.com/grids/hexagons/#coordinates

// 0, 0 will be the starting (base) hexagon with 1, 0 nw, 1, 1 ne,
// 0, 1 to the e, -1, 0 se, -1, -1 sw, and 0, -1 w

type tileLoc struct {
	row int
	col int
}

type move struct {
	// Movement of row and/or column (can be 0, 1 or -1)
	row int
	col int
}

func (myMove move) Equal(otherMove move) bool {
	return myMove.row == otherMove.row && myMove.col == otherMove.col
}

var eMove move = move{row: 0, col: 1}
var neMove move = move{row: 1, col: 1}
var seMove move = move{row: -1, col: 0}
var wMove move = move{row: 0, col: -1}
var nwMove move = move{row: 1, col: 0}
var swMove move = move{row: -1, col: -1}

// BlackTileLayoutCount returns the number of black tiles left after flipping all desiganted tiles in the layout
func BlackTileLayoutCount(input []string, flipRounds int) int {
	moves := parseInput(input)
	outputTileLocSet := flipByMoves(moves)
	outputTileLocSet = flipByRules(outputTileLocSet, flipRounds)

	return countBlackTiles(outputTileLocSet)
}

func countBlackTiles(tileLocSet map[tileLoc]bool) int {
	count := 0
	for _, tile := range tileLocSet {
		if tile {
			count++
		}
	}

	return count
}

func flipByRules(tileLocSet map[tileLoc]bool, rounds int) map[tileLoc]bool {
	if rounds == 0 {
		return tileLocSet
	}
	
	tileLocToNeighborCount := map[tileLoc]int{}
	for aTileLoc, isBlack := range tileLocSet {
		if isBlack {
			tileLocToNeighborCount = updateNeighborCounts(aTileLoc, true, tileLocToNeighborCount)
		}
	}

	for rounds > 0 {
		flipQueue := []tileLoc{}

		for aTileLoc, blackNeighborCount := range tileLocToNeighborCount {
			if (blackNeighborCount == 0 || blackNeighborCount > 2) && tileLocSet[aTileLoc] {
				flipQueue = append(flipQueue, aTileLoc)
			} else if blackNeighborCount == 2 && !tileLocSet[aTileLoc] {
				flipQueue = append(flipQueue, aTileLoc)
			}
		}

		for _, aTileLoc := range flipQueue {
			tileLocSet[aTileLoc] = !tileLocSet[aTileLoc]
			tileLocToNeighborCount = updateNeighborCounts(aTileLoc, tileLocSet[aTileLoc], tileLocToNeighborCount)
		}

		rounds--
	}

	return tileLocSet
}

// Update neighbor counts adding 1 to each neighbor for an addition of a black neighbor
func updateNeighborCounts(aTileLoc tileLoc, isBlack bool, tileToNeighborCount map[tileLoc]int) map[tileLoc]int {
	moves := []move{neMove, eMove, seMove, nwMove, wMove, swMove}

	for _, aMove := range moves {
		if isBlack {
			tileToNeighborCount[tileLoc{row: aTileLoc.row + aMove.row, col: aTileLoc.col + aMove.col}]++
		} else {
			tileToNeighborCount[tileLoc{row: aTileLoc.row + aMove.row, col: aTileLoc.col + aMove.col}]--
		}
	}

	return tileToNeighborCount
}

func flipByMoves(moves [][]move) map[tileLoc]bool {
	// false: white, true: black
	tileLocSet := map[tileLoc]bool{}

	for _, tileMoves := range moves {
		currTile := tileLoc{row: 0, col: 0}
		for _, tileMove := range tileMoves {
			currTile.row += tileMove.row
			currTile.col += tileMove.col
		}
		tileLocSet[currTile] = !tileLocSet[currTile]
	}

	return tileLocSet
}

func parseInput(input []string) [][]move {
	moves := [][]move{}

	for _, line := range input {
		currMoves := []move{}
		var firstChar rune = 0
		for _, char := range line {
			if char == 'n' || char == 's' {
				firstChar = char
			} else if char == 'e' && firstChar == 0 {
				currMoves = append(currMoves, eMove)
			} else if char == 'e' && firstChar == 'n' {
				currMoves = append(currMoves, neMove)
				firstChar = 0
			} else if char == 'e' && firstChar == 's' {
				currMoves = append(currMoves, seMove)
				firstChar = 0
			} else if char == 'w' && firstChar == 0 {
				currMoves = append(currMoves, wMove)
			} else if char == 'w' && firstChar == 'n' {
				currMoves = append(currMoves, nwMove)
				firstChar = 0
			} else if char == 'w' && firstChar == 's' {
				currMoves = append(currMoves, swMove)
				firstChar = 0
			} else {
				log.Fatalf("in a weird state! firstChar: %v, char: %v", firstChar, char)
			}
		}
		moves = append(moves, currMoves)
	}

	return moves
}