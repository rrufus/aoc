package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}

	return ints
}

func main() {
	ints := StringsToInts(ReadFromInput())

	fmt.Println("1")
loop_1:
	for _, i1 := range ints {
		for _, i2 := range ints {
			if i1+i2 == 2020 {
				fmt.Println(i1 * i2)
				break loop_1
			}
		}
	}

	fmt.Println("2")
loop_2:
	for _, i1 := range ints {
		for _, i2 := range ints {
			for _, i3 := range ints {
				if i1+i2+i3 == 2020 {
					fmt.Println(i1 * i2 * i3)
					break loop_2
				}
			}
		}
	}
}
