package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	in := ReadFromInput()
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
	currentCoord := Coord{}
	currentBestInRange := 0
	n := []int{-1, 0, 1}
	multiplier := 1
	multiplierFactor := 100000000

	startTime := time.Now()
main_loop:
	for {
		nBelow := 0
		for _, offsetX := range n {
			for _, offsetY := range n {
				for _, offsetZ := range n {
					newCoord := Coord{currentCoord.X + multiplier*offsetX, currentCoord.Y + multiplier*offsetY, currentCoord.Z + multiplier*offsetZ}
					inRange := calculateNInRange(newCoord, bots)
					if inRange < currentBestInRange {
						nBelow++
					} else if inRange > currentBestInRange {
						currentCoord = newCoord
						currentBestInRange = inRange
						multiplier = 1
						continue main_loop
					}
				}
			}
		}
		// if everything in every direction is less.
		if nBelow == 26 {
			if multiplierFactor == 1 {
				break
			}
			multiplierFactor /= 10
			multiplier = 1
		} else {
			multiplier += multiplierFactor
		}
	}
	endTime := time.Now()

	fmt.Println(abs(currentCoord.X) + abs(currentCoord.Y) + abs(currentCoord.Z)) // 160646364
	fmt.Println("Part 2 took", endTime.Sub(startTime))
}

type Coord struct {
	X, Y, Z int
}

type Bot struct {
	Coord
	R int
}

func calculateNInRange(c Coord, bots []*Bot) (total int) {
	for _, b := range bots {
		if b.inRange(c) {
			total++
		}
	}
	return
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

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
