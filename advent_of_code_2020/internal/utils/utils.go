package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// ReadInts reads every line into a list of integers
func ReadInts(reader io.Reader) ([]int, error) {
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

// ReadStrings reads every line into a list of strings
func ReadStrings(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

// ReadFileLinesToSlice reads a file and returns each line in a slice
func ReadFileLinesToSlice(fileName string) []string {
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}

	list, err := ReadStrings(reader)
	if err != nil {
		log.Fatalf("Couldn't read strings from file: %v\n", err)
	}

	return list
}

// GenerateIntSet generates a map from a list of ints
func GenerateIntSet(list []int) map[int]bool {
	set := make(map[int]bool)
	for _, integer := range list {
		set[integer] = true
	}
	return set
}

// ConvertStringListToIntList returns an int list from an input stringList
func ConvertStringListToIntList(stringList []string) []int {
	intList := []int{}

	for _, stringInt := range stringList {
		intValue, error := strconv.Atoi(stringInt)
		if error != nil {
			log.Fatalf("Error converting string to int: %v", stringInt)
		}

		intList = append(intList, intValue)
	}

	return intList
}
