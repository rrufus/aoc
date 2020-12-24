package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type HexCoord struct {
	E, N int
}

func (h *HexCoord) adjacent() []HexCoord {
	return []HexCoord{{h.E + 2, h.N}, {h.E - 2, h.N}, {h.E + 1, h.N + 1}, {h.E + 1, h.N - 1}, {h.E - 1, h.N + 1}, {h.E - 1, h.N - 1}}
}

func main() {
	in := ReadFromInput()

	blackTiles := map[HexCoord]bool{}
	for _, line := range in {
		currCoord := &HexCoord{}
		currIdx := 0

		for currIdx < len(line) {
			if string(line[currIdx]) == "e" {
				currCoord.E += 2
				currIdx++
				continue
			}
			if string(line[currIdx]) == "w" {
				currCoord.E -= 2
				currIdx++
				continue
			}

			if string(line[currIdx:currIdx+2]) == "ne" {
				currCoord.N += 1
				currCoord.E += 1
			}
			if string(line[currIdx:currIdx+2]) == "sw" {
				currCoord.N -= 1
				currCoord.E -= 1
			}
			if string(line[currIdx:currIdx+2]) == "nw" {
				currCoord.N += 1
				currCoord.E -= 1
			}
			if string(line[currIdx:currIdx+2]) == "se" {
				currCoord.N -= 1
				currCoord.E += 1
			}
			currIdx = currIdx + 2
		}

		_, exists := blackTiles[*currCoord]
		if !exists {
			blackTiles[*currCoord] = true
		} else {
			delete(blackTiles, *currCoord)
		}
	}
	fmt.Println("Part 1", len(blackTiles))

	for i := 0; i < 100; i++ {
		nextBlackTiles := map[HexCoord]bool{}
		for coord := range blackTiles {
			adjacentTiles := coord.adjacent()

			nBlackAdjacent := 0
			for _, tile := range adjacentTiles {
				if blackTiles[tile] {
					nBlackAdjacent++
				} else {

					// handle white tiles turning black
					adjacentTileAdjacents := 0
					for _, nextTile := range tile.adjacent() {
						if blackTiles[nextTile] {
							adjacentTileAdjacents++
						}
					}
					if adjacentTileAdjacents == 2 {
						nextBlackTiles[tile] = true
					}
				}
			}

			// handle cases of black tiles stay black
			if nBlackAdjacent == 1 || nBlackAdjacent == 2 {
				nextBlackTiles[coord] = true
			}
		}
		blackTiles = nextBlackTiles
	}
	fmt.Println("Part 2:", len(blackTiles))
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
