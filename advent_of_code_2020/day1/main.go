package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	goalSum := 2020
	reader, err := os.Open("input")
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}

	list, err := readInts(reader)
	if err != nil {
		log.Fatalf("Couldn't read ints from file: %v\n", err)
	}

	int1, int2, err := FindPairSum(list, goalSum)
	if err != nil {
		log.Fatalf("error finding a pair that adds up to %v: %v\n", goalSum, err)
	}

	fmt.Printf("Found %v and %v which add up to %v\n", int1, int2, goalSum)
	fmt.Printf("The product is: %v\n", int1 * int2)

	int1, int2, int3, err := FindTripleSum(list, goalSum)
	if err != nil {
		log.Fatalf("error finding a triple that adds up to %v: %v\n", goalSum, err)
	}

	fmt.Printf("Found %v, %v, and %v which add up to %v\n", int1, int2, int3, goalSum)
	fmt.Printf("The product is: %v\n", int1 * int2 * int3)
}