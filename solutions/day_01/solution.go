package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFiles() ([]int, []int) {
	file, err := os.Open("files/day_01/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the files and save them into slices.
	first := make([]int, 0)
	second := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Fields(scanner.Text())
		firstStr, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		secondStr, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		first = append(first, firstStr)
		second = append(second, secondStr)
	}
	return first, second
}

func firstProblem() {
	first, second := readFiles()

	// Sort the arrays. Calculate the differences and sum them.
	sort.Ints(first)
	sort.Ints(second)
	sum := 0.0
	for i := range first {
		firstNum := first[i]
		secondNum := second[i]
		sum += math.Abs(float64(firstNum) - float64(secondNum))
	}
	fmt.Printf("%d\n", int32(sum))
}

func secondProblem() {
	first, second := readFiles()

	// Sort the arrays.
	sort.Ints(first)
	sort.Ints(second)

	sum := 0
	i, j := 0, 0
	for {
		// Count how many occurrences are in the left list.
		currFirst := first[i]
		countFirst := 0
		for i < len(first) && first[i] == currFirst {
			countFirst++
			i++
		}

		// Move the index to the first same value.
		for j < len(first) && second[j] < currFirst {
			j++
		}

		if j == len(first) {
			break
		} else if second[j] > currFirst {
			// In case the element is bigger, keep moving the first index until you
			// get to an equal or bigger value.
			for i < len(first) && first[i] < currFirst {
				i++
			}
			if i == len(first) {
				break
			}
			continue
		}

		// Count how many occurences are there in the right list.
		countSecond := 0
		for j < len(first) && second[j] == currFirst {
			countSecond++
			j++
		}

		sum += countFirst * (currFirst * countSecond)

		// Check if the indexes are still in range.
		if i >= len(first) || j >= len(first) {
			break
		}
	}
	fmt.Printf("%d\n", int64(sum))
}

func main() {
	firstProblem()
	secondProblem()
}
