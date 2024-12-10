package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFiles() [][]int {
	file, err := os.Open("files/day_10/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]int, 0)
	for scanner.Scan() {
		row := scanner.Text()
		rowSplit := strings.Split(row, "")
		rowInt := make([]int, 0)
		for _, x := range rowSplit {
			intVal, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}
			rowInt = append(rowInt, intVal)
		}
		grid = append(grid, rowInt)
	}

	return grid
}

func foundInEnd(pts [][4]int, endPt [2]int, startPt [2]int) bool {
	for _, x := range pts {
		if startPt[0] == x[0] && startPt[1] == x[1] &&
			endPt[0] == x[2] && endPt[1] == x[3] {
			return true
		}
	}
	return false
}

func (p *Point) findStart() [2]int {
	for p.prev != nil {
		p = p.prev
	}
	return p.coords
}

type Point struct {
	coords [2]int
	value  int
	prev   *Point
}

func firstProblem(distinct bool) {
	grid := readFiles()

	// Find all zeros as they are possible starting positions.
	open := make([]Point, 0)
	for i, row := range grid {
		for j, elem := range row {
			if elem == 0 {
				open = append(open, Point{[2]int{i, j}, 0, nil})
			}
		}
	}

	// Do BFS starting from zero.
	directions := [][2]int{
		{-1, 0}, // N
		// {-1, 1}, // NE
		{0, 1}, // E
		// {1, 1},  // SE
		{1, 0}, // S
		// {1, -1}, // SW
		{0, -1}, // W
		// {1, -1}, // NW
	}
	numFoundTrails := 0
	// (istart, jstart, iend, jend)
	endPts := make([][4]int, 0)
	for len(open) != 0 {
		curr := open[0]
		open = open[1:]

		for _, d := range directions {
			iNew := curr.coords[0] + d[0]
			jNew := curr.coords[1] + d[1]

			// Skip if it is out of bounds.
			if iNew < 0 || iNew >= len(grid) || jNew < 0 || jNew >= len(grid[0]) {
				continue
			}

			// Add if the value is equal to current plus one.
			oldValue := curr.value
			newValue := grid[iNew][jNew]

			if newValue == (oldValue + 1) {
				possibleEnd := Point{[2]int{iNew, jNew}, newValue, &curr}
				start := possibleEnd.findStart()
				end := possibleEnd.coords
				if newValue == 9 && ((distinct && !foundInEnd(endPts, end, start)) || !distinct) {
					numFoundTrails++
					endPts = append(endPts, [4]int{start[0], start[1], end[0], end[1]})
				} else {
					open = append(
						[]Point{
							{
								[2]int{iNew, jNew},
								newValue,
								&curr,
							},
						},
						open...,
					)
				}
			}
		}
	}
	fmt.Println(numFoundTrails)

}

func secondProblem() {
	firstProblem(false)
}

func main() {
	firstProblem(true)
	secondProblem()
}
