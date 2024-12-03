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
	file, err := os.Open("../files/day2/day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	totalSafe := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		unsafe := true
		index := -2

		for unsafe {
			numbers := strings.Fields(line)
			index++

			if index == len(numbers) {
				break
			}

			isSafe := removeElementAndCheckResults(numbers, index)

			if isSafe {
				unsafe = false
				break
			}

		}
		//isSafe := parseLine(numbers, true)

		//fmt.Println(line + " || " + strconv.FormatBool(!unsafe))

		if !unsafe {
			totalSafe++
		}

	}

	fmt.Println(totalSafe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func removeElementAndCheckResults(tempNumbers []string, index int) bool {
	if index > -1 {
		tempNumbers = append(tempNumbers[:index], tempNumbers[index+1:]...)
	}

	//	fmt.Println(index)
	//	fmt.Println(tempNumbers)
	return parseLine(tempNumbers)

}

func parseLine(numbers []string) bool {
	increasing := true

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
