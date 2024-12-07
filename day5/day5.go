package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../files/day5/day5.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	updateRules := []string{}
	updates := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			updateRules = append(updateRules, line)
		} else {
			updates = append(updates, line)
		}

	}

	rulesMap := mapTheRules(updateRules)

	validUpdates := checkUpdates(rulesMap, updates)

	middlePagesSum := processUpdates(validUpdates)

	fmt.Println(middlePagesSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func mapTheRules(updateRules []string) map[string][]string {
	rulesMap := make(map[string][]string)
	for _, rule := range updateRules {
		pages := strings.Split(rule, "|")
		page := pages[0]
		laterPage := pages[1]

		if rulesMap[laterPage] == nil {
			rulesMap[laterPage] = []string{}
		}

		rulesMap[laterPage] = append(rulesMap[laterPage], page)
	}

	return rulesMap
}

func checkUpdates(rulesMap map[string][]string, updates []string) []string {
	validUpdates := []string{}

	for _, update := range updates {
		valid := true

		pages := strings.Split(update, ",")
		for i, page := range pages {
			beforePages := rulesMap[page]
			for j, otherPage := range pages {
				if j <= i {
					continue
				}

				if slices.Contains(beforePages, otherPage) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}

func processUpdates(validUpdates []string) int {
	middlePageSum := 0

	for _, update := range validUpdates {
		pages := strings.Split(update, ",")
		length := len(pages)

		if length == 0 {
			continue
		}

		middleIndex := length / 2

		middlePage, _ := strconv.Atoi(pages[middleIndex])

		middlePageSum += middlePage
	}

	return middlePageSum
}
