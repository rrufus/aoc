package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	valid, myTicket := ComputeValidValuesAndMyTicket(in)
	ticketsToCheck := strings.Split(strings.Split(strings.Join(in, "\n"), "nearby tickets:\n")[1], "\n")
	validTickets, errors := DiscardErrors(ticketsToCheck, valid)
	fieldPositions := GetFieldPositions(validTickets, valid)

	fmt.Println("Part 1")
	fmt.Println(errors)
	fmt.Println("Part 2")
	fmt.Println(myTicket[fieldPositions["departure location"]] *
		myTicket[fieldPositions["departure station"]] *
		myTicket[fieldPositions["departure platform"]] *
		myTicket[fieldPositions["departure track"]] *
		myTicket[fieldPositions["departure date"]] *
		myTicket[fieldPositions["departure time"]])
}

func GetFieldPositions(validTickets [][]int, valid map[string]map[int]bool) map[string]int {
	byPosition := map[int][]int{}
	positions := map[string]int{}
	possibleFields := map[int][]string{}

	for _, ticket := range validTickets {
		for idx, n := range ticket {
			if _, exists := byPosition[idx]; !exists {
				byPosition[idx] = []int{n}
			} else {
				byPosition[idx] = append(byPosition[idx], n)
			}
			possibleFields[idx] = []string{}
		}
	}
	for position, values := range byPosition {
		for field, validValues := range valid {
			validForPosition := true
			for _, observedValue := range values {
				if !validValues[observedValue] {
					validForPosition = false
				}
			}
			if validForPosition {
				possibleFields[position] = append(possibleFields[position], field)
			}
		}
	}
	for !hasOnePossibleFieldPerPosition(possibleFields) {
		for position, fields := range possibleFields {
			if len(fields) == 1 {
				field := fields[0]
				for p, f := range possibleFields {
					if p != position {
						possibleFields[p] = remove(f, field)
					}
				}
			}
		}
	}

	for idx, fields := range possibleFields {
		field := fields[0]
		positions[field] = idx
	}

	return positions
}

func remove(in []string, item string) []string {
	out := []string{}
	for _, str := range in {
		if str != item {
			out = append(out, str)
		}
	}
	return out
}

func hasOnePossibleFieldPerPosition(possibleFields map[int][]string) bool {
	for _, v := range possibleFields {
		if len(v) != 1 {
			return false
		}
	}
	return true
}

func DiscardErrors(tickets []string, valid map[string]map[int]bool) ([][]int, int) {
	errors := 0
	remaining := [][]int{}
	for _, ticket := range tickets {
		ints := StringsToInts(strings.Split(ticket, ","))
		ticketHasError := false
		for _, currentInt := range ints {
			validInt := false
			for _, rangeOfInts := range valid {
				if rangeOfInts[currentInt] {
					validInt = true
				}
			}
			if !validInt {
				ticketHasError = true
				errors += currentInt
			}
		}
		if !ticketHasError {
			remaining = append(remaining, ints)
		}
	}
	return remaining, errors
}

func ComputeValidValuesAndMyTicket(in []string) (map[string]map[int]bool, []int) {
	validValues := map[string]map[int]bool{}

	for idx, line := range in {
		if line == "" {
			continue
		}
		if line == "your ticket:" {
			return validValues, StringsToInts(strings.Split(in[idx+1], ","))
		}

		field := strings.Split(line, ": ")[0]
		secondPart := strings.Split(line, ": ")[1]
		ranges := strings.Split(secondPart, " or ")
		for _, rangeStr := range ranges {
			startAndEnd := strings.Split(rangeStr, "-")
			start, _ := strconv.Atoi(startAndEnd[0])
			end, _ := strconv.Atoi(startAndEnd[1])
			for i := start; i <= end; i++ {
				if _, exists := validValues[field]; exists {
					validValues[field][i] = true
				} else {
					validValues[field] = map[int]bool{i: true}
				}
			}
		}
	}
	return validValues, []int{}
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
