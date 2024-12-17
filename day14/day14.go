package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	x float64
	y float64

	vx float64
	vy float64
}

type Quadrant struct {
	minX float64
	minY float64

	maxX float64
	maxY float64
}

func main() {
	file, err := os.Open("../files/day14/day14.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	robots := []*Robot{}
	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`-?\d+`)

		numbers := regex.FindAllString(line, -1)
		if numbers != nil {
			x, _ := strconv.Atoi(numbers[0])
			y, _ := strconv.Atoi(numbers[1])
			vx, _ := strconv.Atoi(numbers[2])
			vy, _ := strconv.Atoi(numbers[3])
			robot := &Robot{x: float64(x), y: float64(y), vx: float64(vx), vy: float64(vy)}

			robots = append(robots, robot)
		}
	}

	mapWidth := float64(101)
	mapHeight := float64(103)
	seconds := 100
	for i := 1; i <= seconds; i++ {
		moveRobots(robots, mapWidth, mapHeight)
	}

	safetyScore := calculateSafetyScore(robots, mapWidth, mapHeight)

	fmt.Println(safetyScore)
}

func calculateSafetyScore(robots []*Robot, mapWidth, mapHeight float64) int {
	score := 1
	quadrants := [4]Quadrant{}
	robotsInQuadrants := make([]int, len(quadrants))

	midX := math.Floor(mapWidth / 2)
	midY := math.Floor(mapHeight / 2)

	q0 := Quadrant{minX: 0, maxX: midX - 1, minY: 0, maxY: midY - 1}
	q1 := Quadrant{minX: midX + 1, maxX: mapWidth - 1, minY: 0, maxY: midY - 1}
	q2 := Quadrant{minX: midX + 1, maxX: mapWidth - 1, minY: midY + 1, maxY: mapHeight - 1}
	q3 := Quadrant{minX: 0, maxX: midX - 1, minY: midY + 1, maxY: mapHeight - 1}

	quadrants[0] = q0
	quadrants[1] = q1
	quadrants[2] = q2
	quadrants[3] = q3

	for i, q := range quadrants {
		robotsInQuadrants[i] = countRobots(robots, q)
	}

	for _, r := range robotsInQuadrants {
		score *= r
	}

	return score
}

func countRobots(robots []*Robot, q Quadrant) int {
	count := 0
	for _, robot := range robots {
		if robot.x >= q.minX &&
			robot.x <= q.maxX &&
			robot.y >= q.minY &&
			robot.y <= q.maxY {
			count++
		}
	}

	return count
}

func moveRobots(robots []*Robot, mapWidth float64, mapHeight float64) {
	for _, robot := range robots {
		robot.x = robot.x + robot.vx
		robot.y = robot.y + robot.vy

		if robot.x > mapWidth-1 {
			robot.x -= mapWidth
		}

		if robot.x < 0 {
			robot.x += mapWidth
		}

		if robot.y > mapHeight-1 {
			robot.y -= mapHeight
		}

		if robot.y < 0 {
			robot.y += mapHeight
		}
	}
}
