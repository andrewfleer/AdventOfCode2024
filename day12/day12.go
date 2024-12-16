package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Square struct {
	plotType rune
	north    *Square
	east     *Square
	south    *Square
	west     *Square

	x, y    int
	fences  int
	visited bool
}

func main() {
	file, err := os.Open("../files/day12/day12.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	gardenMap := [][]*Square{}
	rowNum := 0
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune(line)

		// buildMap
		gardenMap = buildMap(gardenMap, row, rowNum)

		rowNum++
	}

	fences := calculateFenceCost(gardenMap)

	fmt.Println(fences)
}

func calculateFenceCost(gardenMap [][]*Square) int {
	fenceCost := 0
	for _, row := range gardenMap {
		for _, square := range row {
			if !square.visited {
				plotFences, area := makePlot(square, square.plotType)

				fenceCost += (plotFences * area)
			}
		}
	}

	return fenceCost
}

func makePlot(square *Square, plotType rune) (int, int) {
	fences := 0
	area := 0

	if square != nil &&
		!square.visited &&
		square.plotType == plotType {
		square.visited = true

		// check north
		nFences, nArea := makePlot(square.north, square.plotType)
		fences += nFences
		area += nArea

		// check east
		eFences, eArea := makePlot(square.east, square.plotType)
		fences += eFences
		area += eArea

		// check south
		sFences, sArea := makePlot(square.south, square.plotType)
		fences += sFences
		area += sArea

		// check west
		wFences, wArea := makePlot(square.west, square.plotType)
		fences += wFences
		area += wArea

		// add this plot
		fences += square.fences
		area++
	}
	return fences, area
}

func buildMap(gardenMap [][]*Square, row []rune, rowNum int) [][]*Square {
	mapRow := []*Square{}

	for i, r := range row {
		square := &Square{plotType: r, fences: 4, y: rowNum, x: i}

		// set North
		if rowNum != 0 {
			link(square, gardenMap[rowNum-1][i], "north")
		}

		// set West
		if i != 0 {
			link(square, mapRow[i-1], "west")
		}

		mapRow = append(mapRow, square)
	}

	gardenMap = append(gardenMap, mapRow)

	return gardenMap
}

func link(from *Square, to *Square, direction string) {
	switch direction {
	case "north":
		from.north = to
		to.south = from

		// Remove fences if plots match
		if to.plotType == from.plotType {
			from.fences--
			to.fences--
		}
	case "west":
		from.west = to
		to.east = from

		// Remove fences if plots match
		if to.plotType == from.plotType {
			from.fences--
			to.fences--
		}
	}
}
