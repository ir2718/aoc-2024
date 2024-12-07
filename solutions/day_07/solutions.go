package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFiles() ([]int, [][]int) {
	file, err := os.Open("files/day_07/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scores := make([]int, 0)
	inputs := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")

		score, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		scores = append(scores, score)

		nums := strings.Fields(lineSplit[1])
		numsInt := make([]int, len(nums))
		for i, n := range nums {
			nInt, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numsInt[i] = nInt
		}
		inputs = append(inputs, numsInt)
	}
	return scores, inputs
}

func convertToBaseNString(n int, b int, decimals int) []string {
	stringRep := make([]string, decimals)
	for i := range stringRep {
		stringRep[i] = "0"
	}

	pos := decimals - 1
	for n > 0 && pos >= 0 {
		digit := n % b
		stringRep[pos] = strconv.Itoa(digit)
		n = int(math.Floor(float64(n) / float64(b)))
		pos--
	}

	return stringRep
}

func generateNaryCombMasks(n int, base int) [][]string {
	numElems := int(math.Pow(float64(base), float64(n)))
	masks := make([][]string, numElems)
	for i := 0; i < numElems; i++ {
		masks[i] = convertToBaseNString(i, base, n)
	}
	return masks
}

func firstProblem() {
	scores, inputs := readFiles()

	calibration := 0
	for i := 0; i < len(inputs); i++ {
		currInput := inputs[i]
		currScore := scores[i]
		masks := generateNaryCombMasks(len(currInput)-1, 2)
		for _, mask := range masks {
			result := currInput[0]
			for j, m := range mask {
				if m == "0" {
					result = result + currInput[j+1]
				} else {
					result = result * currInput[j+1]
				}
			}

			if result == currScore {
				calibration += currScore
				break
			}
		}
	}
	fmt.Println(calibration)
}

func secondProblem() {
	scores, inputs := readFiles()

	calibration := 0
	for i := 0; i < len(inputs); i++ {
		currInput := inputs[i]
		currScore := scores[i]
		masks := generateNaryCombMasks(len(currInput)-1, 3)
		for _, mask := range masks {
			result := currInput[0]
			for j, m := range mask {
				if m == "0" {
					result = result + currInput[j+1]
				} else if m == "1" {
					result = result * currInput[j+1]
				} else {
					str := fmt.Sprintf("%d%d", result, currInput[j+1])
					conversion, err := strconv.Atoi(str)
					if err != nil {
						log.Fatal(err)
					}
					result = conversion
				}
			}

			if result == currScore {
				calibration += currScore
				break
			}
		}
	}
	fmt.Println(calibration)
}

func main() {
	firstProblem()
	secondProblem()
}
