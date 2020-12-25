package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	publicKeys := StringsToInts(in)

	cardLoops := FindLoopSize(7, publicKeys[0])
	encryptionKey := Transform(1, publicKeys[1], cardLoops)

	fmt.Println("Part 1:", encryptionKey)
	fmt.Println("Merry Christmas!")
}

func FindLoopSize(start, out int) int {
	nLoops := 0
	output := 7
	currentValue := 1
	for {
		if currentValue == out {
			return nLoops
		}
		currentValue = Transform(currentValue, output, 1)
		nLoops++
	}
}

func Transform(in, subjectNumber, nTimes int) int {

	for i := 0; i < nTimes; i++ {
		in = (in * subjectNumber) % 20201227
	}

	return in
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
