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
	// in := ReadFromInput()
	in := ReadFromStdIn()
	bots := []*Bot{}

	fmt.Println("Part 1")
	highestSignalBot := &Bot{}
	for _, line := range in {
		bot := lineToBot(line)
		if bot.R > highestSignalBot.R {
			highestSignalBot = bot
		}
		bots = append(bots, bot)
	}
	inRange := 0
	for _, bot := range bots {
		if highestSignalBot.inRange(bot.Coord) {
			inRange++
		}
	}
	fmt.Println(inRange)

	fmt.Println("Part 2")

}

type Coord struct {
	X, Y, Z int
}

type Bot struct {
	Coord
	R int
}

func (b *Bot) inRange(c Coord) bool {
	dx := abs(c.X - b.X)
	dy := abs(c.Y - b.Y)
	dz := abs(c.Z - b.Z)

	return dx+dy+dz <= b.R
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func lineToBot(line string) *Bot {
	split := strings.Split(line, "=")
	numbers := strings.Trim(split[1], "<>, r")
	ints := strings.Split(numbers, ",")
	x, _ := strconv.Atoi(ints[0])
	y, _ := strconv.Atoi(ints[1])
	z, _ := strconv.Atoi(ints[2])
	r, _ := strconv.Atoi(split[2])
	return &Bot{Coord: Coord{x, y, z}, R: r}
}

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

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
