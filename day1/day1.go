package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	totalDistance := 0
	var leftArray []int
	var rightArray []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftArray, rightArray = parseLine(scanner.Text(), leftArray, rightArray)

	}

	for i, leftVal := range leftArray {
		rightVal := rightArray[i]

		distance := math.Abs(float64(leftVal) - float64(rightVal))
		totalDistance += int(distance)
	}

	fmt.Println(totalDistance)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine(text string, leftArray, rightArray []int) ([]int, []int) {
	numbers := strings.Fields(text)

	leftInt, _ := strconv.Atoi(numbers[0])
	rightInt, _ := strconv.Atoi(numbers[1])

	leftArray = append(leftArray, leftInt)
	rightArray = append(rightArray, rightInt)

	sort.Ints(leftArray)
	sort.Ints(rightArray)
	return leftArray, rightArray
}
