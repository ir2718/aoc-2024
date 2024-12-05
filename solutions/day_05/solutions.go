package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFiles() (map[int][]int, [][]int) {
	file, err := os.Open("files/day_05/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := map[int][]int{}
	orderings := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Scan the maps.
		if strings.Contains(text, "|") {
			split := strings.Split(text, "|")
			k, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}
			v, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			_, ok := m[k]
			if !ok {
				values := make([]int, 1)
				values[0] = v
				m[k] = values
			} else {
				m[k] = append(m[k], v)
			}

			// Scan the rows.
		} else if strings.TrimSpace(text) != "" {
			split := strings.Split(strings.TrimSpace(text), ",")
			splitInt := make([]int, len(split))
			for i := range split {
				iInt, err := strconv.Atoi(split[i])
				if err != nil {
					log.Fatal(err)
				}
				splitInt[i] = iInt
			}
			orderings = append(orderings, splitInt)
		}
	}

	return m, orderings
}

func isInSlice(x int, slice []int) bool {
	for _, y := range slice {
		if x == y {
			return true
		}
	}
	return false
}

func firstProblem() [][]int {
	m, orderings := readFiles()

	sum := 0
	incorrectlyOrdered := make([][]int, 0)
	for _, o := range orderings {

		// Go over each index to the right and check if the left values are ok.
		correctOrder := true
		for i := 0; i < len(o); i++ {
			currLeft := o[i]
			for j := i + 1; j < len(o); j++ {
				currRight := o[j]
				currRightValues := m[currRight]
				if isInSlice(currLeft, currRightValues) {
					correctOrder = false
					break
				}
			}
			if !correctOrder {
				break
			}
		}
		if correctOrder {
			sum += o[len(o)/2]
		} else {
			incorrectlyOrdered = append(incorrectlyOrdered, o)
		}
	}
	fmt.Println(sum)
	return incorrectlyOrdered
}

func secondProblem(incorrectlyOrdered [][]int) {
	m, _ := readFiles()

	sum := 0
	for _, o := range incorrectlyOrdered {

		for {
			// Go over each index to the right and check if the left values are
			// ok. If they are not, just swap them to achieve the correct order.
			correctOrder := true
			iSwap, jSwap := -1, -1
			for i := 0; i < len(o); i++ {
				currLeft := o[i]
				for j := i + 1; j < len(o); j++ {
					currRight := o[j]
					currRightValues := m[currRight]
					if isInSlice(currLeft, currRightValues) {
						correctOrder = false
						iSwap, jSwap = i, j
						break
					}
				}
				if !correctOrder {
					break
				}
			}
			if correctOrder {
				sum += o[len(o)/2]
				break
			} else {
				// Swap i and j elements.
				tmp := o[iSwap]
				o[iSwap] = o[jSwap]
				o[jSwap] = tmp
			}
		}

	}
	fmt.Println(sum)
}

func main() {
	incorrectlyOrdered := firstProblem()
	secondProblem(incorrectlyOrdered)
}
