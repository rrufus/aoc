package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1:", CalculateFinalManhattanDistance(in, image.Point{X: 1, Y: 0}, false))

	fmt.Println("Part 2:", CalculateFinalManhattanDistance(in, image.Point{X: 10, Y: 1}, true))
}

func CalculateFinalManhattanDistance(instructions []string, heading image.Point, part2 bool) int {
	shipCoord := image.Point{}
	for _, instruction := range instructions {
		action := rune(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])

		coordToUpdate := &shipCoord
		if part2 {
			coordToUpdate = &heading
		}

		switch action {
		case 'N':
			coordToUpdate.Y += amount
		case 'S':
			coordToUpdate.Y -= amount
		case 'E':
			coordToUpdate.X += amount
		case 'W':
			coordToUpdate.X -= amount
		case 'L':
			fallthrough
		case 'R':
			heading = RotateCoord(heading, action, amount/90)
		case 'F':
			shipCoord.X += heading.X * amount
			shipCoord.Y += heading.Y * amount
		}
	}
	return int(math.Abs(float64(shipCoord.X)) + math.Abs(float64(shipCoord.Y)))
}

func RotateCoord(in image.Point, direction rune, times int) image.Point {
	rotateDirection := map[rune]int{'L': -1, 'R': 1}
	newPosition := image.Point{X: in.X, Y: in.Y}

	for i := 0; i < times; i++ {
		newPosition = image.Point{
			X: rotateDirection[direction] * newPosition.Y,
			Y: -rotateDirection[direction] * newPosition.X,
		}
	}
	return newPosition
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
