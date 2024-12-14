package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

		numbers := []int64{}
		for _, r := range line {
			number, _ := strconv.Atoi(string(r))
			numbers = append(numbers, int64(number))
		}

		checksum := processDisk(numbers)

		fmt.Println(checksum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processDisk(numbers []int64) int64 {
	// Build String
	diskString := buildDiskString(numbers)

	// for _, val := range diskString {
	// 	if val == -1 {
	// 		fmt.Print(".")
	// 	} else {

	// 		fmt.Print(val)
	// 	}
	// }

	// fmt.Println()

	//foo := len(diskString)

	// Compress File
	compressedFile := compress(diskString)

	// for _, val := range compressedFile {
	// 	if val == -1 {
	// 		fmt.Print(".")
	// 	} else {

	// 		fmt.Print(val)
	// 	}
	// }

	fmt.Println()

	//bar := len(compressedFile)

	//fmt.Printf("%d | %d\n", foo, bar)

	//get checksum
	checkSum := calculateChecksum(compressedFile)

	return checkSum
}

func calculateChecksum(compressedFile []int64) int64 {
	checksum := int64(0)

	for i, value := range compressedFile {
		if value == -1 {
			continue
		}

		checksum += int64(i) * value
	}

	return checksum
}

func compress(diskString []int64) []int64 {
	newString := make([]int64, len(diskString))

	copy(newString, diskString)

	lastFileMoved := int64(math.MaxInt64)
	bitVal := int64(-1)
	fileSize := int64(0)
	for tailPointer := int64(len(newString) - 1); tailPointer >= 0; tailPointer-- {
		tailVal := newString[tailPointer]

		if tailVal != bitVal {
			if bitVal != -1 && bitVal < lastFileMoved {
				// Try to move the file.
				indexToMove := findEmptySpace(newString, fileSize, tailPointer)

				if indexToMove != -1 {
					for i := int64(0); i < fileSize; i++ {
						newString[indexToMove+i] = bitVal
						newString[tailPointer+i+1] = -1
					}

					// for _, val := range newString {
					// 	if val == -1 {
					// 		fmt.Print(".")
					// 	} else {

					// 		fmt.Print(val)
					// 	}
					// }

					// fmt.Println()
					lastFileMoved = bitVal

					//fmt.Printf("last file: %d\n", lastFileMoved)

					bitVal = -1
				}
			}
			bitVal = tailVal
			fileSize = 1
		} else {
			fileSize++
		}
	}

	/*newStringTailPointer := len(newString) - 1

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

	}*/

	return newString
}

func findEmptySpace(diskString []int64, size, currentIndex int64) int64 {
	for i := int64(0); i <= currentIndex; i++ {
		if diskString[i] == -1 {
			for j := i; j <= currentIndex; j++ {
				if diskString[j] != -1 {
					break
				}

				if j-i == size-1 {
					return i
				}
			}
		}
	}

	return -1
}

func buildDiskString(numbers []int64) []int64 {
	diskString := []int64{}
	fileID := int64(0)
	for i := range numbers {
		char := int64(-1)
		if i%2 == 0 {
			char = fileID
			fileID++
		}

		for j := int64(0); j <= numbers[i]-1; j++ {
			diskString = append(diskString, char)
		}
	}

	return diskString
}
