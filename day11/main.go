package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EMPTY    = 'L'
	OCCUPIED = '#'
	FLOOR    = '.'
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
	iterationsPart1 := [][]string{in}

	for {
		currentIteration := iterationsPart1[len(iterationsPart1)-1]
		nextIteration := make([]string, len(in))

		for colIdx, row := range currentIteration {
			for rowIdx, seat := range row {

				if seat == EMPTY && countAdjacentOccupied(currentIteration, colIdx, rowIdx) == 0 {
					nextIteration[colIdx] += string(OCCUPIED)
				} else if seat == OCCUPIED && countAdjacentOccupied(currentIteration, colIdx, rowIdx) >= 4 {
					nextIteration[colIdx] += string(EMPTY)
				} else {
					nextIteration[colIdx] += string(seat)
				}
			}
		}

		iterationsPart1 = append(iterationsPart1, nextIteration)

		// this is an assumption that numbers not changing means positions not changing
		if countTotalOccupied(currentIteration) == countTotalOccupied(nextIteration) {
			fmt.Println(countTotalOccupied(nextIteration))
			break
		}
	}

	fmt.Println("Part 2")
	iterationsPart2 := [][]string{in}

	for {
		currentIteration := iterationsPart2[len(iterationsPart2)-1]
		nextIteration := make([]string, len(in))

		for colIdx, row := range currentIteration {
			for rowIdx, seat := range row {

				if seat == EMPTY && countVisibleOccupied(currentIteration, colIdx, rowIdx) == 0 {
					nextIteration[colIdx] += string(OCCUPIED)
				} else if seat == OCCUPIED && countVisibleOccupied(currentIteration, colIdx, rowIdx) > 4 {
					nextIteration[colIdx] += string(EMPTY)
				} else {
					nextIteration[colIdx] += string(seat)
				}
			}
		}

		iterationsPart2 = append(iterationsPart2, nextIteration)

		// this is an assumption that numbers not changing means positions not changing
		if countTotalOccupied(currentIteration) == countTotalOccupied(nextIteration) {
			fmt.Println(countTotalOccupied(nextIteration))
			break
		}

	}
}

func countAdjacentOccupied(seatPlan []string, colIdx, rowIdx int) int {
	coords := [][]int{
		{colIdx - 1, rowIdx - 1},
		{colIdx - 1, rowIdx},
		{colIdx - 1, rowIdx + 1},
		{colIdx, rowIdx - 1},
		{colIdx, rowIdx + 1},
		{colIdx + 1, rowIdx - 1},
		{colIdx + 1, rowIdx},
		{colIdx + 1, rowIdx + 1}}

	occupied := 0

	for _, coord := range coords {
		col, row := coord[0], coord[1]
		if col < 0 || row < 0 || col >= len(seatPlan) || row >= len(seatPlan[0]) {
			continue
		}
		if seatPlan[col][row] == OCCUPIED {
			occupied++
		}
	}

	return occupied
}

func countVisibleOccupied(seatPlan []string, colIdx, rowIdx int) int {
	vectors := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	visible := 0

	// think about caching result, ie from this coordinate, we already calculated that this is visible.
	for _, vector := range vectors {
		dy, dx := vector[0], vector[1]
		nextCol := colIdx
		nextRow := rowIdx

		for {
			nextCol += dy
			nextRow += dx

			if nextCol < 0 || nextCol >= len(seatPlan) || nextRow < 0 || nextRow >= len(seatPlan[0]) || seatPlan[nextCol][nextRow] == EMPTY {
				// reached the edge or can see empty seat in this direction. Count this as a 0
				break
			}
			if seatPlan[nextCol][nextRow] == OCCUPIED {
				visible++
				break
			}
		}
	}
	return visible
}

func countTotalOccupied(rows []string) (count int) {
	for _, row := range rows {
		count += strings.Count(row, string(OCCUPIED))
	}
	return
}
