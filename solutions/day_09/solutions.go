package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFiles() []int {
	s, err := os.ReadFile("files/day_09/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	sStr := strings.Split(strings.TrimSpace(string(s)), "")
	sInt := make([]int, len(sStr))
	for i, x := range sStr {
		currInt, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		sInt[i] = currInt
	}
	return sInt
}

func expand(nums []int) []string {
	newStr := make([]string, 0)
	currValue := 0
	for i := 0; i < len(nums); i++ {
		useValue := "."
		if i%2 == 0 {
			useValue = strconv.Itoa(currValue)
		}

		count := 0
		for count < nums[i] {
			newStr = append(newStr, useValue)
			count++
		}

		if i%2 == 0 {
			currValue++
		}
	}
	return newStr
}

func findLastNonDot(curr int, s []string) int {
	lastIdx := len(s)
	for i := curr; i >= 0; i-- {
		if s[i] != "." {
			lastIdx = i
			break
		}
	}
	return lastIdx
}

func firstFirstDot(curr int, s []string) int {
	firstIdx := -1
	for i := curr; i < len(s); i++ {
		if s[i] == "." {
			firstIdx = i
			break
		}
	}
	return firstIdx
}

func moveToCorrect(s []string) []string {
	last := findLastNonDot(len(s)-1, s)
	first := firstFirstDot(0, s)
	for last > first && last != len(s) && first != -1 {
		tmp := s[last]
		s[last] = s[first]
		s[first] = tmp

		last = findLastNonDot(last-1, s)
		first = firstFirstDot(first, s)
	}
	return s
}

func checksum(s []string) int64 {
	checksum := int64(0)
	for i, x := range s {
		if x == "." {
			continue
		}
		currValue, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(x)
		}
		checksum += int64(currValue * i)
	}
	return checksum
}

func firstProblem() {
	nums := readFiles()
	expandedStr := expand(nums)
	correct := moveToCorrect(expandedStr)
	checksumValue := checksum(correct)
	fmt.Println(checksumValue)
}

func findBlocks(s []string, dot bool) ([]int, []int) {
	starts, lengths := make([]int, 0), make([]int, 0)
	for i := 0; i < len(s); i++ {
		x := s[i]
		filter := false
		if dot {
			filter = x == "."
		} else {
			filter = x != "."
		}

		if filter {
			j := i
			for ; j < len(s); j++ {
				if s[j] != x {
					break
				}
			}
			starts = append(starts, i)
			lengths = append(lengths, j-i)
			i = j - 1
		}
	}
	return starts, lengths
}

func reverseArr(arr []int) []int {
	rev := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		rev[i] = arr[len(arr)-i-1]
	}
	return rev
}

func secondProblem() {
	nums := readFiles()
	expandedStr := expand(nums)

	dotStarts, dotLengths := findBlocks(expandedStr, true)

	numStarts, numLengths := findBlocks(expandedStr, false)
	reversedNumStarts := reverseArr(numStarts)
	reversedNumLengths := reverseArr(numLengths)

	currIdx := 0
	for i := 0; i < len(reversedNumStarts); i++ {
		currNumStart, currNumLength := reversedNumStarts[currIdx], reversedNumLengths[currIdx]

		for j := 0; j < len(dotStarts); j++ {
			currDotStart := dotStarts[j]
			currDotLength := dotLengths[j]
			if currNumLength <= currDotLength && currNumStart > currDotStart {
				for k := range currNumLength {
					tmp := expandedStr[currNumStart+k]
					expandedStr[currNumStart+k] = expandedStr[currDotStart+k]
					expandedStr[currDotStart+k] = tmp
				}
				if dotLengths[j]-currNumLength == 0 {
					dotStarts = append(dotStarts[:j], dotStarts[j+1:]...)
					dotLengths = append(dotLengths[:j], dotLengths[j+1:]...)
				} else {
					dotStarts[j] = dotStarts[j] + currNumLength
					dotLengths[j] = dotLengths[j] - currNumLength
				}
				break
			}
		}
		currIdx++
	}
	checksumValue := checksum(expandedStr)
	fmt.Println(checksumValue)
}

func main() {
	firstProblem()
	secondProblem()
}
