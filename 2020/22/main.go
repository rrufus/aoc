package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()
	// in := ReadFromStdIn()

	p1, p2 := Parse(in)

	fmt.Println("Part 1")
	fmt.Println(CalculateScorePart1(PlayGamePart1(p1, p2)))

	fmt.Println("Part 2")

}

func CalculateScorePart1(deck []int) (score int) {
	for idx, card := range deck {
		multiplier := len(deck) - idx
		score += card * multiplier
	}
	return
}

func PlayGamePart1(d1, d2 []int) []int {

	for len(d1) > 0 && len(d2) > 0 {
		d1, d2 = Round(d1, d2)
	}

	if len(d1) > 0 {
		return d1
	} else {
		return d2
	}
}

func Round(d1, d2 []int) ([]int, []int) {
	cardd1, restDeck1 := d1[0], d1[1:]
	cardd2, restDeck2 := d2[0], d2[1:]

	if cardd1 > cardd2 {
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

func ReadFromStdIn() []string {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)

read_loop:
	for {
		text, _ := reader.ReadString('\n')
		if text == "go\n" {
			break read_loop
		}
		lines = append(lines, strings.TrimSpace(text))
	}

	return lines
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
