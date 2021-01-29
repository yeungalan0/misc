package day25

import (
	"log"
	"math"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

const mod = 20201227

// GetEncryptionKey returns the calculated encryption key given the two public keys and subject number
func GetEncryptionKey(publicKeyStrings []string, subjectNumber int) int {
	publicKeys := utils.ConvertStringListToIntList(publicKeyStrings)

	loopSize1, exponentToRemainderMap := getLoopSize(publicKeys[0], subjectNumber, map[int]int{})
	loopSize2, _ := getLoopSize(publicKeys[1], subjectNumber, exponentToRemainderMap)

	encryptionKey1, _ := runTransformation(loopSize1, publicKeys[1], map[int]int{})
	encryptionKey2, _ := runTransformation(loopSize2, publicKeys[0], map[int]int{})

	if encryptionKey1 == encryptionKey2 {
		return encryptionKey1
	} 

	log.Fatalf("Sanity check failed: key1: %v, key2: %v", encryptionKey1, encryptionKey2)
	return -1
}

func getLoopSize(publicKey, subjectNumber int, exponentToRemainderMap map[int]int) (int, map[int]int) {
	loopSize := 0

	for {
		remainder, exponentToRemainderMap := runTransformation(loopSize, subjectNumber, exponentToRemainderMap)
		if remainder == publicKey {
			return loopSize, exponentToRemainderMap
		}
		loopSize++
	}
}

func runTransformation(loopSize, subjectNumber int, exponentToRemainderMap map[int]int) (int, map[int]int) {
	if remainder, isPresent := exponentToRemainderMap[loopSize]; isPresent {
		return remainder, exponentToRemainderMap
	}

	if loopSize == 1 {
		return subjectNumber % mod, exponentToRemainderMap
	}

	var closestLowPower2 int
	transformedNumber := 1

	for loopSize > 0 {
		if loopSize == 1 || loopSize == 2 {
			closestLowPower2 = 1
		} else {
			closestLowPower2 = int(math.Pow(2, math.Floor(math.Log2(float64(loopSize - 1)))))
		}
		remainder, exponentToRemainderMap := runTransformation(closestLowPower2, subjectNumber, exponentToRemainderMap)
		exponentToRemainderMap[closestLowPower2] = remainder

		transformedNumber *= remainder
		transformedNumber %= mod

		loopSize -= closestLowPower2
	}

	return transformedNumber, exponentToRemainderMap
}