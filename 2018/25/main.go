package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	W, X, Y, Z int
}

type Point struct {
	Constellation *Constellation
	Coord
}

type Constellation map[*Point]bool

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (c *Coord) manhattan(c1 *Coord) int {
	return abs(c.W-c1.W) + abs(c.X-c1.X) + abs(c.Y-c1.Y) + abs(c.Z-c1.Z)
}

func NewPoint(w, x, y, z int) *Point {
	constellation := Constellation{}
	point := &Point{Coord: Coord{W: w, X: x, Y: y, Z: z}, Constellation: &constellation}
	constellation[point] = true
	return point
}

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1 (and only)")

	fmt.Println(CountConstellations(in))
}

func CountConstellations(in []string) int {
	points := []*Point{}
	constellationMaps := map[*Constellation]bool{}

	for _, line := range in {
		ints := StringsToInts(strings.Split(line, ","))
		point := NewPoint(ints[0], ints[1], ints[2], ints[3])

		matchingConstellations := map[*Constellation]bool{point.Constellation: true}
		for _, existingPoint := range points {
			if point.Coord.manhattan(&existingPoint.Coord) <= 3 {
				matchingConstellations[existingPoint.Constellation] = true
			}
		}
		if len(matchingConstellations) > 1 {
			newConstellation := Constellation{}
			for constellation := range matchingConstellations {
				for point := range *constellation {
					newConstellation[point] = true
					point.Constellation = &newConstellation
				}
				delete(constellationMaps, constellation)
			}
			constellationMaps[&newConstellation] = true
		}

		if len(matchingConstellations) == 1 {
			constellationMaps[point.Constellation] = true
		}

		points = append(points, point)
	}

	return len(constellationMaps)
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
