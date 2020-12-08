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
