package day7

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// CountPossibleContainerBags returns the number of bags that can contain the input bag given the rules
func CountPossibleContainerBags(rules []string, bag string, number int) int {
	containerBagsSet := map[string]bool{}
	bagsToContained := createRulesMap(rules)
	// Logic to loop through and keep counting containing bags
	containerBags := []string{bag}

	for len(containerBags) > 0 {
		containBag := containerBags[0]
		containerBags = containerBags[1:]

		for bag, containsMap := range bagsToContained {
			if containsMap[containBag] >= number && !containerBagsSet[bag] {
				containerBagsSet[bag] = true
				containerBags = append(containerBags, bag)
			}
		}
	}

	// fmt.Printf("containerBagsSet: %#v\n", containerBagsSet)
	return len(containerBagsSet)
}

// CountAllContainedBags returns the number of bags that need to be contained in the given bag based on the rules
func CountAllContainedBags(rules []string, bag string) int {
	bagsToContained := createRulesMap(rules)

	return recursiveCounter(bagsToContained, bag)
}

func recursiveCounter(bagsToContained map[string]map[string]int, bag string) int {
	containedBags := 0
	containedBagsMap := bagsToContained[bag]

	for containedBag, number := range containedBagsMap {
		containedBags += number
		containedBags += number * recursiveCounter(bagsToContained, containedBag)
	}

	return containedBags
}

func createRulesMap(rules []string) map[string]map[string]int {
	bagsToContained := map[string]map[string]int{}

	for _, rule := range rules {
		// Starting rule must be of form: "light red bags contain 1 bright white bag, 2 muted yellow bags."
		if strings.Contains(rule, "no other bags") {
			continue
		}

		// Convert to: "light red bags contain 1 bright white bag, 2 muted yellow"
		reEnd := regexp.MustCompile(" bag(s)?\\.")
		rule = reEnd.ReplaceAllString(rule, "")

		// bagKey: "light red", containedBags: []string{"1 bright white", "2 muted yellow"}
		reSplit := regexp.MustCompile(" bags contain | bag(s)?, ")
		rulesSlice := reSplit.Split(rule, -1)
		bagKey := rulesSlice[0]
		containedBagsSlice := rulesSlice[1:]

		containedBagToNumber := generateBagToNumber(containedBagsSlice)

		bagsToContained[bagKey] = containedBagToNumber
	}

	return bagsToContained
}

func generateBagToNumber(containedBagsSlice []string) map[string]int {
	containedBagToNumber := map[string]int{}

	for _, containedBag := range containedBagsSlice {
		numberAdjBag := strings.Split(containedBag, " ")
		if len(numberAdjBag) != 3 {
			log.Fatalf("Couldn't process rule, containedBag: %v, numberAdjBag: %#v", containedBag, numberAdjBag)
		}

		number, error := strconv.Atoi(numberAdjBag[0])
		if error != nil {
			log.Fatalf("Couldn't process as number: %v", numberAdjBag[0])
		}

		bagName := fmt.Sprintf("%v %v", numberAdjBag[1], numberAdjBag[2])
		containedBagToNumber[bagName] = number
	}

	return containedBagToNumber
}
