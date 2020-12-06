package main

import (
	"fmt"
	"internal/day2"
	"internal/utils"
	"log"
	"os"
)

func main() {
	valid1 := 0
	valid2 := 0

	reader, err := os.Open("input")

	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}

	list, err := utils.ReadStrings(reader)

	for _, inputLine := range list {
		isValid1, error := day2.ValidatePassword1(inputLine)
		if error != nil {
			log.Fatalf("Couldn't validate input: %v, error: %v\n", inputLine, err)
		}
		if isValid1 {
			valid1++
		}

		isValid2, error := day2.ValidatePassword2(inputLine)
		if error != nil {
			log.Fatalf("Couldn't validate2 input: %v, error: %v\n", inputLine, err)
		}
		if isValid2 {
			valid2++
		}
	}

	fmt.Printf("Problem 1: detected %v valid passwords based on criteria\n", valid1)
	fmt.Printf("Problem 2: detected %v valid passwords based on criteria\n", valid2)

}
