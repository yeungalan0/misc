package day5

import (
	"log"
	"math"
	"strings"
)

// Seat is your row/col seat on the plane
type Seat struct {
	Row int
	Col int
}

// GetSeatID takes a seat and returns the cooresponding seatID
func GetSeatID(aSeat Seat) int {
	seatID := aSeat.Row*8 + aSeat.Col
	return seatID
}

// GetSeat takes in a binary space partitioning string key and returns the row, column, and seatID associated
func GetSeat(bspCompoundKey string, rowChars int, colChars int) Seat {
	bspCompoundKey = strings.ReplaceAll(bspCompoundKey, "F", "0")
	bspCompoundKey = strings.ReplaceAll(bspCompoundKey, "B", "1")
	bspCompoundKey = strings.ReplaceAll(bspCompoundKey, "L", "0")
	bspCompoundKey = strings.ReplaceAll(bspCompoundKey, "R", "1")

	rowString := bspCompoundKey[:rowChars]
	colString := bspCompoundKey[len(bspCompoundKey)-colChars:]

	row := binaryPicker(rowString)
	col := binaryPicker(colString)

	return Seat{Row: row, Col: col}
}

// DetermineSeat will return the missing seat following the rules from the list of binary space partitioned seats
func DetermineSeat(bspSeats []string, rowChars int, colChars int) Seat {
	possibleSeats := []Seat{}
	seatMap := map[Seat]int{}
	seatIDSet := map[int]bool{}

	for _, bspSeat := range bspSeats {
		seat := GetSeat(bspSeat, rowChars, colChars)
		seatID := GetSeatID(seat)
		seatMap[seat] = seatID
		seatIDSet[seatID] = true
	}

	rows := int(math.Pow(2, float64(rowChars)))
	cols := int(math.Pow(2, float64(colChars)))

	// Exclude the front and back rows
	for row := 1; row < rows-1; row++ {
		for col := 0; col < cols; col++ {
			seat := Seat{Row: row, Col: col}
			seatID, isPresent := seatMap[seat]
			if isPresent {
				continue
			}

			seatID = GetSeatID(seat)
			// Ensure seatID+1 and seatID-1 are present
			if seatIDSet[seatID+1] && seatIDSet[seatID-1] {
				possibleSeats = append(possibleSeats, seat)
			}
		}
	}

	if len(possibleSeats) > 1 {
		log.Fatalf("more than 1 possible seat found! %v \n", possibleSeats)
	}

	return possibleSeats[0]
}

func binaryPicker(bspKey string) int {
	seatCount := 0
	optionsCount := int(math.Pow(2, float64(len(bspKey))))

	for _, binaryChar := range bspKey {
		if binaryChar == '0' {
			optionsCount = optionsCount / 2
		} else {
			optionsCount = optionsCount / 2
			seatCount += optionsCount
		}
	}

	return seatCount
}
