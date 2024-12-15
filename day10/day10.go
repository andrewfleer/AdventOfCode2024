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

	visited bool
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
	trails := mapTrails(volcanoMap)

	fmt.Println(trails)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func mapTrails(volcanoMap [][]*Square) int {
	trails := 0
	for _, row := range volcanoMap {
		for _, square := range row {
			if square.height == 0 {
				trails += attemptToMapTrail(square)
				resetTrails(volcanoMap)
			}
		}
	}
	return trails
}

func resetTrails(volcanoMap [][]*Square) {
	for _, row := range volcanoMap {
		for _, square := range row {
			square.visited = false
		}
	}
}

func attemptToMapTrail(square *Square) int {
	height := square.height

	if height == 9 {
		if !square.visited {
			square.visited = true
			return 1
		}

		return 0
	}

	trails := 0

	if square.north != nil &&
		square.north.height == height+1 {
		trails += attemptToMapTrail(square.north)
	}

	if square.east != nil &&
		square.east.height == height+1 {
		trails += attemptToMapTrail(square.east)
	}

	if square.south != nil &&
		square.south.height == height+1 {
		trails += attemptToMapTrail(square.south)
	}

	if square.west != nil &&
		square.west.height == height+1 {
		trails += attemptToMapTrail(square.west)
	}

	return trails
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
