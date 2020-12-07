package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day2"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	valid1 := 0
	valid2 := 0

	reader, err := os.Open("config/day2/input")

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
