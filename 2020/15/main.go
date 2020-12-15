package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	initialState := StringsToInts(strings.Split(in[0], ","))
	mem := map[int]int{}
	for idx, turn := range initialState {
		mem[turn] = idx
	}
	PlayGame(initialState, 2020, 30000000, mem)

}

func PlayGame(turnHistory []int, targetTurnPart1, targetTurnPart2 int, mem map[int]int) {
	for len(turnHistory) < targetTurnPart2 {
		lastTurn := len(turnHistory) - 1
		lastSpoken := turnHistory[lastTurn]
		previousTurnSameNumber, exists := mem[lastSpoken]
		if exists {
			diff := lastTurn - previousTurnSameNumber
			mem[lastSpoken] = len(turnHistory) - 1
			turnHistory = append(turnHistory, diff)
		} else {
			mem[lastSpoken] = lastTurn
			turnHistory = append(turnHistory, 0)
		}
		if len(turnHistory) == targetTurnPart1 {
			fmt.Println("Part 1:", turnHistory[targetTurnPart1-1])
		}
	}
	fmt.Println("Part 2:", turnHistory[targetTurnPart2-1])
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
