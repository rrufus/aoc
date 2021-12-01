package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()
	ints := StringsToInts(in)

	fmt.Println("Part 1")

	fmt.Println(Part1(ints))

	fmt.Println("Part 2")
	sums := ThreeMeasurementSums(ints)
	fmt.Println(Part1(sums))
}

func Part1(ints []int) int {
	prev := ints[0]
	count := 0
	for _, n := range ints[1:] {
		if prev < n {
			count++
		}
		prev = n
	}

	return count
}

func ThreeMeasurementSums(ints []int) []int {
	sums := []int{}

	for idx := range ints {
		if idx < len(ints)-2 {
			sums = append(sums, ints[idx]+ints[idx+1]+ints[idx+2])
		}
	}
	return sums
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}
	return ints
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
