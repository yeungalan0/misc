package main

import (
	"errors"
	"io"
	"strconv"
	"bufio"
	"os"
	"fmt"
	"log"
)

func FindPairSum(list []int, sum int) (int, int, error) {
	set := generateMap(list)
	for _, element := range list {
		compliment := sum - element
		if set[compliment] {
			return element, compliment, nil
		}
	}

	return -1, -1, errors.New("No pair found that adds to sum :(")
}

func generateMap(list []int) map[int]bool {
	set := make(map[int]bool)
	for _, integer := range list {
		set[integer] = true
	}
	return set
}

func readInts(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	var result []int

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}

	return result, scanner.Err()
}

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
}