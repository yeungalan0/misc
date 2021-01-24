package day23

import (
	"fmt"
	"log"
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
	// Note: don't run this when you try 1,000,000 cups or tests will timeout
	// printCircularLinkedList(cupCircle)

	cupsCircleAfterPlay := simulateCrabCups(rounds, cups[0], cupCircle)
	nodeWith1Val, e := getNodeByVal(1, cupsCircleAfterPlay)
	if e != nil {
		log.Fatalf("Error finding value in list: val: %v, error: %v\n", 1, e)
	}

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

func simulateCrabCups(rounds, currCupVal int, cupCircle circularLinkedList) circularLinkedList {
	currNode, e := getNodeByVal(currCupVal, cupCircle)
	if e != nil {
		log.Fatalf("Error finding value in list: val: %v, error: %v\n", currCupVal, e)
	}

	for rounds > 0 {
		closeNodeList := getCloseNodeList(currNode)
		destinationNode := getDestinationNode(currNode, closeNodeList, cupCircle)

		// Splice in close cups
		currNode.next = closeNodeList.tail.next
		destinationNodeNext := destinationNode.next
		destinationNode.next = closeNodeList.head
		closeNodeList.tail.next = destinationNodeNext

		currNode = currNode.next
		rounds--
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
		if destinationNodeVal < 1 {
			destinationNodeVal = len(cupCircle.nodes)
		}

		if !valueInList(destinationNodeVal, closeNodeList) {
			destinationNode, e := getNodeByVal(destinationNodeVal, cupCircle)
			if e != nil {
				log.Fatalf("Error finding value in list: val: %v, error: %v\n", destinationNode, e)
			}
			
			return destinationNode
		}
	}
}

func valueInList(val int, list linkedList) bool {
	currNode := list.head
	for currNode != nil {
		if list.tail.next != nil && currNode.val == list.tail.next.val {
			break
		} else if currNode.val == val {
			return true
		}

		currNode = currNode.next
	}

	return false
}

func getNodeByVal(val int, cupCircle circularLinkedList) (*node, error) {
	i := val - 1
	if i < 0 || i >= len(cupCircle.nodes) {
		return nil, fmt.Errorf("value not in list")
	}
	return cupCircle.nodes[val - 1], nil
}