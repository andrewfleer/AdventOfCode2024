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
	file, err := os.Open("../files/day1/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	totalScore := 0
	var leftArray []int
	var rightMap = make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftArray, rightMap = parseLine(scanner.Text(), leftArray, rightMap)

	}

	for _, leftVal := range leftArray {
		score := calculateSimilarityScore(leftVal, rightMap)
		totalScore += score
	}

	fmt.Println(totalScore)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateSimilarityScore(leftVal int, rightMap map[int]int) int {
	frequency := rightMap[leftVal]

	return leftVal * frequency
}

func parseLine(text string, leftArray []int, rightMap map[int]int) ([]int, map[int]int) {
	numbers := strings.Fields(text)

	leftInt, _ := strconv.Atoi(numbers[0])
	rightInt, _ := strconv.Atoi(numbers[1])

	leftArray = append(leftArray, leftInt)

	rightMap[rightInt] = rightMap[rightInt] + 1

	return leftArray, rightMap
}
