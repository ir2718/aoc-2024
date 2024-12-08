package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFiles() ([][]string, map[string][][2]int) {
	file, err := os.Open("files/day_08/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)
	antennas := map[string][][2]int{}
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		rowSplit := strings.Split(row, "")
		for j, x := range rowSplit {
			if x != "." {
				antennas[x] = append(antennas[x], [2]int{i, j})
			}
		}
		grid = append(grid, rowSplit)
		i += 1
	}

	return grid, antennas
}

func firstProblem() {
	grid, antennas := readFiles()
	maxHeight, maxWidth := len(grid), len(grid[0])

	antinodes := map[[2]int]bool{}
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				diff := [2]int{v[i][0] - v[j][0], v[i][1] - v[j][1]}
				antinodeFirst := [2]int{v[i][0] + diff[0], v[i][1] + diff[1]}
				antinodeOther := [2]int{v[j][0] - diff[0], v[j][1] - diff[1]}
				if antinodeFirst[0] >= 0 && antinodeFirst[0] < maxHeight &&
					antinodeFirst[1] >= 0 && antinodeFirst[1] < maxWidth {
					antinodes[antinodeFirst] = true
					grid[antinodeFirst[0]][antinodeFirst[1]] = "#"
				}
				if antinodeOther[0] >= 0 && antinodeOther[0] < maxHeight &&
					antinodeOther[1] >= 0 && antinodeOther[1] < maxWidth {
					antinodes[antinodeOther] = true
					grid[antinodeOther[0]][antinodeOther[1]] = "#"
				}
			}
		}
	}

	numAntinodes := len(antinodes)
	fmt.Println(numAntinodes)
}

func secondProblem() {
	grid, antennas := readFiles()
	maxHeight, maxWidth := len(grid), len(grid[0])

	antinodes := map[[2]int]bool{}
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				diff := [2]int{v[i][0] - v[j][0], v[i][1] - v[j][1]}

				antinodeFirst := v[i]
				for {
					antinodeFirst = [2]int{
						antinodeFirst[0] + diff[0], antinodeFirst[1] + diff[1]}
					if antinodeFirst[0] >= 0 && antinodeFirst[0] < maxHeight &&
						antinodeFirst[1] >= 0 && antinodeFirst[1] < maxWidth {
						antinodes[antinodeFirst] = true
						grid[antinodeFirst[0]][antinodeFirst[1]] = "#"
					} else {
						antinodes[v[i]] = true
						break
					}
				}

				antinodeOther := v[j]
				for {
					antinodeOther = [2]int{
						antinodeOther[0] - diff[0], antinodeOther[1] - diff[1]}
					if antinodeOther[0] >= 0 && antinodeOther[0] < maxHeight &&
						antinodeOther[1] >= 0 && antinodeOther[1] < maxWidth {
						antinodes[antinodeOther] = true
						grid[antinodeOther[0]][antinodeOther[1]] = "#"
					} else {
						antinodes[v[j]] = true
						break
					}
				}
			}
		}
	}

	numAntinodes := len(antinodes)
	fmt.Println(numAntinodes)
}

func main() {
	// firstProblem()
	secondProblem()
}
