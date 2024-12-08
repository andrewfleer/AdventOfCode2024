package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Coordinate struct {
	x int
	y int
}

type Guard struct {
	coordinate Coordinate
	symbol     rune
}

func main() {
	file, err := os.Open("../files/day6/day6.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	labMap := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		runes := []rune(line)

		labMap = append(labMap, runes)
	}

	uniqueSquares := processMap(labMap)

	fmt.Println(uniqueSquares)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processMap(labMap [][]rune) int {
	guard := Guard{}

	for i, mapRow := range labMap {
		if guardLocation := findGuard(mapRow); guardLocation > -1 {
			guard.coordinate.x = guardLocation
			guard.coordinate.y = i
			guard.symbol = mapRow[guardLocation]
		}
	}

	uniqueSpots := patrol(labMap, guard)

	return uniqueSpots
}

func patrol(labMap [][]rune, guard Guard) int {
	guardOnScreen := true

	for {
		/*for row := 0; row < len(labMap); row++ {
			for column := 0; column < len(labMap[row]); column++ {
				fmt.Printf("%q ", labMap[row][column])
			}
			fmt.Print("\n")
		}

		fmt.Println("----------------------------------")*/

		labMap, guard, guardOnScreen = moveGuard(labMap, guard)

		if !guardOnScreen {
			break
		}
	}

	uniqueSpots := countSpots(labMap)

	return uniqueSpots
}

func moveGuard(labMap [][]rune, guard Guard) ([][]rune, Guard, bool) {
	nextSpot := getIntendedMovement(guard)

	if nextSpotOffScreen(labMap, nextSpot) {
		labMap[guard.coordinate.y][guard.coordinate.x] = 'X'

		return labMap, guard, false
	}

	if labMap[nextSpot.y][nextSpot.x] == '#' {
		guard.symbol = turnGuard(guard.symbol)
		labMap[guard.coordinate.y][guard.coordinate.x] = guard.symbol

		return labMap, guard, true
	}

	labMap[guard.coordinate.y][guard.coordinate.x] = 'X'
	guard.coordinate.x = nextSpot.x
	guard.coordinate.y = nextSpot.y
	labMap[guard.coordinate.y][guard.coordinate.x] = guard.symbol

	return labMap, guard, true

}

func nextSpotOffScreen(labMap [][]rune, nextSpot Coordinate) bool {
	maxX := len(labMap[0]) - 1
	maxY := len(labMap) - 1

	return nextSpot.x < 0 ||
		nextSpot.x > maxX ||
		nextSpot.y < 0 ||
		nextSpot.y > maxY
}

func getIntendedMovement(guard Guard) Coordinate {
	nextSpot := Coordinate{
		x: guard.coordinate.x,
		y: guard.coordinate.y,
	}

	switch guard.symbol {
	case '^':
		nextSpot.y--
	case '>':
		nextSpot.x++
	case 'V':
		nextSpot.y++
	case '<':
		nextSpot.x--
	}

	return nextSpot
}

func countSpots(labMap [][]rune) int {
	uniqueSpots := 0

	for _, row := range labMap {
		for _, spot := range row {
			if spot == 'X' {
				uniqueSpots++
			}
		}
	}

	return uniqueSpots
}

func findGuard(mapRow []rune) int {

	if guardSpot := slices.Index(mapRow, '^'); guardSpot > -1 {
		return guardSpot
	}

	if guardSpot := slices.Index(mapRow, '>'); guardSpot > -1 {
		return guardSpot
	}

	if guardSpot := slices.Index(mapRow, 'V'); guardSpot > -1 {
		return guardSpot
	}

	if guardSpot := slices.Index(mapRow, '<'); guardSpot > -1 {
		return guardSpot
	}

	return -1
}

func turnGuard(guard rune) rune {
	switch guard {
	case '^':
		return '>'
	case '>':
		return 'V'
	case 'V':
		return '<'
	case '<':
		return '^'
	default:
		log.Fatal("INVALID GUARD")
		return '!'
	}
}
