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

	timesToBlink := 25
	//fmt.Println(stones)
	for i := 0; i < timesToBlink; i++ {
		stones = blink(stones)

		//fmt.Println(stones)
	}

	fmt.Println(len(stones))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		stone := stones[i]
		if stone == 0 {
			stone = 1
			stones[i] = stone
		} else {
			digits := []rune(strconv.Itoa(stone))
			if len(digits)%2 == 0 {
				split := len(digits) / 2

				stone1, _ := strconv.Atoi(string(digits[:split]))
				stone2, _ := strconv.Atoi(string(digits[split:]))

				//insert first stone
				if i == 0 {
					stones = append([]int{stone1}, stones...)
				} else {
					stones = append(stones[:i], append([]int{stone1}, stones[i:]...)...)
				}

				// Replace stone with the second stone
				stones[i+1] = stone2
				i++
			} else {
				stone *= 2024
				stones[i] = stone
			}
		}
	}

	return stones
}
