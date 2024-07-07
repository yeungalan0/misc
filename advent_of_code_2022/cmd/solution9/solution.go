package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
	"log"

	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"
)

//go:embed input
var s string

func main() {
	lines := strings.Split(s, "\n")

	solution1, err := solve1(lines)
	if err != nil {
		log.Fatalf("Error solving day: %v\n", err)
	}

	fmt.Printf("Solution 1 to day: %v\n", solution1)

	solution2, err := solve2(lines)
	if err != nil {
		log.Fatalf("Error solving day: %v\n", err)
	}

	fmt.Printf("Solution 2 to day: %v\n", solution2)
}

type coord struct {
	x int
	y int
}

func solve1(lines []string) (int, error) {
	currCoordHead := coord{x: 0, y: 0}
	currCoordTail := currCoordHead
	// Add in starting coord
	beenTo := map[coord]bool{currCoordTail: true}

	for _, line := range lines {
		motions := strings.Split(line, " ")

		direction, mag := motions[0], utils.HandleErr(strconv.Atoi(motions[1]))

		for i := 0; i < mag; i++ {
			// fmt.Printf("CURR: %#v, %#v\n", currCoordHead, currCoordTail)
			currCoordHead, currCoordTail = simulate(currCoordHead, currCoordTail, direction)
			// diffInX := math.Abs(math.Abs(float64(currCoordHead.x))-math.Abs(float64(currCoordTail.x))) > 1
			// diffInY := math.Abs(math.Abs(float64(currCoordHead.y))-math.Abs(float64(currCoordTail.y))) > 1

			// if diffInX || diffInY {
			// 	fmt.Printf("ERROR: %#v, %#v\n", currCoordHead, currCoordTail)
			// }

			beenTo[currCoordTail] = true
		}
	}

	// fmt.Printf("BEENTO: %#v", beenTo)

	return len(beenTo), nil
}

func simulate(head, tail coord, direction string) (coord, coord) {
	switch direction {
	case "R":
		head.x += 1
	case "L":
		head.x -= 1
	case "U":
		head.y += 1
	case "D":
		head.y -= 1
	}

	xMovement := 0
	yMovement := 0

	// Diagonal movement is necessary
	diffXY := head.x != tail.x && head.y != tail.y
	diffInX := math.Abs(float64(head.x)-float64(tail.x)) > 1
	diffInY := math.Abs(float64(head.y)-float64(tail.y)) > 1

	if diffXY && (diffInX || diffInY) {
		xMovement = int(math.Round(float64(head.x-tail.x) / float64(2.0)))
		yMovement = int(math.Round(float64(head.y-tail.y) / float64(2.0)))
	} else if diffInX {
		xMovement = int(math.Round(float64(head.x-tail.x) / float64(2.0)))
	} else if diffInY {
		yMovement = int(math.Round(float64(head.y-tail.y) / float64(2.0)))
	}

	tail.x += xMovement
	tail.y += yMovement

	return head, tail
}

func solve2(lines []string) (int, error) {
	knots := []coord{
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
		coord{x: 0, y: 0},
	}

	tailIdx := len(knots) - 1
	// Add in starting coord
	beenTo := map[coord]bool{knots[tailIdx]: true}

	for _, line := range lines {
		motions := strings.Split(line, " ")

		direction, mag := motions[0], utils.HandleErr(strconv.Atoi(motions[1]))

		for i := 0; i < mag; i++ {
			for j := 0; j < len(knots)-1; j++ {
				currHead := knots[j]
				currTail := knots[j+1]

				if j == 0 {
					knots[j], knots[j+1] = simulate(currHead, currTail, direction)
				} else {
					knots[j], knots[j+1] = simulate(currHead, currTail, "noOp")
				}

			}

			beenTo[knots[tailIdx]] = true
		}
	}

	// fmt.Printf("BEENTO: %#v", beenTo)

	return len(beenTo), nil
}
