package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	x int
	y int
}

const (
	NORTH = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

func main() {
	file, err := os.Open("../files/day4/day4.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	letterCoordinates := make(map[rune][]Coordinate)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i, letter := range line {
			coords := letterCoordinates[letter]
			coord := Coordinate{
				x: i,
				y: lineCount,
			}

			coords = append(coords, coord)
			letterCoordinates[letter] = coords

		}

		lineCount++

	}

	//remove the first line

	//for _, row := range wordSearch {
	//	fmt.Printf("%q\n", row)
	//}
	totalMatches := findMatches(letterCoordinates, "XMAS")

	fmt.Println(totalMatches)
}

func findMatches(letterCoordinates map[rune][]Coordinate, word string) int {
	if len(word) == 0 {
		return 0
	}

	letterIndex := 0

	letter := word[letterIndex]

	firstLetterCoords := letterCoordinates[rune(letter)]

	if len(firstLetterCoords) == 0 {
		return 0
	}

	matches := 0
	for _, coords := range firstLetterCoords {
		if wordExists(letterCoordinates, letterIndex, word, coords, NORTH) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, NORTHEAST) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, EAST) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, SOUTHEAST) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, SOUTH) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, SOUTHWEST) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, WEST) {
			matches++
		}
		if wordExists(letterCoordinates, letterIndex, word, coords, NORTHWEST) {
			matches++
		}
	}
	return matches
}

func wordExists(letterCoordinates map[rune][]Coordinate, letterIndex int, word string, coords Coordinate, direction int) bool {
	letterIndex++
	if letterIndex >= len(word) {
		return true
	}

	letter := word[letterIndex]

	letterCoords := letterCoordinates[rune(letter)]

	nextCoords := Coordinate{
		x: coords.x,
		y: coords.y,
	}

	switch direction {
	case NORTH:
		nextCoords.y--
	case NORTHEAST:
		nextCoords.y--
		nextCoords.x++
	case EAST:
		nextCoords.x++
	case SOUTHEAST:
		nextCoords.x++
		nextCoords.y++
	case SOUTH:
		nextCoords.y++
	case SOUTHWEST:
		nextCoords.y++
		nextCoords.x--
	case WEST:
		nextCoords.x--
	case NORTHWEST:
		nextCoords.y--
		nextCoords.x--
	}

	for _, c := range letterCoords {
		if c.x == nextCoords.x && c.y == nextCoords.y {
			return wordExists(letterCoordinates, letterIndex, word, c, direction)
		}
	}

	return false

}
