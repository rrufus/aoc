package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// in := ReadFromInput()
	in := ReadFromStdIn()

	fmt.Println("Part 1")
	turnHistory := StringsToInts(strings.Split(in[0], ","))

	fmt.Println(PlayGame(turnHistory, 2020, map[int]int{}))

	fmt.Println("Part 2")

}

func PlayGame(turnHistory []int, targetTurn int, mem map[int]int) int {
	for len(turnHistory) < targetTurn {
		lastSpoken := turnHistory[len(turnHistory)-1]
		turnsMentioned := FindTurns(turnHistory, lastSpoken)
		if len(turnsMentioned) == 1 {
			turnHistory = append(turnHistory, 0)
		} else {
			turnHistory = append(turnHistory, turnsMentioned[len(turnsMentioned)-1]-turnsMentioned[len(turnsMentioned)-2])
		}
	}
	return turnHistory[targetTurn-1]
}

func FindTurns(ints []int, intToFind int) []int {
	foundIndex := []int{}

	for idx, i := range ints {
		if i == intToFind {
			foundIndex = append(foundIndex, idx+1)
		}
	}
	return foundIndex
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

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
