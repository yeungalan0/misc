package day11

import (
	"fmt"
	"math"
)

type location struct {
	row int
	col int
}

type seat struct {
	seatType              rune
	adjacentOccupiedCount int
	adjacentCells         []location
}

type adjacentOccupied func(seatGrid [][]seat, row, col int) (int, []location)

// GetStabalizedSeatCount returns the number of occupied seats after modeling the seating
// system stabilizes
func GetStabalizedSeatCount(seatList []string, lineOfSightAdjacent bool) int {
	seatGrid := convertToSeatMatrix(seatList)
	if !lineOfSightAdjacent {
		seatGrid = populateAdjacent(seatGrid, getAdjacentOccupied)
	} else {
		seatGrid = populateAdjacent(seatGrid, getAdjacentOccupied2)
	}

	stableSeatGrid := modelSystemToStability(seatGrid, lineOfSightAdjacent)
	return countOccupied(stableSeatGrid)
}

func modelSystemToStability(seatGrid [][]seat, lineOfSightAdjacent bool) [][]seat {
	newSeatGrid := copySeatMatrix(seatGrid)
	newSeatGrid, changed := modelSystemIteration(newSeatGrid, lineOfSightAdjacent)

	for changed {
		newSeatGrid, changed = modelSystemIteration(newSeatGrid, lineOfSightAdjacent)
	}

	return newSeatGrid
}

func modelSystemIteration(seatGrid [][]seat, lineOfSightAdjacent bool) ([][]seat, bool) {
	var minOccupiedTransformCount int
	if lineOfSightAdjacent {
		minOccupiedTransformCount = 5
	} else {
		minOccupiedTransformCount = 4
	}
	updateQueue := []location{}
	for row, seatRow := range seatGrid {
		for col, seat := range seatRow {
			newSeatValue := seat.seatType
			if seat.seatType == 'L' && seat.adjacentOccupiedCount == 0 {
				newSeatValue = '#'
				updateQueue = append(updateQueue, location{row: row, col: col})
			} else if seat.seatType == '#' && seat.adjacentOccupiedCount >= minOccupiedTransformCount {
				newSeatValue = 'L'
				updateQueue = append(updateQueue, location{row: row, col: col})
			}

			seatGrid[row][col].seatType = newSeatValue
		}
	}

	for _, changedSeatLocation := range updateQueue {
		seatGrid = updateAdjacentCells(seatGrid, changedSeatLocation)
	}

	return seatGrid, len(updateQueue) > 0
}

func populateAdjacent(seatGrid [][]seat, adjacentOccupiedFn adjacentOccupied) [][]seat {
	for row, seatRow := range seatGrid {
		for col := range seatRow {
			seatGrid[row][col].adjacentOccupiedCount, seatGrid[row][col].adjacentCells = adjacentOccupiedFn(seatGrid, row, col)
		}
	}

	return seatGrid
}

// getAdjacentOccupied returns the number of occupied seats and locations of seats directly adjacent to the input seat
func getAdjacentOccupied(seatGrid [][]seat, row, col int) (int, []location) {
	occupiedCount := 0
	adjacentCells := []location{}

	startRow := int(math.Max(0, float64(row-1)))
	startCol := int(math.Max(0, float64(col-1)))
	// +2 since these are exclusive
	endRow := int(math.Min(float64(len(seatGrid)), float64(row+2)))
	endCol := int(math.Min(float64(len(seatGrid[0])), float64(col+2)))

	for r := startRow; r < endRow; r++ {
		for c := startCol; c < endCol; c++ {
			if r == row && c == col || seatGrid[r][c].seatType == '.' {
				continue
			} else if seatGrid[r][c].seatType == '#' {
				occupiedCount++
			}

			adjacentCells = append(adjacentCells, location{row: r, col: c})
		}
	}

	return occupiedCount, adjacentCells
}

// getAdjacentOccupied2 count all the adjacent occupied seats and returns all the adjacent seats form a seat's line of sight
func getAdjacentOccupied2(seatGrid [][]seat, row, col int) (int, []location) {
	occupiedCount := 0
	adjacentCells := []location{}
	currentLocation := location{row: row, col: col}

	for r := -1; r < 2; r++ {
		for c := -1; c < 2; c++ {
			if r == 0 && c == 0 {
				continue
			}

			adjacentLocation, error := crawlDirection(seatGrid, currentLocation, r, c)
			if error == nil {
				adjacentCells = append(adjacentCells, adjacentLocation)
				if seatGrid[adjacentLocation.row][adjacentLocation.col].seatType == '#' {
					occupiedCount++
				}
			}
		}
	}

	return occupiedCount, adjacentCells
}

// crawlDirection recursively moves in direction and returns the location of the first seat found or errors
func crawlDirection(seatGrid [][]seat, lastCheckLocation location, rowMovement, colMovement int) (location, error) {
	newRow := lastCheckLocation.row + rowMovement
	newCol := lastCheckLocation.col + colMovement

	if newRow < 0 || newRow >= len(seatGrid) || newCol < 0 || newCol >= len(seatGrid[0]) {
		return location{row: -1, col: -1}, fmt.Errorf("Out of range: newRow: %v, newCol: %v", newRow, newCol)
	}

	checkLocation := location{row: newRow, col: newCol}
	if seatGrid[newRow][newCol].seatType == '#' ||
		seatGrid[newRow][newCol].seatType == 'L' {
		return checkLocation, nil
	}

	return crawlDirection(seatGrid, checkLocation, rowMovement, colMovement)
}

func copySeatMatrix(seatMatrix [][]seat) [][]seat {
	newMatrix := [][]seat{}
	for rowIndex, matrixRow := range seatMatrix {
		newMatrix = append(newMatrix, []seat{})
		for _, element := range matrixRow {
			newMatrix[rowIndex] = append(newMatrix[rowIndex], element)
		}
	}

	return newMatrix
}

func updateAdjacentCells(seatGrid [][]seat, seatLocation location) [][]seat {
	currentSeat := seatGrid[seatLocation.row][seatLocation.col]
	currentSeatIsOccupied := currentSeat.seatType == '#'

	for _, cell := range currentSeat.adjacentCells {
		if currentSeatIsOccupied {
			seatGrid[cell.row][cell.col].adjacentOccupiedCount++
		} else {
			seatGrid[cell.row][cell.col].adjacentOccupiedCount--
		}
	}

	return seatGrid
}

func countOccupied(seatGrid [][]seat) int {
	occupiedCount := 0

	for _, seatRow := range seatGrid {
		for _, seat := range seatRow {
			if seat.seatType == '#' {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}

func printSeatGrid(seatGrid [][]seat) {
	for _, seatRow := range seatGrid {
		for _, seat := range seatRow {
			fmt.Printf("%v", string(seat.seatType))
		}
		fmt.Println("")
	}
	fmt.Printf("\n\n")
}

func convertToSeatMatrix(input []string) [][]seat {
	output := [][]seat{}

	for lineIndex, line := range input {
		output = append(output, []seat{})
		for _, char := range line {
			output[lineIndex] = append(output[lineIndex], seat{seatType: char, adjacentOccupiedCount: 0})
		}
	}

	return output
}
