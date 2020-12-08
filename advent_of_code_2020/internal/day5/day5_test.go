package day5

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	input  string
	seat   Seat
	seatID int
}

var testCases = []testCase{
	{
		input:  "BFFFBBFRRR",
		seat:   Seat{Row: 70, Col: 7},
		seatID: 567,
	},
	{
		input:  "FFFBBBFRRR",
		seat:   Seat{Row: 14, Col: 7},
		seatID: 119,
	},
	{
		input:  "BBFFBBFRLL",
		seat:   Seat{Row: 102, Col: 4},
		seatID: 820,
	},
}

func TestGetSeat(t *testing.T) {
	for _, testCase := range testCases {
		seat := GetSeat(testCase.input, 7, 3)
		if !cmp.Equal(testCase.seat, seat) || testCase.seatID != GetSeatID(seat) {
			t.Errorf("Expected seat: %v, seatID: %v, but got seat: %v, seatID: %v\n",
				testCase.seat, testCase.seatID, seat, GetSeatID(seat))
		}
	}
}
