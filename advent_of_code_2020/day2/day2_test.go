package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	input        string
	password     string
	validation   validation
	expectValid1 bool
	expectValid2 bool
}

const expectedValid1 = 2
const expectedValid2 = 1

var testCases = []testCase{
	testCase{
		input:        "1-3 a: abcde",
		password:     "abcde",
		validation:   validation{Min: 1, Max: 3, Character: "a"},
		expectValid1: true,
		expectValid2: true,
	},
	testCase{
		input:        "1-3 b: cdefg",
		password:     "cdefg",
		validation:   validation{Min: 1, Max: 3, Character: "b"},
		expectValid1: false,
		expectValid2: false,
	},
	testCase{
		input:        "2-9 c: ccccccccc",
		password:     "ccccccccc",
		validation:   validation{Min: 2, Max: 9, Character: "c"},
		expectValid1: true,
		expectValid2: false,
	},
}

func TestValidatePassword1(t *testing.T) {
	valid := 0
	for _, testCase := range testCases {
		isValid, error := validatePassword1(testCase.input)
		if error != nil {
			t.Errorf("Encountered error validating input! input: %v, error: %v", testCase.input, error)
		}
		if testCase.expectValid1 != isValid {
			t.Errorf("Expected valid to be %v, but was %v! input: %v", testCase.expectValid1, isValid, testCase.input)
		}
		if isValid {
			valid++
		}
	}

	if valid != expectedValid1 {
		t.Errorf("Failed to correctly validate passwords, expected %v valid, but got %v\n", expectedValid1, valid)
	}
}

func TestValidatePassword2(t *testing.T) {
	valid := 0
	for _, testCase := range testCases {
		isValid, error := validatePassword2(testCase.input)
		if error != nil {
			t.Errorf("Encountered error validating input! input: %v, error: %v", testCase.input, error)
		}
		if testCase.expectValid2 != isValid {
			t.Errorf("Expected valid to be %v, but was %v! input: %v", testCase.expectValid2, isValid, testCase.input)
		}
		if isValid {
			valid++
		}
	}

	if valid != expectedValid2 {
		t.Errorf("Failed to correctly validate passwords, expected %v valid, but got %v\n", expectedValid2, valid)
	}
}

func TestSplitValidationAndPassword(t *testing.T) {
	for _, testCase := range testCases {
		validation, password, error := splitValidationAndPassword(testCase.input)
		if error != nil {
			t.Errorf("error splitting validation and password - input: %v, error: %v", testCase.input, error)
		}
		if !cmp.Equal(testCase.validation, *validation) {
			t.Errorf("expected and actual validation are not equal! expected: %v, actual: %v", testCase.validation, *validation)
		}
		if testCase.password != *password {
			t.Errorf("expected and actual passwords are not equal! expected: %v, actual: %v", testCase.password, *password)
		}
	}
}
