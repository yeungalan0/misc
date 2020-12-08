package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day5"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	bspCompoundKeys := utils.ReadFileLinesToSlice("config/day5/input")
	highestSeatID := 0

	for _, bspCompoundKey := range bspCompoundKeys {
		seat := day5.GetSeat(bspCompoundKey, 7, 3)
		seatID := day5.GetSeatID(seat)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	mySeat := day5.DetermineSeat(bspCompoundKeys, 7, 3)

	fmt.Printf("Problem 1: Highest seatID detected %v\n", highestSeatID)
	fmt.Printf("Problem 2: My seat detected %v, seatID: %v\n", mySeat, day5.GetSeatID(mySeat))
}
