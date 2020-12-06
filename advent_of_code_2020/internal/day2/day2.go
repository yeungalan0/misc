package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type validation struct {
	Min       int
	Max       int
	Character string
}

// ValidatePassword1 validates the password based on min/max being the min/max character count
func ValidatePassword1(input string) (bool, error) {
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

// ValidatePassword2 validates the password based on min/max being the index of the expected character
func ValidatePassword2(input string) (bool, error) {
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
