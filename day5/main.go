package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	START_ROWS    = 128
	START_COLUMNS = 8
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

func FindRow(rowIdentifier string) int {
	rangeStart := 1
	rangeEnd := START_ROWS
	for _, character := range rowIdentifier {
		if character == 'F' {
			rangeEnd -= (rangeEnd-rangeStart)/2 + 1
		}
		if character == 'B' {
			rangeStart += (rangeEnd-rangeStart)/2 + 1
		}
	}

	if rowIdentifier[len(rowIdentifier)-1] == 'F' {
		return rangeStart - 1
	}
	return rangeEnd - 1
}

func FindColumn(columnIdentifier string) int {
	rangeStart := 1
	rangeEnd := START_COLUMNS
	for _, character := range columnIdentifier {
		if character == 'L' {
			rangeEnd -= (rangeEnd-rangeStart)/2 + 1
		}
		if character == 'R' {
			rangeStart += (rangeEnd-rangeStart)/2 + 1
		}
	}

	if columnIdentifier[len(columnIdentifier)-1] == 'L' {
		return rangeStart - 1
	}
	return rangeEnd - 1
}

func main() {
	in := ReadFromStdIn()

	fmt.Println("Part 1")
	maxSeatId := 0
	seatIds := []int{}
	for _, line := range in {
		rowIdentifier := line[0:7]
		columnIdentifier := line[7:]

		row := FindRow(rowIdentifier)
		column := FindColumn(columnIdentifier)

		seatId := row*8 + column
		seatIds = append(seatIds, seatId)

		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	fmt.Println(maxSeatId)

	fmt.Println("Part 2")

	sort.Ints(seatIds)

	for idx, seatId := range seatIds {
		if idx != 0 && idx != len(seatIds)-1 {
			if seatIds[idx+1] != seatId+1 {
				fmt.Println(seatId + 1)
			}
		}
	}

}
