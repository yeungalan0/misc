package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"
	"log"

	"github.com/yeungalan0/misc/advent_of_code_2022/internal/utils"
)

type node struct {
	name     string
	size     int
	parent   *node
	children []*node
}

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

func solve1(lines []string) (int, error) {
	root := buildTree(lines)
	dfs(root)

	return sumDir(root), nil
}

func dfs(n *node) {
	// fmt.Printf("%#v\n", n)

	if len(n.children) == 0 {
		return
	}

	for _, child := range n.children {
		dfs(child)
	}

	size := 0
	for _, child := range n.children {
		size += child.size
	}

	n.size = size
}

func sumDir(n *node) int {
	sum := 0

	if len(n.children) == 0 {
		return sum
	}

	if n.size <= 100000 {
		sum += n.size
	}

	for _, child := range n.children {
		sum += sumDir(child)
	}

	return sum
}

func buildTree(lines []string) *node {
	root := &node{name: "/"}
	curr := root

	// Skip first line as root is already created
	for _, line := range lines[1:] {
		switch {
		case strings.Contains(line, "$ ls"):
			continue

		case strings.Contains(line, "dir "):
			dirName := strings.Split(line, " ")[1]
			newNode := &node{
				name:   dirName,
				parent: curr,
			}

			curr.children = append(curr.children, newNode)

		case strings.Contains(line, "$ cd .."):
			curr = curr.parent

		// Assume this is a cd to a specific directory
		case strings.Contains(line, "$ cd"):
			s := strings.Split(line, " ")
			nextDirName := s[2]

			for _, child := range curr.children {
				if child.name == nextDirName {
					curr = child

					continue
				}
			}

			// panic(fmt.Sprintf("Could not find child: %v, children: %#v, line: %v", nextDirName, curr.children, line))

		// When we get a file
		default:
			sizeAndName := strings.Split(line, " ")
			size, name := sizeAndName[0], sizeAndName[1]
			newNode := &node{
				name:   name,
				size:   utils.HandleErr(strconv.Atoi(size)),
				parent: curr,
			}

			curr.children = append(curr.children, newNode)
		}
	}

	for curr.parent != nil {
		curr = curr.parent
	}

	return curr
}

func sizes(n *node) []int {
	dirSizes := make([]int, 0)
	if len(n.children) == 0 {
		return dirSizes
	}

	dirSizes = append(dirSizes, n.size)

	for _, child := range n.children {
		dirSizes = append(dirSizes, sizes(child)...)
	}

	return dirSizes
}

func solve2(lines []string) (int, error) {
	root := buildTree(lines)
	dfs(root)

	used := root.size
	free := 70000000 - used
	needToFree := 30000000 - free
	// fmt.Println(needToFree)

	dirSizes := sizes(root)

	sort.Slice(dirSizes, func(i, j int) bool {
		return dirSizes[i] > dirSizes[j]
	})
	// fmt.Printf("%#v\n", dirSizes)

	prev := dirSizes[0]

	for _, size := range dirSizes {
		if size < needToFree {
			break
		}

		prev = size
	}

	return prev, nil
}
