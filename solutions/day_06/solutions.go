package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFiles() ([][]string, [2]int) {
	file, err := os.Open("files/day_06/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)
	i := 0
	var currPoint [2]int
	for scanner.Scan() {
		row := scanner.Text()
		rowSplit := strings.Split(row, "")
		for j, x := range rowSplit {
			if x == "^" {
				currPoint = [2]int{i, j}
			}
		}
		grid = append(grid, rowSplit)
		i += 1
	}
	return grid, currPoint
}

func firstProblem() {
	grid, currPoint := readFiles()

	gridHeight, gridWidth := len(grid), len(grid[0])
	visited := map[[2]int]bool{}
	visited[currPoint] = true
	currDirection := [2]int{-1, 0}
	for {
		possibleNextPoint := [2]int{
			currPoint[0] + currDirection[0],
			currPoint[1] + currDirection[1],
		}

		if possibleNextPoint[0] < 0 || possibleNextPoint[0] >= gridHeight ||
			possibleNextPoint[1] < 0 || possibleNextPoint[1] >= gridWidth {
			break
		} else if grid[possibleNextPoint[0]][possibleNextPoint[1]] == "#" {
			if currDirection[0] == -1 && currDirection[1] == 0 {
				currDirection[0] = 0
				currDirection[1] = 1
			} else if currDirection[0] == 0 && currDirection[1] == 1 {
				currDirection[0] = 1
				currDirection[1] = 0
			} else if currDirection[0] == 1 && currDirection[1] == 0 {
				currDirection[0] = 0
				currDirection[1] = -1
			} else {
				currDirection[0] = -1
				currDirection[1] = 0
			}
		} else {
			currPoint[0] += currDirection[0]
			currPoint[1] += currDirection[1]
			visited[currPoint] = true
		}
	}
	numVisited := len(visited)
	fmt.Println(numVisited)
}

func secondProblem() {
	originalGrid, originalCurrPoint := readFiles()

	gridHeight, gridWidth := len(originalGrid), len(originalGrid[0])

	numPossibleCycles := 0
	for i := 0; i < gridHeight; i++ {
		for j := 0; j < gridWidth; j++ {
			if originalGrid[i][j] == "#" {
				continue
			}

			currPoint := [2]int{originalCurrPoint[0], originalCurrPoint[1]}
			visited := map[[2]int][2]int{}
			currDirection := [2]int{-1, 0}
			visited[currPoint] = currDirection

			originalGrid[i][j] = "#"

			for {
				// Simulate moving around the grid with lazy implementation of
				// rotation.
				possibleNextPoint := [2]int{
					currPoint[0] + currDirection[0],
					currPoint[1] + currDirection[1],
				}

				originalDirection, ok := visited[possibleNextPoint]
				if ok && originalDirection[0] == currDirection[0] && originalDirection[1] == currDirection[1] {
					numPossibleCycles++
					break
				} else if possibleNextPoint[0] < 0 || possibleNextPoint[0] >= gridHeight ||
					possibleNextPoint[1] < 0 || possibleNextPoint[1] >= gridWidth {
					break
				} else if originalGrid[possibleNextPoint[0]][possibleNextPoint[1]] == "#" {
					if currDirection[0] == -1 && currDirection[1] == 0 {
						currDirection[0] = 0
						currDirection[1] = 1
					} else if currDirection[0] == 0 && currDirection[1] == 1 {
						currDirection[0] = 1
						currDirection[1] = 0
					} else if currDirection[0] == 1 && currDirection[1] == 0 {
						currDirection[0] = 0
						currDirection[1] = -1
					} else {
						currDirection[0] = -1
						currDirection[1] = 0
					}
				} else {
					currPoint[0] += currDirection[0]
					currPoint[1] += currDirection[1]
					visited[currPoint] = currDirection
					originalGrid[currPoint[0]][currPoint[1]] = "*"
				}
			}

			originalGrid[i][j] = "."
		}
	}

	fmt.Println(numPossibleCycles)
}

func main() {
	firstProblem()
	secondProblem()
}
