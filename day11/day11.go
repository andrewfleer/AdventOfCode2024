package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/day11/day11.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	stones := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, " ")

		for _, val := range vals {
			number, _ := strconv.Atoi(val)
			stones = append(stones, number)
		}
	}

	timesToBlink := 75
	//fmt.Println(stones)
	numStones := 0

	for _, stone := range stones {
		numStones += blinkCache(stone, timesToBlink)
	}

	fmt.Println(numStones)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type StoneBlinkResult struct {
	Value     int
	NumBlinks int
}

var blinkCacher = make(map[StoneBlinkResult]int)

func blinkCache(stone int, timesToBlink int) int {
	if c, ok := blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: timesToBlink}]; ok {
		return c
	}

	if timesToBlink == 0 {
		return 1
	}

	if stone == 0 {
		return blinkCache(1, timesToBlink-1)
	} else {
		digits := []rune(strconv.Itoa(stone))
		if len(digits)%2 == 0 {
			split := len(digits) / 2

			stone1, _ := strconv.Atoi(string(digits[:split]))
			stone2, _ := strconv.Atoi(string(digits[split:]))

			newStones := blinkCache(stone1, timesToBlink-1) + blinkCache(stone2, timesToBlink-1)
			blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: timesToBlink}] = newStones

			return newStones
		} else {
			newStones := blinkCache(stone*2024, timesToBlink-1)
			blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: timesToBlink}] = newStones

			return newStones
		}
	}
}
