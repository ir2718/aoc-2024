package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFiles() string {
	file, err := os.ReadFile("files/day_03/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func firstProblem() {
	str := readFiles()

	// Define regex for solving the problem.
	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		log.Fatal(err)
	}

	// Find all matches.
	matches := r.FindAllString(str, -1)
	sum := 0
	for _, m := range matches {
		splitLine := strings.Split(m, ",")
		splitLineFirst, err := strconv.Atoi(splitLine[0][4:])
		if err != nil {
			log.Fatal(err)
		}
		splitLineSecond, err := strconv.Atoi(splitLine[1][:len(splitLine[1])-1])
		if err != nil {
			log.Fatal(err)
		}
		sum += splitLineFirst * splitLineSecond
	}
	fmt.Println(sum)
}

func secondProblem() {
	str := readFiles()

	// Define regex for solving the problem.
	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	// Find all matches.
	matches := r.FindAllString(str, -1)

	mulEnabled := true
	sum := 0
	for _, m := range matches {
		if m == "do()" {
			mulEnabled = true
		} else if m == "don't()" {
			mulEnabled = false
		} else if mulEnabled {
			splitLine := strings.Split(m, ",")
			splitLineFirst, err := strconv.Atoi(splitLine[0][4:])
			if err != nil {
				log.Fatal(err)
			}
			splitLineSecond, err := strconv.Atoi(splitLine[1][:len(splitLine[1])-1])
			if err != nil {
				log.Fatal(err)
			}
			sum += splitLineFirst * splitLineSecond
		}
	}
	fmt.Println(sum)
}

func main() {
	// firstProblem()
	secondProblem()
}
