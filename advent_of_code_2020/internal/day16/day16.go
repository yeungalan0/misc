package day16

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

type fieldRange struct {
	field       string
	valueRanges []valueRange
}

type valueRange struct {
	min int
	max int
}

func (myFR fieldRange) Equal(otherFR fieldRange) bool {
	return cmp.Equal(myFR.valueRanges, otherFR.valueRanges) && myFR.field == otherFR.field
}

func (myVR valueRange) Equal(otherVR valueRange) bool {
	return myVR.min == otherVR.min && myVR.max == otherVR.max
}

// DetermineTicketFields returns the ticket fields determined by the possible ranges
func DetermineTicketFields(input []string) []string {
	fieldsForIndexes := [][]string{}
	possibleFieldRanges, _, nearbyTickets := parseInput(input)
	validNearbyTickets := getValidTickets(nearbyTickets, possibleFieldRanges)

	for i := 0; i < len(validNearbyTickets[0]); i++ {
		fieldsForIndex := getFieldsForIndex(i, possibleFieldRanges, validNearbyTickets)
		fieldsForIndexes = append(fieldsForIndexes, fieldsForIndex)
	}

	fields, error := getUniqueFieldsForIndex(fieldsForIndexes, []string{})

	if error != nil {
		log.Fatal(error)
	}
	return fields
}

func getUniqueFieldsForIndex(fieldsForIndexes [][]string, determinedFields []string) ([]string, error) {
	if len(fieldsForIndexes) == 0 {
		return determinedFields, nil
	}

	// TODO: could sort fields by decreasing possibilities to increasing and keep track of index
	// which should improve the backtracking efficiency
	possibleFields := fieldsForIndexes[0]
	for _, possibleField := range possibleFields {
		guessedFields := append(determinedFields, possibleField)
		newFieldsForIndex, error := updatePossibleFields(possibleField, fieldsForIndexes[1:])
		if error == nil {
			fields, error := getUniqueFieldsForIndex(newFieldsForIndex, guessedFields)
			if error == nil {
				return fields, nil
			}
		}
	}

	return []string{}, fmt.Errorf("No unique fields could be determined for path %v", determinedFields)
}

func updatePossibleFields(fieldGuess string, fieldsForIndexes [][]string) ([][]string, error) {
	newFieldsForIndex := [][]string{}

	for _, possibleFields := range fieldsForIndexes {
		possibleFieldsCopy := make([]string, len(possibleFields))
		copy(possibleFieldsCopy, possibleFields)
		possibleFieldsCopy = utils.DeleteElement(fieldGuess, possibleFieldsCopy)
		if len(possibleFieldsCopy) == 0 {
			return nil, fmt.Errorf("This is not a possible branch")
		}
		newFieldsForIndex = append(newFieldsForIndex, possibleFieldsCopy)
	}

	return newFieldsForIndex, nil
}

func getFieldsForIndex(index int, possibleFieldRanges []fieldRange, nearbyTickets [][]int) []string {
	fieldsForIndex := []string{}

	for _, fr := range possibleFieldRanges {
		isFieldForIndex := true
		for _, ticket := range nearbyTickets {
			if !isWithinRanges(ticket[index], []fieldRange{fr}) {
				isFieldForIndex = false
				break
			}
		}

		if isFieldForIndex {
			fieldsForIndex = append(fieldsForIndex, fr.field)
		}
	}

	if len(fieldsForIndex) == 0 {
		log.Fatalf("Couldn't find a ticket value for index! Index: %v\n", index)
	}

	return fieldsForIndex
}

// getValidTickets returns the valid tickets and input
func getValidTickets(nearbyTickets [][]int, fieldRanges []fieldRange) [][]int {
	validTickets := [][]int{}

	for _, ticket := range nearbyTickets {
		valid := true
		for _, ticketNumber := range ticket {
			if !isWithinRanges(ticketNumber, fieldRanges) {
				valid = false
				break
			}
		}

		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

// GetTicketErrorRate returns the error rate of ticket scanning (really sum of erroneous numbers)
func GetTicketErrorRate(input []string) int {
	errorRateSum := 0

	fieldRanges, _, nearbyTickets := parseInput(input)
	for _, ticket := range nearbyTickets {
		for _, ticketNumber := range ticket {
			if !isWithinRanges(ticketNumber, fieldRanges) {
				errorRateSum += ticketNumber
			}
		}
	}

	return errorRateSum
}

func isWithinRanges(ticketNumber int, fieldRanges []fieldRange) bool {
	for _, fr := range fieldRanges {
		for _, vr := range fr.valueRanges {
			if ticketNumber >= vr.min && ticketNumber <= vr.max {
				return true
			}
		}
	}

	return false
}

func parseInput(input []string) ([]fieldRange, []int, [][]int) {
	sections := getSections(input)
	valueRangesSlice, yourTicketSlice, nearbyTicketsSlice := sections[0], sections[1], sections[2]
	fieldRanges := convertToValueRanges(valueRangesSlice)
	yourTicket := utils.ConvertStringListToIntList(strings.Split(yourTicketSlice[1], ","))
	nearbyTickets := [][]int{}

	for _, nearbyTicketString := range nearbyTicketsSlice[1:] {
		nearbyTicketInts := utils.ConvertStringListToIntList(strings.Split(nearbyTicketString, ","))
		nearbyTickets = append(nearbyTickets, nearbyTicketInts)
	}

	return fieldRanges, yourTicket, nearbyTickets
}

func convertToValueRanges(valueRangesStrings []string) []fieldRange {
	fieldRanges := []fieldRange{}

	for _, valueRangeString := range valueRangesStrings {
		descriptionAndRanges := strings.Split(valueRangeString, ": ")
		rangesString := strings.Split(descriptionAndRanges[1], " or ")
		newFieldRange := fieldRange{field: descriptionAndRanges[0]}
		for _, minMaxString := range rangesString {
			minMaxStrings := strings.Split(minMaxString, "-")

			min, error := strconv.Atoi(minMaxStrings[0])
			if error != nil {
				log.Fatalf("Couldn't convert %v to int!", minMaxStrings[0])
			}

			max, error := strconv.Atoi(minMaxStrings[1])
			if error != nil {
				log.Fatalf("Couldn't convert %v to int!", minMaxStrings[0])
			}

			newValueRange := valueRange{min: min, max: max}
			newFieldRange.valueRanges = append(newFieldRange.valueRanges, newValueRange)
		}
		fieldRanges = append(fieldRanges, newFieldRange)
	}

	return fieldRanges
}

func getSections(input []string) [][]string {
	sections := [][]string{}
	newSection := []string{}

	for _, line := range input {
		if line == "" {
			sections = append(sections, newSection)
			newSection = []string{}
			continue
		}

		newSection = append(newSection, line)
	}

	sections = append(sections, newSection)

	return sections
}
