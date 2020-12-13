package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bus struct {
	FirstTime, RepeatsEvery int
}

func main() {
	in := ReadFromInput()

	earliestArrivalTime, buses := Parse(in)

	fmt.Println("Part 1")
	fmt.Println(Part1(earliestArrivalTime, buses))

	fmt.Println("Part 2")
	fmt.Println(Part2(buses))
}

func Parse(in []string) (int, []*Bus) {
	earliestArrivalTime, _ := strconv.Atoi(in[0])
	busesStr := strings.Split(in[1], ",")
	buses := []*Bus{}

	for idx, busStr := range busesStr {
		if busStr == "x" {
			continue
		}
		every, _ := strconv.Atoi(busStr)
		buses = append(buses, &Bus{idx, every})
	}

	return earliestArrivalTime, buses
}

func Part1(earliestArrivalTime int, buses []*Bus) int {
	time := earliestArrivalTime

	for {
		for _, bus := range buses {
			if time%bus.RepeatsEvery == 0 {
				return (time - earliestArrivalTime) * bus.RepeatsEvery
			}
		}
		time++
	}
}

// Part2 has got to be a mathematical solution as "surely the actual earliest timestamp will be larger than 100,000,000,000,000"
func Part2(buses []*Bus) int {

	foldedMultiple, rest := buses[0], buses[1:]

	for _, m := range rest {
		foldedMultiple = FoldPair(foldedMultiple, m)
	}
	return foldedMultiple.FirstTime

}

// FoldPair should be able to find a way of combining two multiples which
// results in another with the combined properties of the first
func FoldPair(memo, new *Bus) (m *Bus) {

	repeatsEvery := LCM(memo.RepeatsEvery, new.RepeatsEvery)

	for i := 0; true; i++ {
		time := i*memo.RepeatsEvery + memo.FirstTime
		if time < 0 {
			continue
		}
		if (time+new.FirstTime)%new.RepeatsEvery == 0 {
			return &Bus{RepeatsEvery: repeatsEvery, FirstTime: time}
		}
	}
	return
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

// thanks google for the below...
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
