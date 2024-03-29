package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// in := ReadFromInput()
	in := ReadFromStdIn()

	fmt.Println("Part 1")

	fmt.Println("Part 2")

}

func ReadFromStdIn() []string {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)

read_loop:
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			break read_loop
		}
		lines = append(lines, strings.TrimSpace(text))
	}

	return lines
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}
	return ints
}

func SliceEqual(s1, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}
	for idx, item := range s1 {
		if s2[idx] != item {
			return false
		}
	}
	return true
}

func ReadFromInput() []string {
	bytes, _ := os.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
