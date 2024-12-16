package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Square struct {
	plotType                       rune
	north, east, south, west       *Square
	nFence, eFence, sFence, wFence bool
	visited                        bool

	// For debugging
	x, y int
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
	totalCost := 0
	for _, row := range gardenMap {
		for _, square := range row {
			if !square.visited {
				corners, area := makePlot(square, square.plotType)
				plotCost := (corners * area)
				totalCost += plotCost
			}
		}
	}

	return totalCost
}

func makePlot(square *Square, plotType rune) (int, int) {
	corners := 0
	area := 0

	if square != nil &&
		!square.visited &&
		square.plotType == plotType {
		square.visited = true

		// check north
		nCorners, nArea := makePlot(square.north, square.plotType)
		corners += nCorners
		area += nArea

		// check east
		eCorners, eArea := makePlot(square.east, square.plotType)
		corners += eCorners
		area += eArea

		// check south
		sCorners, sArea := makePlot(square.south, square.plotType)
		corners += sCorners
		area += sArea

		// check west
		wCorners, wArea := makePlot(square.west, square.plotType)
		corners += wCorners
		area += wArea

		// add this plot
		area++

		// look for containing-corners
		if square.nFence {
			if square.eFence {
				corners++
			}

			if square.wFence {
				corners++
			}
		}

		if square.sFence {
			if square.eFence {
				corners++
			}
			if square.wFence {
				corners++
			}
		}

		// Look for extruding corners
		nSquare := square.north
		eSquare := square.east
		sSquare := square.south
		wSquare := square.west

		if nSquare != nil &&
			// Northeast
			nSquare.plotType == square.plotType {
			if eSquare != nil &&
				eSquare.plotType == square.plotType {
				if nSquare.eFence && eSquare.nFence {
					corners++
				}
			}
			// Northwest
			if wSquare != nil &&
				wSquare.plotType == square.plotType {
				if nSquare.wFence && wSquare.nFence {
					corners++
				}
			}
		}

		if sSquare != nil &&
			sSquare.plotType == square.plotType {
			// Southeast
			if eSquare != nil &&
				eSquare.plotType == square.plotType {
				if sSquare.eFence && eSquare.sFence {
					corners++
				}
			}

			// Southwest
			if wSquare != nil &&
				wSquare.plotType == square.plotType {
				if sSquare.wFence && wSquare.sFence {
					corners++
				}
			}
		}
	}
	return corners, area
}

func buildMap(gardenMap [][]*Square, row []rune, rowNum int) [][]*Square {
	mapRow := []*Square{}

	for i, r := range row {
		square := &Square{plotType: r, nFence: true, eFence: true, sFence: true, wFence: true, y: rowNum, x: i}

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
			from.nFence = false
			to.sFence = false
		}
	case "west":
		from.west = to
		to.east = from

		// Remove fences if plots match
		if to.plotType == from.plotType {
			from.wFence = false
			to.eFence = false
		}
	}
}
