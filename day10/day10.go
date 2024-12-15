package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Square struct {
	height int
	north  *Square
	east   *Square
	south  *Square
	west   *Square

	pathsToNine int
}

func main() {
	file, err := os.Open("../files/day10/day10.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	volcanoMap := [][]*Square{}
	rowNum := 0
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)

		// build Map
		volcanoMap = buildMap(volcanoMap, row, rowNum)

		rowNum++
	}

	// Map trails
	mapTrails(volcanoMap)

	trails := countTrails(volcanoMap)

	fmt.Println(trails)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countTrails(volcanoMap [][]*Square) int {
	trails := 0
	for _, row := range volcanoMap {
		for _, square := range row {
			if square.height == 0 {
				trails += square.pathsToNine
			}
		}
	}

	return trails
}

func mapTrails(volcanoMap [][]*Square) {
	for _, row := range volcanoMap {
		for _, square := range row {
			if square.height == 9 {
				attemptToMapTrail(square)
			}
		}
	}

}

func attemptToMapTrail(square *Square) {
	height := square.height

	if square.north != nil &&
		square.north.height == height-1 {
		square.north.pathsToNine += 1
		attemptToMapTrail(square.north)
	}

	if square.east != nil &&
		square.east.height == height-1 {
		square.east.pathsToNine += 1
		attemptToMapTrail(square.east)
	}

	if square.south != nil &&
		square.south.height == height-1 {
		square.south.pathsToNine += 1
		attemptToMapTrail(square.south)
	}

	if square.west != nil &&
		square.west.height == height-1 {
		square.west.pathsToNine += 1
		attemptToMapTrail(square.west)
	}
}

func buildMap(volcanoMap [][]*Square, row []rune, rowNum int) [][]*Square {
	mapRow := []*Square{}
	for i, r := range row {
		height, _ := strconv.Atoi(string(r))

		square := &Square{height: height}

		// set North
		if rowNum != 0 {
			link(square, volcanoMap[rowNum-1][i], "north")
		}

		// setWest
		if i != 0 {
			link(square, mapRow[i-1], "west")
		}

		mapRow = append(mapRow, square)

	}

	volcanoMap = append(volcanoMap, mapRow)

	return volcanoMap
}

func link(from, to *Square, direction string) {
	switch direction {
	case "north":
		from.north = to
		to.south = from
	case "west":
		from.west = to
		to.east = from
	}
}
