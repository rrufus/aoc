package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PREAMBLE_LENGTH = 25

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

func main() {
	ints := StringsToInts(ReadFromStdIn())

	fmt.Println("Part 1")
	cursor := PREAMBLE_LENGTH

	part1 := 0
main_loop:
	for {
		activeValues := ints[cursor-PREAMBLE_LENGTH : cursor]
		underInvestigation := ints[cursor]
		for _, v1 := range activeValues {
			for _, v2 := range activeValues {
				if v1 != v2 && v1+v2 == underInvestigation {
					cursor++
					continue main_loop

				}
			}
		}
		part1 = underInvestigation
		break
	}
	fmt.Println(part1)

	fmt.Println("Part 2")
	valuesRange := ints[:cursor]
	windowStart := 0
	windowEnd := 2
	for {
		currentWindow := valuesRange[windowStart:windowEnd]
		result := sum(currentWindow)
		if result < part1 {
			windowEnd++
		}
		if result == part1 {
			min, max := minmax(currentWindow)
			fmt.Println(min + max)
			break
		}
		if result > part1 {
			windowStart++
			windowEnd = windowStart + 2
		}
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func minmax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
