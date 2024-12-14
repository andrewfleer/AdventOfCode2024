package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../files/day9/day9.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := []int{}
		for _, r := range line {
			number, _ := strconv.Atoi(string(r))
			numbers = append(numbers, number)
		}

		checksum := processDisk(numbers)

		fmt.Println(checksum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processDisk(numbers []int) int {
	// Build String
	diskString := buildDiskString(numbers)

	//fmt.Println(diskString)

	//foo := len(diskString)

	// Compress File
	compressedFile := compress(diskString)

	// for _, b := range compressedFile {
	// 	fmt.Printf("%c", b)
	// }
	// fmt.Println()

	//bar := len(compressedFile)

	//fmt.Printf("%d | %d\n", foo, bar)

	//get checksum
	checkSum := calculateChecksum(compressedFile)

	return checkSum
}

func calculateChecksum(compressedFile []int) int {
	checksum := 0

	for i, value := range compressedFile {
		if value == -1 {
			break
		}

		checksum += i * value
	}

	return checksum
}

func compress(diskString []int) []int {
	newString := make([]int, len(diskString))

	copy(newString, diskString)

	newStringTailPointer := len(newString) - 1

	for i, r := range newString {
		if r == -1 {
			for newStringTailPointer > 0 {
				if newString[newStringTailPointer] != -1 {
					newString[i] = newString[newStringTailPointer]
					newString[newStringTailPointer] = -1

					// for _, b := range newString {
					// 	fmt.Printf("%c", b)
					// }
					// fmt.Println()
					break
				}
				newStringTailPointer--
			}
		}

		if isFileCompressed(newString) {
			break
		}

	}

	return newString
}

func isFileCompressed(fileString []int) bool {
	headPointer := 0
	tailPointer := len(fileString) - 1

	for headPointer < tailPointer {
		if fileString[headPointer] == -1 {
			break
		}
		headPointer++
	}

	for tailPointer > 0 {
		if fileString[tailPointer] != -1 {
			break
		}
		tailPointer--
	}

	return headPointer >= tailPointer
}

func buildDiskString(numbers []int) []int {
	diskString := []int{}
	fileID := 0
	for i := range numbers {
		char := -1
		if i%2 == 0 {
			char = fileID
			fileID++
		}

		for j := 0; j <= numbers[i]-1; j++ {
			diskString = append(diskString, char)
		}
	}

	return diskString
}
