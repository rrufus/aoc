package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	in := ReadFromStdIn()

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
