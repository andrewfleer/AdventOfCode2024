package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val   int
	plus  *Node
	times *Node
}

func main() {
	file, err := os.Open("../files/day7/day7.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sumOfValids := 0

	for scanner.Scan() {
		line := scanner.Text()

		answerAndNumbers := strings.Split(line, ": ")

		answer, _ := strconv.Atoi(answerAndNumbers[0])

		numbers := strings.Split(answerAndNumbers[1], " ")

		if isValid(answer, numbers) {
			sumOfValids += answer
		}
	}

	fmt.Println(sumOfValids)
}

func isValid(answer int, numbers []string) bool {
	tree := buildTree(numbers)

	valid := isTreeValid(tree, answer)

	return valid
}

func isTreeValid(node *Node, answer int) bool {
	if node.plus == nil {
		return node.val == answer
	}

	if isTreeValid(node.plus, answer) {
		return true
	}

	return isTreeValid(node.times, answer)

}

/*func doMath(answer, startingVal int, node *Node) int {
	tempAnswer := answer
	if node.plus == nil {

	}

	tempAnswer = doMath(tempAnswer, startingVal, node.plus)

	if tempAnswer == startingVal {
		return tempAnswer
	}

	tempAnswer = answer
	return doMath(tempAnswer, startingVal, node.times)

}*/

func buildTree(numbers []string) *Node {
	var root *Node
	for _, val := range numbers {
		number, _ := strconv.Atoi(val)

		root = insert(root, number)

	}

	return root
}

func insert(node *Node, number int) *Node {
	if node == nil {
		return &Node{val: number}
	}

	node.plus = insert(node.plus, number)
	node.plus.val = node.val + number
	node.times = insert(node.times, number)
	node.times.val = node.val * number

	return node

}
