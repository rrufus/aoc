package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(string(bytes), "\n")
}

func main() {
	in := ReadFromInput()
	groups := [][]string{}
	groups = append(groups, []string{})
	for _, line := range in {
		if line == "" {
			groups = append(groups, []string{})
		} else {
			groups[len(groups)-1] = append(groups[len(groups)-1], line)
		}
	}

	fmt.Println("Part 1")
	groupsPart1 := make([]map[rune]bool, len(groups))
	count := 0
	for idx, group := range groups {
		for _, line := range group {
			for _, character := range line {
				if groupsPart1[idx] == nil {
					groupsPart1[idx] = map[rune]bool{}
				}
				groupsPart1[idx][character] = true
			}
		}
	}
	for _, group := range groupsPart1 {
		count += len(group)
	}
	fmt.Println(count) // 6549

	fmt.Println("Part 2")
	groupsPart2 := make([]map[rune]bool, len(groups))
	for idx, group := range groups {
		if groupsPart2[idx] == nil {
			groupsPart2[idx] = map[rune]bool{}
		}
		if len(group) == 0 {
			continue
		}

		first := group[0]
		for _, character := range first {
			groupsPart2[idx][character] = true
		}
		if len(group) > 1 {
			rest := group[1:]

			for k := range groupsPart2[idx] {
				for _, line := range rest {
					if !strings.ContainsRune(line, k) {
						delete(groupsPart2[idx], k)
					}
				}
			}
		}
	}
	count = 0
	for _, group := range groupsPart2 {
		count += len(group)
	}
	fmt.Println(count) // 3466

}
