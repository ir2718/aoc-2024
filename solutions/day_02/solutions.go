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
	file, err := os.Open("files/day_02/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slices := make([][]int, 0)
	for scanner.Scan() {
		splitLine := strings.Fields(scanner.Text())
		splitLineInt := make([]int, len(splitLine))
		for i, s := range splitLine {
			intNum, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			splitLineInt[i] = intNum
		}
		slices = append(slices, splitLineInt)
	}
	return slices
}

func firstProblem() {
	slices := readFiles()

	sum := 0
	for _, s := range slices {
		increasing := s[1] > s[0]
		add := true
		for i := 0; i < len(s)-1; i++ {
			if (increasing && (!(s[i+1] >= s[i]+1) || !(s[i+1] <= s[i]+3))) ||
				(!increasing && (!(s[i+1] <= s[i]-1) || !(s[i+1] >= s[i]-3))) {
				add = false
			}
		}
		if add {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func secondProblem() {
	slices := readFiles()

	sum := 0
	for _, s := range slices {
		// Try with the original solution.
		increasing := s[1] > s[0]
		add := true
		for i := 0; i < len(s)-1; i++ {
			if (increasing && (!(s[i+1] >= s[i]+1) || !(s[i+1] <= s[i]+3))) ||
				(!increasing && (!(s[i+1] <= s[i]-1) || !(s[i+1] >= s[i]-3))) {
				add = false
			}
		}
		if add {
			sum += 1
			continue
		}

		// Try by removing a single value and then doing the original solution.
		for j := 0; j < len(s); j++ {

			// Create a copy of slice.
			s_copy := make([]int, len(s))
			copy(s_copy, s)
			s2 := append(s_copy[:j], s_copy[j+1:]...)

			increasing := s2[1] > s2[0]
			add := true
			for i := 0; i < len(s2)-1; i++ {
				if (increasing && (!(s2[i+1] >= s2[i]+1) || !(s2[i+1] <= s2[i]+3))) ||
					(!increasing && (!(s2[i+1] <= s2[i]-1) || !(s2[i+1] >= s2[i]-3))) {
					add = false
				}
			}

			if add {
				sum += 1
				break
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	firstProblem()
	secondProblem()
}
