package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	value int
	prev  []*Node
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func main() {
	ints := StringsToInts(ReadFromInput())

	startNode := &Node{value: 0, prev: []*Node{}}
	allNodes := []*Node{startNode}
	fmt.Println("Part 1")
	sort.Ints(ints)
	diffOf1 := 0
	diffOf3 := 1 // the device one
	currentJolts := 0
	for _, i := range ints {
		allNodes = append(allNodes, &Node{value: i, prev: []*Node{}})

		diff := i - currentJolts
		if diff == 1 {
			diffOf1++
		} else if diff == 3 {
			diffOf3++
		}
		currentJolts += diff

	}
	fmt.Println(diffOf1 * diffOf3)

	fmt.Println("Part 2")
	for idx, node := range allNodes {
		for _, nextNode := range allNodes[idx:] {
			if nextNode.value-node.value <= 3 && nextNode.value-node.value > 0 {
				nextNode.prev = append(nextNode.prev, node)
			}
		}
	}
	fmt.Println(FindPaths(allNodes[len(allNodes)-1], startNode, map[*Node]int{}))
}

// oops, this is O(N!) without cacheing and we have N=102
func FindPaths(node, goal *Node, cache map[*Node]int) int {
	paths := 0
	for _, prev := range node.prev {
		if prev == goal {
			paths++
		} else {
			if n, exists := cache[prev]; exists {
				paths += n
			} else {
				nPaths := FindPaths(prev, goal, cache)
				paths += nPaths
				cache[prev] = nPaths
			}
		}
	}
	return paths
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}
	return ints
}
