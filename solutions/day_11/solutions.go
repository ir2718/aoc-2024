package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func firstProblem(s string, numBlinks int) {
	splitStr := strings.Fields(s)
	m := map[string]int{}
	for _, s := range splitStr {
		m[s] = 1
	}

	for numBlinks > 0 {
		newStr := map[string]int{}

		for s := range m {
			if s == "0" {
				_, ok := newStr["1"]
				if ok {
					newStr["1"] += m[s]
				} else {
					newStr["1"] = m[s]
				}

			} else if len(s)%2 == 0 {
				left, err := strconv.Atoi(s[:len(s)/2])
				if err != nil {
					log.Fatal(err)
				}
				right, err := strconv.Atoi(s[len(s)/2:])
				if err != nil {
					log.Fatal(err)
				}

				leftStr := strconv.Itoa(left)
				_, ok := newStr[leftStr]
				if ok {
					newStr[leftStr] += m[s]
				} else {
					newStr[leftStr] = m[s]
				}

				rightStr := strconv.Itoa(right)
				_, ok = newStr[rightStr]
				if ok {
					newStr[rightStr] += m[s]
				} else {
					newStr[rightStr] = m[s]
				}

			} else {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				numStr := strconv.Itoa(num * 2024)

				_, ok := m[numStr]
				if ok {
					newStr[numStr] += m[s]
				} else {
					newStr[numStr] = m[s]
				}

			}
		}
		m = newStr
		numBlinks--
	}

	total := 0
	for _, v := range m {
		total += v
	}
	fmt.Println(total)
}

func main() {
	inputString := "510613 358 84 40702 4373582 2 0 1584"
	firstProblem(inputString, 25)
	firstProblem(inputString, 75)
}
