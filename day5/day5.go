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

	rulesMapMustGoAfter, rulesMapMustGoBefore := mapTheRules(updateRules)

	invalidUpdates := checkUpdates(rulesMapMustGoAfter, updates)

	fixedUpdates := fixUpdates(rulesMapMustGoAfter, rulesMapMustGoBefore, invalidUpdates)

	middlePagesSum := processUpdates(fixedUpdates)

	fmt.Println(middlePagesSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func mapTheRules(updateRules []string) (map[string][]string, map[string][]string) {
	rulesMapAfter := make(map[string][]string)
	rulesMapBefore := make(map[string][]string)
	for _, rule := range updateRules {
		pages := strings.Split(rule, "|")
		page := pages[0]
		laterPage := pages[1]

		if rulesMapAfter[laterPage] == nil {
			rulesMapAfter[laterPage] = []string{}
		}

		rulesMapAfter[laterPage] = append(rulesMapAfter[laterPage], page)

		if rulesMapBefore[page] == nil {
			rulesMapBefore[page] = []string{}
		}

		rulesMapBefore[page] = append(rulesMapBefore[page], laterPage)
	}

	return rulesMapAfter, rulesMapBefore
}

func checkUpdates(rulesMap map[string][]string, updates []string) [][]string {
	invalidUpdates := [][]string{}

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
		if !valid {
			invalidUpdates = append(invalidUpdates, pages)
		}
	}

	return invalidUpdates
}

func fixUpdates(rulesMapAfter, rulesMapBefore map[string][]string, invalidUpdates [][]string) [][]string {
	fixedUpdates := [][]string{}

	for _, update := range invalidUpdates {
		fixedUpdate := []string{}

		for i := len(update) - 1; i >= 0; i-- {
			page := update[i]
			if i == len(update)-1 {
				fixedUpdate = append(fixedUpdate, page)
				continue
			}

			minIndex := len(fixedUpdate) - 1
			mustGoBeforePages := rulesMapBefore[page]
			for i, fixedPage := range fixedUpdate {
				if slices.Contains(mustGoBeforePages, fixedPage) {
					if i < minIndex {
						minIndex = i
					}
				}
			}

			mustGoAfterPages := rulesMapAfter[page]
			maxIndex := 0
			for i, updatePage := range fixedUpdate {
				if slices.Contains(mustGoAfterPages, updatePage) {
					if i+1 > maxIndex {
						maxIndex = i + 1
					}
				}
			}

			indexToUse := maxIndex
			if minIndex > indexToUse {
				indexToUse = minIndex
			}

			fixedUpdate = append(fixedUpdate[:indexToUse], append([]string{page}, fixedUpdate[indexToUse:]...)...)
		}

		fixedUpdates = append(fixedUpdates, fixedUpdate)
	}

	return fixedUpdates
}

func processUpdates(validUpdates [][]string) int {
	middlePageSum := 0

	for _, update := range validUpdates {

		length := len(update)

		if length == 0 {
			continue
		}

		middleIndex := length / 2

		middlePage, _ := strconv.Atoi(update[middleIndex])

		middlePageSum += middlePage
	}

	return middlePageSum
}
