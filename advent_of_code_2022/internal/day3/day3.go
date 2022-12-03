package day3

import (
	"strings"

	"golang.org/x/exp/slices"
)

func Solve1(lines []string) (int, error) {
	items := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	sum := 0

	for _, sack := range lines {
		middleIndex := len(sack) / 2
		firstComp := sack[0:middleIndex]
		secComp := sack[middleIndex:]

		for _, item := range firstComp {
			found := strings.ContainsRune(secComp, item)
			if found {
				index := slices.Index(items, item)

				// Account for 0 based indexing
				sum += index + 1
				break
			}
		}
	}

	return sum, nil
}

func Solve2(lines []string) (int, error) {
	items := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	sum := 0

	group := make([]string, 0, 3)

	for _, sack := range lines {
		group = append(group, sack)

		if len(group) == 3 {
			for _, item := range group[0] {
				found1 := strings.ContainsRune(group[1], item)
				found2 := strings.ContainsRune(group[2], item)
				if found1 && found2 {
					index := slices.Index(items, item)

					// Account for 0 based indexing
					sum += index + 1
					group = make([]string, 0, 3)
					break
				}
			}
		}
	}

	return sum, nil
}
