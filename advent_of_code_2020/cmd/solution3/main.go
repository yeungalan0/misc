package main

import (
	"fmt"
	"internal/day3"
	"internal/utils"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("input")

	if err != nil {
		log.Fatalf("Couldn't open file: %v\n", err)
	}

	list, err := utils.ReadStrings(reader)

	treesHitRight1Down1, _ := day3.TreesHit(list, day3.Slope{Right: 1, Down: 1})
	treesHitRight3Down1, _ := day3.TreesHit(list, day3.Slope{Right: 3, Down: 1})
	treesHitRight5Down1, _ := day3.TreesHit(list, day3.Slope{Right: 5, Down: 1})
	treesHitRight7Down1, _ := day3.TreesHit(list, day3.Slope{Right: 7, Down: 1})
	treesHitRight1Down2, _ := day3.TreesHit(list, day3.Slope{Right: 1, Down: 2})

	fmt.Printf("Problem 1: trees hit %v\n", treesHitRight3Down1)
	fmt.Printf("Problem 2: total trees hit product %v\n", treesHitRight1Down1*treesHitRight3Down1*treesHitRight5Down1*treesHitRight7Down1*treesHitRight1Down2)
}
