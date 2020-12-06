package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFromStdIn() []string {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)

read_loop:
	for {
		text, _ := reader.ReadString('\n')
		if text == "go\n" {
			break read_loop
		}
		lines = append(lines, strings.TrimSpace(text))
	}

	return lines
}

func main() {
	in := ReadFromStdIn()

	fmt.Println("Part 1")
	groupsPart1 := []map[rune]bool{}

	groupsPart1 = append(groupsPart1, map[rune]bool{})
	for _, line := range in {
		activeGroup := groupsPart1[len(groupsPart1)-1]

		if line == "" {
			groupsPart1 = append(groupsPart1, map[rune]bool{})
		} else {
			for _, character := range line {
				activeGroup[character] = true
			}
		}
	}
	count := 0
	for _, group := range groupsPart1 {
		count += len(group)
	}
	fmt.Println(count)

	fmt.Println("Part 2")
	groupsPart2 := []map[rune]bool{}
	passengersInGroup := []map[rune]bool{}

	groupsPart2 = append(groupsPart2, map[rune]bool{})
	for _, line := range in {
		activeGroup := groupsPart2[len(groupsPart2)-1]

		if line == "" {
			firstPassenger, rest := passengersInGroup[0], passengersInGroup[1:]

		passenger_answer_loop:
			for question := range firstPassenger {
				for _, passenger := range rest {
					if !passenger[question] {
						continue passenger_answer_loop
					}
				}
				activeGroup[question] = true
			}

			// reset passengers
			passengersInGroup = []map[rune]bool{}
			// new group
			groupsPart2 = append(groupsPart2, map[rune]bool{})
		} else {
			passenger := map[rune]bool{}
			for _, character := range line {
				passenger[character] = true
			}
			passengersInGroup = append(passengersInGroup, passenger)
		}
	}
	count = 0
	for _, group := range groupsPart2 {
		count += len(group)
	}
	fmt.Println(count)

}
