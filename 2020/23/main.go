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
	cupInts := StringsToInts(strings.Split(ReadFromInput()[0], ""))

	cupsP1, lookupP1 := SetupCups(cupInts, len(cupInts))

	PlayCups(cupsP1, lookupP1, 100, false)

	fmt.Printf("Part 1: ")
	cursor := lookupP1[1].next
	for i := 0; i < len(cupsP1)-1; i++ {
		fmt.Printf("%v", cursor.val)
		cursor = cursor.next
	}
	fmt.Println("")

	cupsP2, lookupP2 := SetupCups(cupInts, 1000000)

	PlayCups(cupsP2, lookupP2, 10000000, false)

	fmt.Println("Part 2:", lookupP2[1].next.val*lookupP2[1].next.next.val)
}

func SetupCups(cupInts []int, n int) ([]*Cup, map[int]*Cup) {
	valsToPtrs := map[int]*Cup{}
	cups := make([]*Cup, n)
	for idx := range cups {
		cups[idx] = &Cup{}
	}

	for idx, cup := range cups {
		cupInt := idx + 1
		if idx < len(cupInts) {
			cupInt = cupInts[idx]
		}
		valsToPtrs[cupInt] = cup
		cup.val = cupInt
		cup.prev = cups[(idx-1+len(cups))%len(cups)]
		cup.next = cups[(idx+1)%len(cups)]
	}

	return cups, valsToPtrs
}

func PlayCups(cups []*Cup, lookup map[int]*Cup, rounds int, print bool) {
	currentCup := cups[0]
	lowest := 1
	highest := len(cups)
	for i := 0; i < rounds; i++ {
		isolatedCupsStart := currentCup.next
		isolatedCupsEnd := currentCup.next.next.next

		if print {
			fmt.Printf("-- move %v --\n", i+1)
			fmt.Println("pick up:", isolatedCupsStart.val, isolatedCupsStart.next.val, isolatedCupsEnd.val)
		}

		currentCup.next = currentCup.next.next.next.next
		currentCup.next.prev = currentCup

		nLookingFor := currentCup.val - 1
		for nLookingFor == isolatedCupsStart.val ||
			nLookingFor == isolatedCupsStart.next.val ||
			nLookingFor == isolatedCupsEnd.val ||
			nLookingFor < lowest {
			nLookingFor--
			if nLookingFor < lowest {
				nLookingFor = highest
			}
		}
		destination := lookup[nLookingFor]
		if print {
			fmt.Println("destination:", destination.val)
		}
		oldNext := destination.next
		isolatedCupsEnd.next = oldNext
		oldNext.prev = isolatedCupsEnd
		destination.next = isolatedCupsStart
		isolatedCupsStart.prev = destination

		currentCup = currentCup.next
	}
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
