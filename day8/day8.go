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
			antinode := Coordinate{x: coord.x, y: coord.y}
			if !slices.Contains(antinodes, antinode) {
				antinodes = append(antinodes, antinode)
			}
			for j := i + 1; j < len(runes); j++ {
				nextCoord := runes[j]

				antinode = Coordinate{x: nextCoord.x, y: nextCoord.y}
				if !slices.Contains(antinodes, antinode) {
					antinodes = append(antinodes, antinode)
				}

				xDiff := math.Abs(float64(coord.x) - float64(nextCoord.x))
				yDiff := math.Abs(float64(coord.y) - float64(nextCoord.y))

				// go "behind" the first coord
				x := coord.x
				y := coord.y

				if coord.x < nextCoord.x {
					if coord.y < nextCoord.y {
						for {
							x -= int(xDiff)
							y -= int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x < 0 ||
								antinode.y < 0 {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					} else {
						for {
							x -= int(xDiff)
							y += int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x < 0 ||
								antinode.y >= mapHeight {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					}
				} else {
					if coord.y < nextCoord.y {
						for {
							x += int(xDiff)
							y -= int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x >= mapWidth ||
								antinode.y < 0 {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					} else {
						for {
							x += int(xDiff)
							y += int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x >= mapWidth ||
								antinode.y >= mapHeight {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					}
				}

				// go "past" the next node
				x = nextCoord.x
				y = nextCoord.y

				if nextCoord.x < coord.x {
					if nextCoord.y < coord.y {
						for {
							x -= int(xDiff)
							y -= int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x < 0 ||
								antinode.y < 0 {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					} else {
						for {
							x -= int(xDiff)
							y += int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x < 0 ||
								antinode.y >= mapHeight {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					}
				} else {
					if nextCoord.y < coord.y {
						for {
							x += int(xDiff)
							y -= int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x >= mapWidth ||
								antinode.y < 0 {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					} else {
						for {
							x += int(xDiff)
							y += int(yDiff)
							antinode.x = x
							antinode.y = y

							if antinode.x >= mapWidth ||
								antinode.y >= mapHeight {
								break
							}

							if !slices.Contains(antinodes, antinode) {
								antinodes = append(antinodes, antinode)
							}
						}
					}
				}
			}
		}
	}

	return antinodes
}
