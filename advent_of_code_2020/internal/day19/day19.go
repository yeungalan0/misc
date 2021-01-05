package day19

import (
	"fmt"
	"regexp"
	"strings"
)

// NOTE: Definitely had inspiration from the Reddit thread in coming up with this answer. I could've solved this going
// down the path of using a large data structure, but this regex solution is much more optimal.
// https://www.reddit.com/r/adventofcode/comments/kg1mro/2020_day_19_solutions/

// CountMatchingRule returns a count of the number of messages that match the given rule
func CountMatchingRule(input []string, problem2 bool) int {
	count := 0
	ruleRegex, messages := parseInput(input, problem2)

	for _, message := range messages {
		if ruleRegex.MatchString(message) {
			count++
		}
	}

	return count
}

func parseInput(input []string, problem2 bool) (*regexp.Regexp, []string) {
	inputParts := [][]string{}

	part := []string{}
	for _, line := range input {
		if line == "" {
			inputParts = append(inputParts, part)
			continue
		}

		part = append(part, line)
	}
	inputParts = append(inputParts, part)

	rulesMap := buildRulesMap(inputParts[0])

	if problem2 {
		rule42Regex := buildRuleRegex("42", rulesMap)
		rule31Regex := buildRuleRegex("31", rulesMap)
		rule11 := ""

		rulesMap["8"] = []string{"(" + rule42Regex + "+" + ")"}
		for i := 1; i < 10; i++ { // Did this manually at first, but this is WAY smarter
			rule11 += fmt.Sprintf("|%s{%d}%s{%d}", rule42Regex, i, rule31Regex, i)
		}

		rulesMap["11"] = []string{"(" + rule11[1:] + ")"}
	}

	rulesRegex := buildRuleRegex("0", rulesMap)

	re := regexp.MustCompile("^" + rulesRegex + "$")

	return re, inputParts[1]
}

// TODO: Could cache built up rules in map here
func buildRuleRegex(rule string, rulesMap map[string][]string) string {
	re := ""
	rules := rulesMap[rule]

	if rules[0] == "\"a\"" || rules[0] == "\"b\"" || rules[0][0] == '(' {
		return strings.ReplaceAll(rules[0], "\"", "")
	}

	for _, rule := range rules {
		if rule == "|" {
			re += "|"
			continue
		}

		re += buildRuleRegex(rule, rulesMap)
	}

	return "(" + re + ")"
}

func buildRulesMap(ruleDefinitions []string) map[string][]string {
	rulesMap := map[string][]string{}

	for _, ruleDefinition := range ruleDefinitions {
		ruleDefinitionSlice := strings.Split(ruleDefinition, " ")
		ruleKeyString := strings.TrimSuffix(ruleDefinitionSlice[0], ":")

		rulesMap[ruleKeyString] = ruleDefinitionSlice[1:]
	}

	return rulesMap
}
