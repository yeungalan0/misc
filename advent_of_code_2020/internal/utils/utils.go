package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"reflect"
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

// ReverseSlice reverses an arbitrary slice
// Note: You'll have to use a . cast to convert the interface to the right type,
// e.g. reverseSlice(operations).([]bitmaskOperations)
func ReverseSlice(slice interface{}) interface{} {
	size := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}

	return slice
}

// DeleteElements returns a new slice with the input elements deleted (not accounting for duplicates)
func DeleteElements(elements []string, slice []string) []string {
	for _, element := range elements {
		slice = DeleteElement(element, slice)
	}

	return slice
}

// DeleteElement returns a new slice with the element deleted
func DeleteElement(element string, slice []string) []string {
	for index, sliceElement := range slice {
		if element == sliceElement {
			return append(slice[:index], slice[index+1:]...)
		}
	}

	// log.Printf("Element (%v) not found in slice (%v)", element, slice)
	return slice
}

// ReverseString returns the input string reversed
func ReverseString(theString string) string {
	runes := []rune(theString)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// SliceIndex returns the index of an element found when predicate is true
// Source: https://stackoverflow.com/questions/8307478/how-to-find-out-element-position-in-slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// Contains returns a boolean based on whether the element was found by the predicate or not
func Contains(limit int, predicate func(i int) bool) bool {
	if SliceIndex(limit, predicate) == -1 {
		return false
	}

	return true
}