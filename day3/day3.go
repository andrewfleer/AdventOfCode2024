package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../files/day3/day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	totalValue := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		value := readLine(line)

		totalValue += value

	}

	fmt.Println(totalValue)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readLine(line string) int {
	val := 0
	r := regexp.MustCompile(`(mul\([0-9]*,[0-9]*\))`)

	numberRegEx := regexp.MustCompile(`[0-9]+`)

	matches := r.FindAllString(line, -1)

	for _, match := range matches {
		nums := numberRegEx.FindAllString(match, -1)

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		val += num1 * num2
	}

	//fmt.Println(matches)

	return val
}
