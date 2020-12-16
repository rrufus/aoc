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
	in := ReadFromInput()
	// in := ReadFromStdIn()

	valid := ComputeValidValues(in)
	ticketsToCheck := strings.Split(strings.Split(strings.Join(in, "\n"), "nearby tickets:\n")[1], "\n")
	fmt.Println("Part 1")
	fmt.Println(CalculateErrorRate(ticketsToCheck, valid))

	fmt.Println("Part 2")

}

func CalculateErrorRate(tickets []string, valid map[int]bool) int {
	errors := 0
	for _, ticket := range tickets {
		ints := StringsToInts(strings.Split(ticket, ","))
		for _, currentInt := range ints {
			if _, exists := valid[currentInt]; !exists {
				errors += currentInt
			}
		}
	}
	return errors
}

func ComputeValidValues(in []string) map[int]bool {
	validValues := map[int]bool{}

	for _, line := range in {
		if line == "" {
			break
		}
		secondPart := strings.Split(line, ": ")[1]
		ranges := strings.Split(secondPart, " or ")
		for _, rangeStr := range ranges {
			startAndEnd := strings.Split(rangeStr, "-")
			start, _ := strconv.Atoi(startAndEnd[0])
			end, _ := strconv.Atoi(startAndEnd[1])
			for i := start; i <= end; i++ {
				validValues[i] = true
			}
		}
	}
	return validValues
}

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
