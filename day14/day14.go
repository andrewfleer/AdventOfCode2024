package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
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
	seconds := 101 * 103

	const rows, cols = 10, 10
	rowNum := 0
	colNum := 0
	plots := make([][]*plot.Plot, rows)

	for i := 1; i <= seconds; i++ {
		if colNum == 0 {
			plots[rowNum] = make([]*plot.Plot, cols)
		}

		moveRobots(robots, mapWidth, mapHeight)

		p, err := createScatterPlot(robots, i, mapWidth, mapHeight)
		if err != nil {
			panic(err)
		}

		plots[rowNum][colNum] = p

		colNum++
		if colNum == cols {
			rowNum++
			colNum = 0
		}

		if rowNum == rows {
			// save the image
			img := vgimg.New(vg.Points(2000), vg.Points(2000))
			dc := draw.New(img)

			t := draw.Tiles{
				Rows: rows,
				Cols: cols,
				PadX: vg.Millimeter,
				PadY: vg.Millimeter,
			}
			canvases := plot.Align(plots, t, dc)

			for y := 0; y < rows; y++ {
				for x := 0; x < cols; x++ {
					fmt.Printf("Printing %d row %d column\n", y, x)
					if plots[y][x] != nil {
						plots[y][x].Draw(canvases[y][x])
						w, err := os.Create("robits.png")
						if err != nil {
							panic(err)
						}

						png := vgimg.PngCanvas{Canvas: img}
						if _, err := png.WriteTo(w); err != nil {
							panic(err)
						}
					}
				}
			}
			colNum = 0
			rowNum = 0
		}

	}

	safetyScore := calculateSafetyScore(robots, mapWidth, mapHeight)

	fmt.Println(safetyScore)
}

func createScatterPlot(robots []*Robot, seconds int, mapWidth, mapHeight float64) (*plot.Plot, error) {
	p := plot.New()

	// Disable all axis annotations
	p.X.Label.Text = ""
	p.Y.Label.Text = ""
	p.X.Padding = 0
	p.Y.Padding = 0

	// Set ranges before disabling ticks
	p.X.Min, p.X.Max = 0, 10
	p.Y.Min, p.Y.Max = 0, 10

	// Disable ticks and grid
	p.X.Tick.Width = 0
	p.Y.Tick.Width = 0
	p.X.Width = 0
	p.Y.Width = 0

	p.Title.Text = strconv.Itoa(seconds) + " seconds"

	points := make(plotter.XYs, len(robots))
	for r, robot := range robots {
		points[r].X = robot.x
		points[r].Y = robot.y
	}

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		return nil, err
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	scatter.GlyphStyle.Radius = vg.Points(2)

	p.Add(scatter)
	p.X.Min, p.X.Max = 0, mapWidth
	p.Y.Min, p.Y.Max = 0, mapHeight

	return p, nil

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
