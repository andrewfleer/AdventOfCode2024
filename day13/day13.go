package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Button struct {
	x    int
	y    int
	cost int
}

type Prize struct {
	x int
	y int
}

func main() {
	file, err := os.Open("../files/day13/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	aButton := Button{cost: 3}
	bButton := Button{cost: 1}

	prize := Prize{}

	totalCost := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		text := strings.Split(line, ":")

		indicator := text[0]
		xAndY := strings.Split(text[1], ", ")

		if indicator == "Prize" {
			xVal, _ := strconv.Atoi(strings.Split(xAndY[0], "=")[1])
			yVal, _ := strconv.Atoi(strings.Split(xAndY[1], "=")[1])

			prize.x = xVal
			prize.y = yVal

			cost := playGame(aButton, bButton, prize)

			if cost != -1 {
				totalCost += cost
			}
		} else {
			xVal, _ := strconv.Atoi(strings.Split(xAndY[0], "+")[1])
			yVal, _ := strconv.Atoi(strings.Split(xAndY[1], "+")[1])

			if indicator == "Button A" {
				aButton.x = xVal
				aButton.y = yVal
			}

			if indicator == "Button B" {
				bButton.x = xVal
				bButton.y = yVal
			}
		}
	}

	fmt.Println(totalCost)
}

func playGame(aButton, bButton Button, prize Prize) int {
	cost := -1

	bDividend := (prize.y*aButton.x - prize.x*aButton.y)
	bDivisor := (bButton.y*aButton.x - aButton.y*bButton.x)

	if bDivisor == 0 {
		return cost
	} else if bDividend%bDivisor != 0 {
		return cost
	}

	timesToPressB := bDividend / bDivisor

	aDividend := prize.x - timesToPressB*bButton.x
	aDivisor := aButton.x

	if aDivisor == 0 {
		return cost
	} else if aDividend%aDivisor != 0 {
		return cost
	}

	timesToPressA := aDividend / aDivisor

	if timesToPressB > 100 ||
		timesToPressA > 100 {
		return -1
	}

	cost = timesToPressA*aButton.cost + timesToPressB*bButton.cost

	return cost
}
