package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Cup struct {
	prev *Cup
	next *Cup
	val  int
}

func main() {
	in := ReadFromInput()
	cupInts := StringsToInts(strings.Split(in[0], ""))

	cups := make([]*Cup, len(cupInts))
	for idx, _ := range cups {
		cups[idx] = &Cup{}
	}
	var cupContainingOne *Cup
	for idx, cupInt := range cupInts {
		cups[idx].val = cupInt
		cups[idx].prev = cups[(idx-1+len(cupInts))%len(cupInts)]
		cups[idx].next = cups[(idx+1)%len(cupInts)]
		if cupInt == 1 {
			cupContainingOne = cups[idx]
		}
	}

	lowest := cupInts[0]
	highest := cupInts[0]
	for _, cupInt := range cupInts {
		if cupInt < lowest {
			lowest = cupInt
		}
		if cupInt > highest {
			highest = cupInt
		}
	}

	currentCup := cups[0]

	fmt.Println(lowest, highest)

	fmt.Println("Part 1")
	for i := 0; i < 100; i++ {
		isolatedCupsStart := currentCup.next
		isolatedCupsEnd := currentCup.next.next.next

		fmt.Printf("-- move %v --\n", i+1)
		fmt.Println("Pick up:", isolatedCupsStart.val, isolatedCupsStart.next.val, isolatedCupsEnd.val)

		currentCup.next = currentCup.next.next.next.next
		currentCup.next.prev = currentCup

		nLookingFor := currentCup.val - 1
		for nLookingFor == isolatedCupsStart.val ||
			nLookingFor == isolatedCupsStart.next.val ||
			nLookingFor == isolatedCupsEnd.val || nLookingFor < lowest {
			nLookingFor--
			if nLookingFor < lowest {
				nLookingFor = highest
			}
		}
		destination := currentCup.next
		for destination.val != nLookingFor {
			destination = destination.next
		}
		fmt.Println("destination:", destination.val)
		oldNext := destination.next
		isolatedCupsEnd.next = oldNext
		oldNext.prev = isolatedCupsEnd
		destination.next = isolatedCupsStart
		isolatedCupsStart.prev = destination

		currentCup = currentCup.next
	}
	fmt.Println("-- final --")
	cursor := cupContainingOne.next
	for i := 0; i < len(cups)-1; i++ {
		fmt.Printf("%v", cursor.val)
		cursor = cursor.next
	}
	fmt.Println("\n")

	fmt.Println("Part 2")

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
