package day14

import (
	"log"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

type addressToValue struct {
	address int
	value   int64
}

type bitmaskOperations struct {
	bitmask          [36]int          // Index is the position, and value is 0, 1, or 2 (representing both)
	addressesToValue []addressToValue // Memory address to value
}

func (myBMO bitmaskOperations) Equal(otherBMO bitmaskOperations) bool {
	return cmp.Equal(myBMO.bitmask, otherBMO.bitmask) &&
		cmp.Equal(myBMO.addressesToValue, otherBMO.addressesToValue)
}

func (aTV addressToValue) Equal(otherATV addressToValue) bool {
	return aTV.address == otherATV.address && aTV.value == otherATV.value
}

// CalculateBitmaskSum2 returns the sum of all values left in the memory address after running through input
// where mask 'X' values can affect multiple memory addresses
func CalculateBitmaskSum2(input []string) int64 {
	sum := int64(0)
	memoryAddressMap := map[int64]int64{}
	operations := parseInputToBitOperations(input)
	operations = utils.ReverseSlice(operations).([]bitmaskOperations)

	// TODO: Clean this up, these nested loops are pretty ugly
	for _, operation := range operations {
		for _, addressValue := range utils.ReverseSlice(operation.addressesToValue).([]addressToValue) {
			addresses := maskAddress(addressValue.address, operation.bitmask)
			for _, writeAddress := range addresses {
				// Since operations is reversed only run calculations for addresses not already in the map
				if _, isPresent := memoryAddressMap[writeAddress]; !isPresent {
					memoryAddressMap[writeAddress] = addressValue.value
				}
			}
		}
	}

	for _, value := range memoryAddressMap {
		sum += value
	}

	return sum
}

// CalculateBitmaskSum returns the sum of all values left in the memory address after running through input
func CalculateBitmaskSum(input []string) int64 {
	sum := int64(0)
	memoryAddressMap := map[int]int64{}
	operations := parseInputToBitOperations(input)
	operations = utils.ReverseSlice(operations).([]bitmaskOperations)

	for _, operation := range operations {
		for _, addressValue := range utils.ReverseSlice(operation.addressesToValue).([]addressToValue) {
			// Since operations is reversed only run calculations for addresses not already in the map
			if _, isPresent := memoryAddressMap[addressValue.address]; !isPresent {
				maskedValue := maskValue(addressValue.value, operation.bitmask, map[int]bool{2: true})
				memoryAddressMap[addressValue.address] = maskedValue
			}
		}
	}

	for _, value := range memoryAddressMap {
		sum += value
	}

	return sum
}

func parseInputToBitOperations(input []string) []bitmaskOperations {
	operations := []bitmaskOperations{}
	var currentBitmaskOperations *bitmaskOperations

	for _, maskOrOperation := range input {
		operationSlice := strings.Split(maskOrOperation, " = ")
		if operationSlice[0] == "mask" {
			if currentBitmaskOperations != nil {
				operations = append(operations, *currentBitmaskOperations)
			}
			bitmaskArray := convertBitmaskToArray(operationSlice[1])
			currentBitmaskOperations = &bitmaskOperations{
				bitmask:          bitmaskArray,
				addressesToValue: []addressToValue{},
			}
		} else {
			addressString := operationSlice[0][4 : len(operationSlice[0])-1]
			address, error := strconv.Atoi(addressString)
			if error != nil {
				log.Fatalf("couldn't convert %v to int!", addressString)
			}

			value, error := strconv.Atoi(operationSlice[1])
			if error != nil {
				log.Fatalf("couldn't convert %v to int!", operationSlice[1])
			}

			currentBitmaskOperations.addressesToValue = append(
				currentBitmaskOperations.addressesToValue,
				addressToValue{address: address, value: int64(value)},
			)
		}
	}

	operations = append(operations, *currentBitmaskOperations)

	return operations
}

func convertBitmaskToArray(bitmaskString string) [36]int {
	bitmaskArray := [36]int{}

	for index, maskValue := range bitmaskString {
		if maskValue == '0' {
			bitmaskArray[index] = 0
		} else if maskValue == '1' {
			bitmaskArray[index] = 1
		} else {
			bitmaskArray[index] = 2
		}
	}

	return bitmaskArray
}

// Note, position starts with the least significant bit on the right at 0, so right to left
func setBitValue(value int64, bitSet bool, position int) int64 {
	if bitSet {
		value |= (1 << position)
	} else {
		value &^= (1 << position)
	}
	return value
}

func maskValue(value int64, mask [36]int, ignoreBitValue map[int]bool) int64 {
	for index, bitValue := range mask {
		if !ignoreBitValue[bitValue] {
			leastSignificantIndex := len(mask) - 1 - index
			value = setBitValue(value, bitValue == 1, leastSignificantIndex)
		}
	}

	return value
}

func maskAddress(address int, mask [36]int) []int64 {
	addresses := []int64{}
	// Since we're dealing with 36 bit masks, a 32 bit int isn't big enough
	address64 := int64(address)
	// Get the address value with known mask bits set
	address64 = maskValue(address64, mask, map[int]bool{0: true, 2: true})

	addressesSet := map[int64]bool{address64: true}

	for index, bitValue := range mask {
		leastSignificantIndex := len(mask) - 1 - index
		if bitValue != 2 {
			continue
		}
		for addressKey := range addressesSet {
			bitSetAddress := setBitValue(addressKey, true, leastSignificantIndex)
			if !addressesSet[bitSetAddress] {
				addressesSet[bitSetAddress] = true
			}

			bitClearedAddress := setBitValue(addressKey, false, leastSignificantIndex)
			if !addressesSet[bitClearedAddress] {
				addressesSet[bitClearedAddress] = true
			}
		}
	}

	for addressKey := range addressesSet {
		addresses = append(addresses, addressKey)
	}

	return addresses
}
