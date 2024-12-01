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

func main() {
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
