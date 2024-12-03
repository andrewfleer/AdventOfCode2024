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
	fullText := ""
	for scanner.Scan() {
		line := scanner.Text()

		fullText = fullText + line

	}

	value := readLine(fullText)

	totalValue += value

	fmt.Println(totalValue)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readLine(line string) int {
	val := 0
	r := regexp.MustCompile(`(mul\([0-9]*,[0-9]*\))|(do\(\))|(don't\(\))`)

	numberRegEx := regexp.MustCompile(`[0-9]+`)

	matches := r.FindAllString(line, -1)

	do := true
	for _, match := range matches {
		if match == "do()" {
			do = true
		} else if match == "don't()" {
			do = false
		} else {
			if do {
				nums := numberRegEx.FindAllString(match, -1)

				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])

				val += num1 * num2
			}
		}
	}

	//fmt.Println(matches)

	return val
}
