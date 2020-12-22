package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Turns [][][]int

func (t Turns) isEqualToPrevious(d1, d2 []int) bool {
	for _, previousTurn := range t {
		prevDeck1, prevDeck2 := previousTurn[0], previousTurn[1]

		if SliceEqual(d1, prevDeck1) && SliceEqual(d2, prevDeck2) {
			return true
		}
	}
	return false
}

func main() {
	in := ReadFromInput()
	p1, p2 := Parse(in)

	part1, _ := PlayGame(p1, p2, false)
	fmt.Println("Part 1:", part1)
	part2, _ := PlayGame(p1, p2, true)
	fmt.Println("Part 2:", part2)
}

func CalculateScore(deck []int) (score int) {
	for idx, card := range deck {
		multiplier := len(deck) - idx
		score += card * multiplier
	}
	return
}

func PlayGame(d1, d2 []int, part2 bool) (int, bool) {
	turns := Turns{}
	for len(d1) > 0 && len(d2) > 0 {
		if part2 {
			if turns.isEqualToPrevious(d1, d2) {
				return CalculateScore(d1), true
			} else {
				turns = append(turns, [][]int{d1, d2})
			}
		}
		d1, d2 = Round(d1, d2, part2)
	}

	if len(d1) > 0 {
		return CalculateScore(d1), true
	}
	return CalculateScore(d2), false
}

func Round(d1, d2 []int, part2 bool) ([]int, []int) {
	cardd1, restDeck1 := d1[0], d1[1:]
	cardd2, restDeck2 := d2[0], d2[1:]

	if part2 && cardd1 <= len(restDeck1) && cardd2 <= len(restDeck2) {
		// must copy or we get memory overwrites
		copyDeck1 := make([]int, cardd1)
		copyDeck2 := make([]int, cardd2)
		copy(copyDeck1, restDeck1)
		copy(copyDeck2, restDeck2)

		_, d1Wins := PlayGame(copyDeck1, copyDeck2, part2)

		if d1Wins {
			return append(restDeck1, cardd1, cardd2), restDeck2
		} else {
			return restDeck1, append(restDeck2, cardd2, cardd1)
		}

	} else if cardd1 > cardd2 {
		return append(restDeck1, cardd1, cardd2), restDeck2
	} else {
		return restDeck1, append(restDeck2, cardd2, cardd1)
	}
}

func Parse(in []string) ([]int, []int) {
	player1Deck := []int{}
	player2Deck := []int{}
	isPlayer1 := true
	for _, line := range in {
		if line == "Player 1:" || line == "" {
			continue
		}
		if line == "Player 2:" {
			isPlayer1 = false
			continue
		}
		n, _ := strconv.Atoi(line)
		if isPlayer1 {
			player1Deck = append(player1Deck, n)
		} else {
			player2Deck = append(player2Deck, n)
		}
	}
	return player1Deck, player2Deck
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func SliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for idx, item := range s1 {
		if s2[idx] != item {
			return false
		}
	}
	return true
}
