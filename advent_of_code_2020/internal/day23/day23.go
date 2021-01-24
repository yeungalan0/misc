package day23

import (
	"fmt"
	"strings"

	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

type circularLinkedList struct {
	nodes []*node
}

type linkedList struct {
	head *node
	tail *node
}

type node struct {
	val int
	next *node
}

// GetCupLabels returns the labels (in clockwise order) except 1 of the cups after
// N rounds of crab cups
func GetCupLabels(input []string, rounds int, numberOfCups int) string {
	cups := utils.ConvertStringListToIntList(strings.Split(input[0], ""))
	cupCircle := getCircularLinkedList(cups, numberOfCups)
	printCircularLinkedList(cupCircle)

	cupsCircleAfterPlay := simulateCrabCups(rounds, cups[0], cupCircle)
	nodeWith1Val := getNodeByVal(1, cupsCircleAfterPlay)

	if numberOfCups > len(cups) {
		return fmt.Sprint(nodeWith1Val.next.val * nodeWith1Val.next.next.val)
	}

	labels := ""
	currNode := nodeWith1Val.next
	for currNode.val != 1 {
		labels += fmt.Sprint(currNode.val)
		currNode = currNode.next
	}

	return labels
}

func printCircularLinkedList(list circularLinkedList) {
	start := list.nodes[0].val
	output := fmt.Sprint(start)

	curr := *list.nodes[0].next
	for curr.val != start {
		output +=  " -> " + fmt.Sprint(curr.val)
		curr = *curr.next
	}

	fmt.Printf("%v\n", output)
}

func getCircularLinkedList(cups []int, numberOfCups int) circularLinkedList {
	circularList := circularLinkedList{nodes: make([]*node, numberOfCups)}

	var head *node
	var prev *node
	var curr *node

	for i, cup := range cups {
		curr = &node{val: cup}
		if i == 0 {
			head = curr
		} else {
			prev.next = curr
		}
		circularList.nodes[curr.val-1] = curr
		prev = curr
	}

	for i := len(cups); i < numberOfCups; i++ {
		curr = &node{val: i + 1}
		prev.next = curr
		circularList.nodes[i] = curr
		prev = curr
	}

	curr.next = head
	return circularList
}

func simulateCrabCups(rounds, currentCupVal int, cupCircle circularLinkedList) circularLinkedList {
	currNode := getNodeByVal(currentCupVal, cupCircle)

	for rounds > 0 {
		closeNodeList := getCloseNodeList(currNode)
		destinationNode := getDestinationNode(currNode, closeNodeList, cupCircle)

		// Splice in close cups
		currNode.next = closeNodeList.tail.next
		destinationNodeNext := destinationNode.next
		destinationNode.next = closeNodeList.head
		closeNodeList.tail.next = destinationNodeNext

		currNode = currNode.next
	}

	return cupCircle
}

func getCloseNodeList(currNode *node) linkedList {
	return linkedList{head: currNode.next, tail: currNode.next.next.next}
}

func getDestinationNode(currNode *node, closeNodeList linkedList, cupCircle circularLinkedList) *node {
	destinationNodeVal := currNode.val

	for {
		destinationNodeVal--
		if destinationNodeVal < 0 {
			destinationNodeVal = len(cupCircle.nodes)
		}

		if !valueInList(destinationNodeVal, closeNodeList) {
			return getNodeByVal(destinationNodeVal, cupCircle)
		}
	}
}

func valueInList(val int, list linkedList) bool {
	currNode := list.head
	for currNode != nil {
		if currNode.val == val {
			return true
		} else if list.tail.next != nil && currNode.val == list.tail.next.val {
			break
		}

		currNode = currNode.next
	}

	return false
}

func getNodeByVal(val int, cupCircle circularLinkedList) *node {
	return cupCircle.nodes[(val % len(cupCircle.nodes)) - 1]
}

// func simulateCrabCups(cups []int, rounds int) []int {
// 	numberOfCups := len(cups)
// 	newCups := make([]int, numberOfCups)
// 	currCupIndex := 0

// 	for rounds > 0 {
// 		offset := 0
// 		closeCupIndexes := getCupsWithCircularIndexes(currCupIndex + 1, currCupIndex + 4, numberOfCups)
// 		destinationCupIndex := getDestinationCupIndex(currCupIndex, closeCupIndexes, cups)

// 		newCups[0] = cups[currCupIndex]
// 		offset++
// 		newCups = fillSliceIndexes(offset, closeCupIndexes[2] + 1, destinationCupIndex + 1, cups, newCups)
// 		offset += circularIndexLength(closeCupIndexes[2] + 1, destinationCupIndex + 1, numberOfCups)
// 		newCups = fillSliceIndexes(offset, currCupIndex + 1, currCupIndex + 4, cups, newCups)
// 		offset += circularIndexLength(currCupIndex + 1, currCupIndex + 4, numberOfCups)
// 		newCups = fillSliceIndexes(offset, destinationCupIndex + 1, currCupIndex, cups, newCups)

// 		currCupIndex = getNewCurrCupIndex(cups[currCupIndex], newCups)
// 		cups, newCups = newCups, cups
// 		rounds--
// 	}

// 	return cups
// }

// func circularIndexLength(startIndex, endIndex, numberOfCups int) int {
// 	startIndex = startIndex % numberOfCups
// 	endIndex = endIndex % numberOfCups

// 	if endIndex < startIndex {
// 		endIndex += numberOfCups
// 	}

// 	return endIndex - startIndex
// }

// func fillSliceIndexes(offset, oldSliceStartIndex, oldSliceEndIndex int, oldSlice, newSlice []int) []int {
// 	sliceLen := len(oldSlice)
// 	oldSliceStartIndex = oldSliceStartIndex % sliceLen
// 	oldSliceEndIndex = oldSliceEndIndex % sliceLen

// 	if oldSliceEndIndex < oldSliceStartIndex {
// 		oldSliceEndIndex += sliceLen
// 	}

// 	for i := oldSliceStartIndex; i < oldSliceEndIndex; i++ {
// 		newSlice[(offset + i - oldSliceStartIndex) % sliceLen] = oldSlice[i % sliceLen]
// 	}

// 	return newSlice
// }

// func getValuesFromIndexes(slice, indexes []int) []int {
// 	values := []int{}
// 	for _, i := range indexes {
// 		values = append(values, slice[i])
// 	}

// 	return values
// }

// func getNewCurrCupIndex(currCupVal int, newCups []int) int {
// 	numberOfCups := len(newCups)
// 	currCupIndex := utils.SliceIndex(numberOfCups, func(i int) bool {return newCups[i] == currCupVal})
// 	newCurrCupIndex := (currCupIndex + 1) % numberOfCups
// 	return newCurrCupIndex
// }

// func getDestinationCupIndex(currCupIndex int, closeCupIndexes, cups []int) int {
// 	currCupVal := cups[currCupIndex]
// 	destinationIndex := -1

//     for destinationIndex == -1 {
// 		currCupVal--
// 		if currCupVal < 0 {
// 			currCupVal = 9
// 		}

// 		possibleIndex := utils.SliceIndex(len(cups), func(i int) bool {return cups[i] == currCupVal})
// 		isInCloseCupIndexes := utils.Contains(len(closeCupIndexes), func(i int) bool {return closeCupIndexes[i] == possibleIndex})

// 		if possibleIndex == currCupIndex {
// 			log.Fatalf("Looping detected! possbileIndex: %v, currCupIndex: %v", possibleIndex, currCupIndex)
// 		}

// 		if !isInCloseCupIndexes {
// 			destinationIndex = possibleIndex
// 		}
// 	}

// 	return destinationIndex
// }

// // getCupsWithCircularIndexes returns the indexes in cirular fashion inclusive of start and exclusive of end
// func getCupsWithCircularIndexes(startIndex, endIndex, numberOfCups int) []int {
// 	indexes := []int{}
// 	startIndex = startIndex % numberOfCups
// 	endIndex = endIndex % numberOfCups

// 	if endIndex < startIndex {
// 		endIndex += 9
// 	}

// 	for i := startIndex; i < endIndex; i++ {
// 		indexes = append(indexes, i % numberOfCups)
// 	}

// 	return indexes
// }