package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/day2/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	totalSafe := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		isSafe := parseLine(scanner.Text())

		if isSafe {
			totalSafe++
		}

	}

	fmt.Println(totalSafe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine(s string) bool {
	increasing := true
	numbers := strings.Fields(s)

	firstVal, _ := strconv.Atoi(numbers[0])
	secondVal, _ := strconv.Atoi(numbers[1])

	if firstVal == secondVal {
		return false
	}

	if firstVal > secondVal {
		increasing = false
	}

	for i, _ := range numbers {
		if i == len(numbers)-1 {
			break
		}

		val, _ := strconv.Atoi(numbers[i])
		compVal, _ := strconv.Atoi(numbers[i+1])

		change := math.Abs(float64(val) - float64(compVal))

		if change < 1 || change > 3 {
			return false
		}

		if increasing && val > compVal {
			return false
		}

		if !increasing && val < compVal {
			return false
		}

	}

	return true
}
