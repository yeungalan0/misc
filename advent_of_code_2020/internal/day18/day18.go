package day18

import (
	"log"
	"strconv"
	"strings"
)

// SumEvaluatedExpressions returns the sum of the evaluated expressions
func SumEvaluatedExpressions(input []string, advanced bool) int {
	sum := 0
	expressions := parseInput(input)
	for _, expression := range expressions {
		sum += evaluateExpression(expression, advanced)
	}

	return sum
}

func evaluateExpression(expression []string, advanced bool) int {
	var evaluatedValue int
	if advanced {
		evaluatedValue, _ = evaluateExpressionHelperAdvanced(expression)
	} else {
		evaluatedValue, _ = evaluateExpressionHelper(expression)
	}
	return evaluatedValue
}

func evaluateExpressionHelperAdvanced(expression []string) (int, []string) {
	unevaluatedExpressionQueue := make([]string, len(expression))
	copy(unevaluatedExpressionQueue, expression)
	evaluatedValue := 0
	currentOperator := ""

	for len(unevaluatedExpressionQueue) > 0 {
		currentElement := unevaluatedExpressionQueue[0]
		if currentElement == "(" {
			var outputValue int
			outputValue, unevaluatedExpressionQueue = evaluateExpressionHelperAdvanced(unevaluatedExpressionQueue[1:])
			evaluatedValue = evaluate(evaluatedValue, outputValue, currentOperator)
		} else if currentElement == ")" {
			return evaluatedValue, unevaluatedExpressionQueue
		} else if currentElement == "*" {
			currentOperator = currentElement
			var outputValue int
			outputValue, unevaluatedExpressionQueue = evaluateExpressionHelperAdvanced(unevaluatedExpressionQueue[1:])
			evaluatedValue = evaluate(evaluatedValue, outputValue, currentOperator)
			continue
		} else if currentElement == "+" {
			currentOperator = currentElement
		} else {
			intValue, error := strconv.Atoi(currentElement)
			if error != nil {
				log.Fatalf("Couldn't convert %v to int\n", currentElement)
			}
			evaluatedValue = evaluate(evaluatedValue, intValue, currentOperator)
		}

		unevaluatedExpressionQueue = unevaluatedExpressionQueue[1:]
	}

	return evaluatedValue, unevaluatedExpressionQueue
}

func evaluateExpressionHelper(expression []string) (int, []string) {
	unevaluatedExpressionQueue := make([]string, len(expression))
	copy(unevaluatedExpressionQueue, expression)
	evaluatedValue := 0
	currentOperator := ""

	for len(unevaluatedExpressionQueue) > 0 {
		currentElement := unevaluatedExpressionQueue[0]
		if currentElement == "(" {
			var outputValue int
			outputValue, unevaluatedExpressionQueue = evaluateExpressionHelper(unevaluatedExpressionQueue[1:])
			evaluatedValue = evaluate(evaluatedValue, outputValue, currentOperator)
		} else if currentElement == ")" {
			return evaluatedValue, unevaluatedExpressionQueue
		} else if currentElement == "*" || currentElement == "+" {
			currentOperator = currentElement
		} else {
			intValue, error := strconv.Atoi(currentElement)
			if error != nil {
				log.Fatalf("Couldn't convert %v to int\n", currentElement)
			}
			evaluatedValue = evaluate(evaluatedValue, intValue, currentOperator)
		}

		unevaluatedExpressionQueue = unevaluatedExpressionQueue[1:]
	}

	return evaluatedValue, unevaluatedExpressionQueue
}

func evaluate(evaluatedValue int, newValue int, operator string) int {
	if operator == "+" {
		return evaluatedValue + newValue
	} else if operator == "*" {
		return evaluatedValue * newValue
	} else if operator == "" && evaluatedValue == 0 {
		return newValue
	}

	log.Fatalf("Something weird happened with evaluate input, evaluatedValue: %v, newValue %v, operator: %v",
		evaluatedValue, newValue, operator)
	return -1
}

func parseInput(input []string) [][]string {
	expressions := [][]string{}

	for _, line := range input {
		updatedLine := strings.ReplaceAll(line, "(", "( ")
		updatedLine = strings.ReplaceAll(updatedLine, ")", " )")
		expressions = append(expressions, strings.Split(updatedLine, " "))
	}

	return expressions
}
