package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	Forward = "forward"
	Down    = "down"
	Up      = "up"
)

type DivePosition struct {
	Horizontal int
	Depth      int
	Aim        int
}

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1")
	position := &DivePosition{}

	for _, line := range in {
		instruction, amountStr := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		amount, _ := strconv.Atoi(amountStr)

		if instruction == Forward {
			position.Horizontal += amount
		}
		if instruction == Down {
			position.Depth += amount
		}
		if instruction == Up {
			position.Depth -= amount
		}
	}
	fmt.Println(position.Depth * position.Horizontal)

	fmt.Println("Part 2")
	position = &DivePosition{}
	for _, line := range in {
		instruction, amountStr := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		amount, _ := strconv.Atoi(amountStr)

		if instruction == Forward {
			position.Horizontal += amount
			position.Depth += position.Aim * amount
		}
		if instruction == Down {
			position.Aim += amount
		}
		if instruction == Up {
			position.Aim -= amount
		}
	}
	fmt.Println(position.Depth * position.Horizontal)

}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
