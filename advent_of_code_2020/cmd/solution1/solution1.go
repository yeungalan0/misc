package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day1"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"

	"log"
	"os"
)

func main() {
	goalSum := 2020
	reader, err := os.Open("config/day1/input")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}

	list, err := utils.ReadInts(reader)
	if err != nil {
		log.Fatalf("Couldn't read ints from file: %v\n", err)
	}

	set := day1.GenerateIntSet(list)

	int1, int2, err := day1.FindPairSum(set, goalSum)
	if err != nil {
		log.Fatalf("error finding a pair that adds up to %v: %v\n", goalSum, err)
	}

	fmt.Printf("Found %v and %v which add up to %v\n", int1, int2, goalSum)
	fmt.Printf("The product is: %v\n", int1*int2)

	int1, int2, int3, err := day1.FindTripleSum(list, goalSum)
	if err != nil {
		log.Fatalf("error finding a triple that adds up to %v: %v\n", goalSum, err)
	}

	fmt.Printf("Found %v, %v, and %v which add up to %v\n", int1, int2, int3, goalSum)
	fmt.Printf("The product is: %v\n", int1*int2*int3)
}
