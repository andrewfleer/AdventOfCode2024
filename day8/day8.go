package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	file, err := os.Open("../files/day8/day8.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	runeMap := make(map[rune][]Coordinate)

	mapWidth := 0
	mapHeight := 0
	for scanner.Scan() {
		line := scanner.Text()
		mapWidth = len(line)

		runes := []rune(line)

		for i, r := range runes {
			if r != '.' {
				runeMap[r] = append(runeMap[r], Coordinate{x: i, y: mapHeight})
			}
		}

		mapHeight++

	}

	antinodes := findAntinodes(runeMap, mapWidth, mapHeight)

	fmt.Println(len(antinodes))
}

func findAntinodes(runeMap map[rune][]Coordinate, mapWidth, mapHeight int) []Coordinate {
	antinodes := []Coordinate{}

	for _, runes := range runeMap {
		if len(runes) < 2 {
			continue
		}

		for i, coord := range runes {
			for j := i + 1; j < len(runes); j++ {
				nextCoord := runes[j]

				xDiff := math.Abs(float64(coord.x) - float64(nextCoord.x))
				yDiff := math.Abs(float64(coord.y) - float64(nextCoord.y))

				var antinode1, antinode2 Coordinate

				if coord.x < nextCoord.x {
					antinode1.x = coord.x - int(xDiff)
					antinode2.x = nextCoord.x + int(xDiff)
				} else {
					antinode1.x = coord.x + int(xDiff)
					antinode2.x = nextCoord.x - int(xDiff)
				}

				if coord.y < nextCoord.y {
					antinode1.y = coord.y - int(yDiff)
					antinode2.y = nextCoord.y + int(yDiff)
				} else {
					antinode1.y = coord.y + int(yDiff)
					antinode2.y = nextCoord.y - int(yDiff)
				}

				// Check if the antinodes fit on the map.
				if antinode1.x >= 0 &&
					antinode1.x < mapWidth &&
					antinode1.y >= 0 &&
					antinode1.y < mapHeight {

					if !slices.Contains(antinodes, antinode1) {
						antinodes = append(antinodes, antinode1)
					}
				}

				if antinode2.x >= 0 &&
					antinode2.x < mapWidth &&
					antinode2.y >= 0 &&
					antinode2.y < mapHeight {

					if !slices.Contains(antinodes, antinode2) {
						antinodes = append(antinodes, antinode2)
					}
				}
			}
		}
	}

	return antinodes
}
