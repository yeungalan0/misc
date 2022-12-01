package day1

import (
	"sort"
	"strconv"
)

func Solve1(lines []string) (int, error) {
	max := 0
	currSum := 0

	for _, line := range lines {
		if line == "" {
			if currSum > max {
				max = currSum
			}

			currSum = 0
			continue
		}

		intVal, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		currSum += intVal
	}

	return max, nil
}

func Solve2(lines []string) (int, error) {
	lines = append(lines, "") // Append newline to account for last set of calories
	sums := make([]int, 0)
	currSum := 0

	for _, line := range lines {
		if line == "" {
			sums = append(sums, currSum)

			currSum = 0
			continue
		}

		intVal, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		currSum += intVal
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	top3 := sums[0] + sums[1] + sums[2]

	return top3, nil
}
