package day13

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type busIndex struct {
	busID int
	index int
}

type linearCongruence struct {
	modulo    int64
	remainder int64
}

// FindEarliestBus returns the bus ID and time that you could take the earliest bus given the input
func FindEarliestBus(input []string) (int, int) {
	earliestDepartureTime, buses := parseInput(input)
	earliestBus := -1
	earliestBusTime := math.MaxInt32

	for _, myBusIndex := range buses {
		multiplyBy := int(math.Ceil(float64(earliestDepartureTime) / float64(myBusIndex.busID)))

		bestBusTime := myBusIndex.busID * multiplyBy
		if bestBusTime < earliestBusTime {
			earliestBus = myBusIndex.busID
			earliestBusTime = bestBusTime
		}
	}

	return earliestBus, earliestBusTime
}

// FindSynchronousEarliestTime returns the earliest time t such that bus 0 leaves at time t, bus 1 leaves at time t+1, etc.
func FindSynchronousEarliestTime(input []string) int64 {
	linearCongruences := parseInputToLinearCongruence(input)

	minimumTime, _ := chineseRemainderGauss(linearCongruences)

	return minimumTime
}

func parseInput(input []string) (int, []busIndex) {
	buses := []busIndex{}
	earliestTimestamp, error := strconv.Atoi(input[0])
	if error != nil {
		log.Fatalf("Couldn't convert to int: %v", input[0])
	}

	busSlice := strings.Split(input[1], ",")
	for index, busString := range busSlice {
		if busString == "x" {
			continue
		}

		busID, error := strconv.Atoi(busString)
		if error != nil {
			log.Fatalf("Couldn't convert to int: %v", busString)
		}

		buses = append(buses, busIndex{busID: busID, index: index})
	}

	return earliestTimestamp, buses
}

func parseInputToLinearCongruence(input []string) []linearCongruence {
	linearCongruences := []linearCongruence{}

	busSlice := strings.Split(input[1], ",")
	for index, busString := range busSlice {
		if busString == "x" {
			continue
		}

		busID, error := strconv.Atoi(busString)
		if error != nil {
			log.Fatalf("Couldn't convert to int: %v", busString)
		}

		linearCongruences = append(linearCongruences, linearCongruence{modulo: int64(busID), remainder: int64(busID - index%busID)})
	}

	return linearCongruences
}

// 1. Check modulos are coprime
// 2. Calculate N
// 3. sum += a1 * (N/n) * invmod(N/n, n)

// chineseRemainderGauss uses the Gaussian method to return the minimum X and it's modulo
// satisfying a system of linear congruences
func chineseRemainderGauss(linearCongruences []linearCongruence) (int64, int64) {
	if !areCoPrime(linearCongruences) {
		log.Fatalf("Linear congruences don't contain all coprime modulo, CRT does not apply! %#v", linearCongruences)
	}

	sum := int64(0)
	N := int64(1)

	for _, linCongruence := range linearCongruences {
		N *= linCongruence.modulo
	}

	for _, linCongruence := range linearCongruences {
		currentModulo := N / linCongruence.modulo
		sum += linCongruence.remainder * currentModulo * findInverseMod(currentModulo, linCongruence.modulo)
	}

	sum %= N // Get the minimum positive answer
	return sum, N
}

func areCoPrime(linearCongruences []linearCongruence) bool {
	for index, linCongruence := range linearCongruences {
		for _, linCongruence2 := range linearCongruences[index+1:] {
			if gcd, _, _ := extendedEuclid(linCongruence.modulo, linCongruence2.modulo); gcd != 1 {
				return false
			}
		}
	}

	return true
}

// euclidean algorithm to find GCD
// func greatestCommonDivisor(a int64, b int64) int64 {
// 	if a == 0 {
// 		return b
// 	} else if b == 0 {
// 		return a
// 	}

// 	remainder := a % b

// 	return greatestCommonDivisor(b, remainder)
// }

func findInverseMod(a int64, m int64) int64 {
	gcd, x, _ := extendedEuclid(a, m)

	if gcd != 1 {
		log.Fatalf("GCD is not 1, modular inverse doesn't exist for %v (mod %v), gcd: %v", a, m, gcd)
	}
	if x < 0 {
		x = m + x
	}

	return x % m
}

// extendedEuclid takes x and y, and returns gcd(x,y) with a and b such that a*x + b*y = gcd(x,y)
// TODO: Look into this algorithm more, definitely don't understand it as well as I should
func extendedEuclid(x int64, y int64) (int64, int64, int64) {
	x0, x1, y0, y1 := int64(1), int64(0), int64(0), int64(1)
	var q int64

	for y > 0 {
		q, x, y = x/y, y, x%y
		x0, x1 = x1, x0-q*x1
		y0, y1 = y1, y0-q*y1
	}

	return int64(x), int64(x0), int64(y0)
}
