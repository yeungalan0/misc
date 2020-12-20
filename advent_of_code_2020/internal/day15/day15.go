package day15

import (
	"log"
	"strconv"
	"strings"
)

type recitedTurn struct {
	previousTurnSpoken int
	recentTurnSpoken   int
}

// GetRecitationAtTurn returns the number recited at input turn
func GetRecitationAtTurn(input []string, turn int) int {
	recitedNumbersToTurn, lastNumberRecited, currentTurn := parseInput(input)

	for currentTurn <= turn {
		lastNumberRecitedTurn, _ := recitedNumbersToTurn[lastNumberRecited]
		if lastNumberRecitedTurn.previousTurnSpoken == 0 {
			recitedNumbersToTurn[0] = recitedTurn{
				previousTurnSpoken: recitedNumbersToTurn[0].recentTurnSpoken,
				recentTurnSpoken:   currentTurn,
			}
			lastNumberRecited = 0
		} else {
			newNumberRecited := lastNumberRecitedTurn.recentTurnSpoken - lastNumberRecitedTurn.previousTurnSpoken
			recitedNumbersToTurn[newNumberRecited] = recitedTurn{
				previousTurnSpoken: recitedNumbersToTurn[newNumberRecited].recentTurnSpoken,
				recentTurnSpoken:   currentTurn,
			}
			lastNumberRecited = newNumberRecited
		}
		currentTurn++
	}

	return lastNumberRecited
}

func parseInput(input []string) (map[int]recitedTurn, int, int) {
	numbersToTurn := map[int]recitedTurn{}
	startingNumbers := strings.Split(input[0], ",")
	nextTurn := len(startingNumbers) + 1
	var lastNumber int
	var err error

	for index, numberString := range startingNumbers {
		lastNumber, err = strconv.Atoi(numberString)
		if err != nil {
			log.Fatalf("Couldn't parse number %v!", numberString)
		}

		// TODO: Handle duplicate numbers in input? Currently assuming no duplicates
		numbersToTurn[lastNumber] = recitedTurn{recentTurnSpoken: index + 1}
	}

	return numbersToTurn, lastNumber, nextTurn
}
