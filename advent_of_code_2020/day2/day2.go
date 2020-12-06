package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type validation struct {
	Min       int
	Max       int
	Character string
}

func validatePassword1(input string) (bool, error) {
	validation, password, error := splitValidationAndPassword(input)
	if error != nil {
		return false, error
	}

	charCount := strings.Count(*password, validation.Character)

	if charCount >= validation.Min && charCount <= validation.Max {
		return true, nil
	}
	return false, nil
}

func validatePassword2(input string) (bool, error) {
	validation, password, error := splitValidationAndPassword(input)
	if error != nil {
		return false, error
	}

	passwordRunes := []rune(*password)
	characterAtMin := string(passwordRunes[validation.Min-1])
	characterAtMax := string(passwordRunes[validation.Max-1])
	return characterAtMin != characterAtMax &&
			(characterAtMin == validation.Character || characterAtMax == validation.Character),
		nil
}

func splitValidationAndPassword(passwordValidationString string) (*validation, *string, error) {
	passwordValidationString = strings.ReplaceAll(passwordValidationString, ": ", "-")
	passwordValidationString = strings.ReplaceAll(passwordValidationString, " ", "-")

	componentsSlice := strings.Split(passwordValidationString, "-")
	if len(componentsSlice) != 4 {
		return nil, nil, fmt.Errorf(
			"string split into more than 4 pieces! input: %v, split output: %v",
			passwordValidationString,
			componentsSlice,
		)
	}

	min, error := strconv.Atoi(componentsSlice[0])
	if error != nil {
		return nil, nil, error
	}

	max, error := strconv.Atoi(componentsSlice[1])
	if error != nil {
		return nil, nil, error
	}

	validation := validation{
		Min:       min,
		Max:       max,
		Character: componentsSlice[2],
	}

	return &validation, &(componentsSlice[3]), nil
}

func readStrings(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}
