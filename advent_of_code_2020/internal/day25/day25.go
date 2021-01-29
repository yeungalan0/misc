package day25

import (
	"log"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

// GetEncryptionKey returns the calculated encryption key given the two public keys and subject number
func GetEncryptionKey(publicKeyStrings []string, subjectNumber int) int {
	publicKeys := utils.ConvertStringListToIntList(publicKeyStrings)

	loopSize1 := getLoopSize(publicKeys[0], subjectNumber)
	loopSize2 := getLoopSize(publicKeys[1], subjectNumber)

	encryptionKey1 := runTransformation(loopSize1, publicKeys[0], subjectNumber)
	encryptionKey2 := runTransformation(loopSize2, publicKeys[1], subjectNumber)

	if encryptionKey1 == encryptionKey2 {
		return encryptionKey1
	} else {
		log.Fatalf("Sanity check failed: key1: %v, key2: %v", encryptionKey1, encryptionKey2)
	}

	return -1
}

func getLoopSize(publicKey, subjectNumber int) int {
	loopSize := 0

	
}