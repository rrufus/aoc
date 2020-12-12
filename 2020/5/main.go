package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1")
	seatIds := []int{}

	for _, line := range in {

		// added afterwards, didnt realise it was binary shift before.
		f := strings.Map(func(r rune) rune {
			switch r {
			case 'F', 'L':
				return '0'
			default:
				return '1'
			}
		}, line)
		seatId, _ := strconv.ParseInt(f, 2, 64)
		seatIds = append(seatIds, int(seatId))
	}
	sort.Ints(seatIds)
	fmt.Println(seatIds[len(seatIds)-1])

	fmt.Println("Part 2")

	for idx, seatId := range seatIds {
		if idx != 0 && idx != len(seatIds)-1 {
			if seatIds[idx+1] != seatId+1 {
				fmt.Println(seatId + 1)
			}
		}
	}

}
