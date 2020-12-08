package main

import (
	"fmt"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/day4"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

func main() {
	passports := utils.ReadFileLinesToSlice("config/day4/input")
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	optionalFields := []string{"cid"}

	_, totalValid := day4.ValidatePassports(passports, requiredFields, optionalFields, false)
	fmt.Printf("Problem 1: total valid passports detected %v\n", totalValid)

	_, totalValid2 := day4.ValidatePassports(passports, requiredFields, optionalFields, true)
	fmt.Printf("Problem 2: total valid passports detected %v\n", totalValid2)
}
