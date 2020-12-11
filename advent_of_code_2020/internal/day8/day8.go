package day8

import (
	"log"
	"strconv"
	"strings"
)

// Operation struct to determine the operation, value, if it's been visited, or if it's been modified
type Operation struct {
	name     string
	value    int
	visited  bool
	modified bool
}

// GetAccumulatorAtTermination edits the operation list by replacing a jump with noop or vice versa, efficiently finding one
// that will terminate the operation list successfully
func GetAccumulatorAtTermination(input []string) int {
	operationList := ParseOperations(input)

	accumulator, didLoop, modificationList := GetAccumulatorAtLoop(operationList)

	for didLoop {
		operationList := ParseOperations(input)
		for index, operation := range modificationList {
			if operation.name == "jmp" && operation.modified == false && modificationList[index+1].visited == false {
				operationList[index].name = "nop"
				modificationList[index].modified = true
				break
			} else if operation.name == "nop" && operation.modified == false && modificationList[index+operation.value].visited == false {
				operationList[index].name = "jmp"
				modificationList[index].modified = true
				break
			}
		}

		accumulator, didLoop, _ = GetAccumulatorAtLoop(operationList)
	}

	return accumulator
}

// GetAccumulatorAtLoop returns the value of the accumulator right before doing a previous operation, and boolean representing a loop (true)
// or the accumulator after executing the last line, and a false for didn't loop
func GetAccumulatorAtLoop(operationList []Operation) (int, bool, []Operation) {
	accumulator := 0
	index := 0

	for true {
		if index >= len(operationList) {
			return accumulator, false, nil
		}

		// Without the & you get the element by value and infinite loop yourself...
		currentOperation := &operationList[index]
		if currentOperation.visited {
			break
		}

		if currentOperation.name == "acc" {
			accumulator += currentOperation.value
			index++
		} else if currentOperation.name == "jmp" {
			index += currentOperation.value
		} else {
			index++
		}

		currentOperation.visited = true
	}

	return accumulator, true, operationList
}

// ParseOperations parses an input string list and turns it into a list of Operation structs
func ParseOperations(operations []string) []Operation {
	operationList := []Operation{}

	for _, op := range operations {
		operationSlice := strings.Split(op, " ")
		value, error := strconv.Atoi(operationSlice[1])
		if error != nil {
			log.Fatalf("Couldn't process as number: %v, operation: %v", operationSlice[1], op)
		}

		operationList = append(operationList, Operation{name: operationSlice[0], value: value, visited: false})
	}

	return operationList
}
