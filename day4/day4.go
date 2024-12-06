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
	NULL = iota
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
	totalMatches := findXMAS(letterCoordinates)

	fmt.Println(totalMatches)
}

func findXMAS(letterCoordinates map[rune][]Coordinate) int {
	matches := 0
	aCoords := letterCoordinates[rune('A')]

	possibleLetters := []rune{'M', 'S'}
	for _, coords := range aCoords {

		if isValid(letterCoordinates, coords, possibleLetters, NORTHWEST) {
			if isValid(letterCoordinates, coords, possibleLetters, SOUTHWEST) {
				matches++
			}
		}

	}

	return matches
}

func isValid(letterCoordinates map[rune][]Coordinate, coords Coordinate, possibleLetters []rune, direction int) bool {
	var nextDirection int
	nextCoords := Coordinate{
		x: coords.x,
		y: coords.y,
	}

	switch direction {
	case NULL:
		return true
	case NORTHEAST:
		nextCoords.y--
		nextCoords.x++
		nextDirection = NULL
	case SOUTHEAST:
		nextCoords.x++
		nextCoords.y++
		nextDirection = NULL
	case SOUTHWEST:
		nextCoords.y++
		nextCoords.x--
		nextDirection = NORTHEAST
	case NORTHWEST:
		nextCoords.y--
		nextCoords.x--
		nextDirection = SOUTHEAST
	}

	for _, letter := range possibleLetters {
		letterCoords := letterCoordinates[letter]

		for _, c := range letterCoords {
			if c.x == nextCoords.x &&
				c.y == nextCoords.y {

				remainingLetters := []rune{'S'}
				if letter == 'S' {
					remainingLetters = []rune{'M'}
				}
				return isValid(letterCoordinates, coords, remainingLetters, nextDirection)
			}
		}
	}

	return false
}
