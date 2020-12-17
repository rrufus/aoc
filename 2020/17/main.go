package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type State rune
type Space map[Coord]State
type Space4D map[Coord4D]State

type Coord struct {
	X, Y, Z int
}

type Coord4D struct {
	W, X, Y, Z int
}

const (
	Active   = State('#')
	Inactive = State('.')
)

func main() {
	in := ReadFromInput()

	space := Space{}
	space4d := Space4D{}

	for yIdx, line := range in {
		for xIdx, cube := range line {
			c := Coord{xIdx, yIdx, 0}
			c4d := Coord4D{0, xIdx, yIdx, 0}
			space[c] = State(cube)
			space4d[c4d] = State(cube)
		}
	}

	cycles := 0
	for cycles < 6 {
		space = next(space)
		space4d = next4d(space4d)
		cycles++
	}
	fmt.Println("Part 1:", space.CountActiveCubes())
	fmt.Println("Part 2:", space4d.CountActiveCubes())
}

func next(oldSpace Space) (newSpace Space) {
	newSpace = Space{}

	for coord := range oldSpace {
		if _, exists := newSpace[coord]; !exists {
			newSpace[coord] = coord.nextState(oldSpace)
		}

		nextToCoords := coord.nextTo()
		for _, nextToCoord := range nextToCoords {
			if _, exists := newSpace[nextToCoord]; exists {
				continue
			}
			newSpace[nextToCoord] = nextToCoord.nextState(oldSpace)
		}
	}
	return
}

func next4d(oldSpace Space4D) (newSpace Space4D) {
	newSpace = Space4D{}

	for coord := range oldSpace {
		if _, exists := newSpace[coord]; !exists {
			newSpace[coord] = coord.nextState(oldSpace)
		}

		nextToCoords := coord.nextTo()
		for _, nextToCoord := range nextToCoords {
			if _, exists := newSpace[nextToCoord]; exists {
				continue
			}
			newSpace[nextToCoord] = nextToCoord.nextState(oldSpace)
		}
	}
	return
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func (s Space) CountActiveInCoords(coords []Coord) (total int) {
	for _, c := range coords {
		if state, exists := s[c]; exists && state == Active {
			total++
		}
	}
	return
}

func (s Space) CountActiveCubes() (total int) {
	for _, v := range s {
		if v == Active {
			total++
		}
	}
	return
}

func (s Space4D) CountActiveInCoords(coords []Coord4D) (total int) {
	for _, c := range coords {
		if state, exists := s[c]; exists && state == Active {
			total++
		}
	}
	return
}

func (s Space4D) CountActiveCubes() (total int) {
	for _, v := range s {
		if v == Active {
			total++
		}
	}
	return
}

func (c *Coord) nextTo() (result []Coord) {
	close := []int{-1, 0, 1}

	result = []Coord{}
	for _, x := range close {
		for _, y := range close {
			for _, z := range close {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				result = append(result, Coord{c.X + x, c.Y + y, c.Z + z})
			}
		}
	}
	return
}

func (c *Coord) nextState(space Space) State {
	currentState, exists := space[*c]
	if !exists {
		currentState = Inactive
	}
	activeNearby := space.CountActiveInCoords(c.nextTo())
	if currentState == Active {
		if activeNearby == 2 || activeNearby == 3 {
			return Active
		} else {
			return Inactive
		}
	} else if activeNearby == 3 {
		return Active
	} else {
		return Inactive
	}
}

func (c *Coord4D) nextTo() (result []Coord4D) {
	close := []int{-1, 0, 1}

	result = []Coord4D{}
	for _, w := range close {
		for _, x := range close {
			for _, y := range close {
				for _, z := range close {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					result = append(result, Coord4D{c.W + w, c.X + x, c.Y + y, c.Z + z})
				}
			}
		}
	}
	return
}

func (c *Coord4D) nextState(space Space4D) State {
	currentState, exists := space[*c]
	if !exists {
		currentState = Inactive
	}
	activeNearby := space.CountActiveInCoords(c.nextTo())
	if currentState == Active {
		if activeNearby == 2 || activeNearby == 3 {
			return Active
		} else {
			return Inactive
		}
	} else if activeNearby == 3 {
		return Active
	} else {
		return Inactive
	}
}
