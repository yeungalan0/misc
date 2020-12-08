package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ValidatePassports determines the number of valid passports with the expected fields in it
func ValidatePassports(passports []string, requiredFields []string, optionalFields []string, deepValidation bool) ([]bool, int) {
	totalValid := 0
	validSlice := []bool{}
	passportLines := []string{}

	for _, line := range passports {
		if line == "" {
			if isValidPassport(passportLines, requiredFields, optionalFields, deepValidation) {
				totalValid++
				validSlice = append(validSlice, true)
			} else {
				validSlice = append(validSlice, false)
			}
			passportLines = []string{}
			continue
		}

		passportLines = append(passportLines, line)
	}
	if isValidPassport(passportLines, requiredFields, optionalFields, deepValidation) {
		totalValid++
		validSlice = append(validSlice, true)
	} else {
		validSlice = append(validSlice, false)
	}
	return validSlice, totalValid
}

func isValidPassport(passportLines []string, requiredFields []string, optionalFields []string, deepValidation bool) bool {
	passport := map[string]string{}

	for _, line := range passportLines {
		keyValueStrings := strings.Split(line, " ")
		for _, keyValueString := range keyValueStrings {
			keyValueSlice := strings.Split(keyValueString, ":")
			if _, isPresent := passport[keyValueSlice[0]]; isPresent {
				fmt.Printf("found duplicate key, returning not valid!\n")
				return false
			}

			passport[keyValueSlice[0]] = keyValueSlice[1]
		}
	}

	return isValidPassportHelper(passport, requiredFields, optionalFields, deepValidation)
}

// Does the actual check on the fields of the passport
func isValidPassportHelper(passport map[string]string, requiredFields []string, optionalFields []string, deepValidation bool) bool {
	if len(passport) < len(requiredFields) || len(passport) > len(requiredFields)+len(optionalFields) {
		return false
	}

	allFields := []string{}
	allFields = append(allFields, requiredFields...)
	allFields = append(allFields, optionalFields...)

	keys := []string{}
	for key := range passport {
		keys = append(keys, key)
	}

	if isSubset(requiredFields, keys) && isSubset(keys, allFields) {
		if deepValidation {
			return deepValidate(passport)
		}
		return true
	}

	// fmt.Printf("Not all fields found! keys: %v, required: %v, all: %v\n", keys, requiredFields, allFields)
	return false
}

func isSubset(first, second []string) bool {
	// TODO: create generic function
	set := map[string]bool{}
	for _, element := range second {
		set[element] = true
	}

	for _, element := range first {
		if _, isPresent := set[element]; !isPresent {
			// fmt.Printf("Not present! element: %v, set: %v\n", element, set)
			return false
		}
	}
	return true
}

func deepValidate(passport map[string]string) bool {
	byrInt, err := strconv.Atoi(passport["byr"])
	if err != nil || byrInt < 1920 || byrInt > 2002 {
		return false
	}

	iyrInt, err := strconv.Atoi(passport["iyr"])
	if err != nil || iyrInt < 2010 || iyrInt > 2020 {
		return false
	}

	eyrInt, err := strconv.Atoi(passport["eyr"])
	if err != nil || eyrInt < 2020 || eyrInt > 2030 {
		return false
	}

	heightString := passport["hgt"]
	hgtMatch, err := regexp.MatchString("^[0-9]{2}in|[0-9]{3}cm$", heightString)
	if err != nil || !hgtMatch {
		return false
	}
	hgtVal, err := strconv.Atoi(heightString[:len(heightString)-2])
	if err != nil || !hgtMatch {
		return false
	}
	hgtUnit := heightString[len(heightString)-2:]
	if hgtUnit == "cm" && (hgtVal < 150 || hgtVal > 193) ||
		hgtUnit == "in" && (hgtVal < 59 || hgtVal > 76) {
		return false
	}

	hclMatch, err := regexp.MatchString("^#[0-9a-f]{6}$", passport["hcl"])
	if err != nil || !hclMatch {
		return false
	}

	eclMatch, err := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", passport["ecl"])
	if err != nil || !eclMatch {
		return false
	}

	pidMatch, err := regexp.MatchString("^[0-9]{9}$", passport["pid"])
	if err != nil || !pidMatch {
		return false
	}

	return true
}
