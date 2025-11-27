package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

var directions = []Point{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

type Grid struct {
	rowsLen int
	colsLen int
	data    [][]int
}

func (g *Grid) isPositionValid(x, y int) bool {
	return x >= 0 && x < g.rowsLen && y >= 0 && y < g.colsLen
}

func (g *Grid) bfs(start Point) int {
	visited := make(map[Point]bool)
	queue := []Point{start}
	reachableNines := make(map[Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

		if g.data[current.X][current.Y] == 9 {
			reachableNines[current] = true
		}

		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}

			if g.isPositionValid(next.X, next.Y) && !visited[next] && g.data[next.X][next.Y] == g.data[current.X][current.Y]+1 {
				queue = append(queue, next)
			}
		}
	}

	return len(reachableNines)
}

func NewGrid(data [][]int) *Grid {
	return &Grid{
		rowsLen: len(data),
		colsLen: len(data[0]),
		data:    data,
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error with file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var gridData [][]int

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, ch := range line {
			row[i] = int(ch - '0')
		}

		gridData = append(gridData, row)
	}

	grid := NewGrid(gridData)

	totalScore := 0
	for x := 0; x < grid.rowsLen; x++ {
		for y := 0; y < grid.colsLen; y++ {
			if grid.data[x][y] == 0 {
				totalScore += grid.bfs(Point{x, y})
			}
		}
	}
	fmt.Printf("Total: %d\n", totalScore)
}

func convertLineToArray(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		if field == "" {
			continue
		}
		num, err := strconv.Atoi(field)
		if err != nil {
			log.Fatalf("Invalid number: %q: %v", field, err)
		}
		nums[i] = num
	}

	return nums
}
