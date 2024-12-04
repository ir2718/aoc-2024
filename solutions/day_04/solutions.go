package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CoordString struct {
	col, row                   int
	currWord                   []string
	colDirection, rowDirection int
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func (s *Stack[T]) Pop() T {
	if s.Len() == 0 {
		log.Fatal("Calling `pop()` is not possible as stack is empty.")
	}
	elem := s.elements[s.Len()-1]
	s.elements = s.elements[:s.Len()-1]
	return elem
}

func (s *Stack[T]) Push(elem *T) {
	s.elements = append(s.elements, *elem)
}

func readFiles() [][]string {
	file, err := os.Open("files/day_04/problem_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	letters := make([][]string, 0)
	for scanner.Scan() {
		row := scanner.Text()
		rowSplit := strings.Split(row, "")
		letters = append(letters, rowSplit)
	}
	return letters
}

func firstProblem() {
	letters := readFiles()

	// Assign the word to each coordinate in the map.
	word := "XMAS"
	splitWord := strings.Split(word, "")
	open := Stack[CoordString]{}

	// Left, right, top, bottom positions.
	l_i, l_j := 0, -1
	r_i, r_j := 0, 1
	t_i, t_j := -1, 0
	b_i, b_j := 1, 0

	// Diagonal positions.
	lt_i, lt_j := -1, -1
	rt_i, rt_j := -1, 1
	lb_i, lb_j := 1, -1
	rb_i, rb_j := 1, 1

	rowDirections := []int{l_i, r_i, t_i, b_i, lt_i, rt_i, lb_i, rb_i}
	colDirections := []int{l_j, r_j, t_j, b_j, lt_j, rt_j, lb_j, rb_j}
	for i, row := range letters {
		for j := range row {
			if letters[i][j] == splitWord[0] {
				for k := range rowDirections {
					rowDirection := rowDirections[k]
					colDirection := colDirections[k]
					open.Push(&CoordString{j, i, splitWord[1:], colDirection, rowDirection})
				}
			}
		}
	}

	// Do DFS starting from each coordinate and sum the number of successful DFSs.
	numFound := 0

	for open.Len() != 0 {
		coordString := open.Pop()

		i, j := coordString.row, coordString.col
		currWord := coordString.currWord
		colDirection, rowDirection := coordString.colDirection, coordString.rowDirection

		newRow, newCol := i+rowDirection, j+colDirection

		if newRow > (len(letters)-1) || newCol > (len(letters[0])-1) || newRow < 0 || newCol < 0 {
			continue
		}

		if letters[newRow][newCol] == currWord[0] {
			// In the case there is a match and this is the last letter,
			// it means the word is found.
			if len(currWord) == 1 {
				numFound++
			} else {
				open.Push(&CoordString{newCol, newRow, currWord[1:], colDirection, rowDirection})
			}
		}
	}
	fmt.Println(numFound)
}

func secondProblem() {

	letters := readFiles()

	// Assign the word to each coordinate in the map.
	word := "MAS"
	splitWord := strings.Split(word, "")

	foundNum := 0
	for i, row := range letters {
		for j := range row {
			if i-1 < 0 || j-1 < 0 || i+1 >= len(letters) || j+1 >= len(letters[0]) {
				continue
			}

			// Detect candidates for the X shaped 'MAS' and check for symmetry.
			if letters[i][j] == splitWord[1] && (((letters[i-1][j-1] == "M" && letters[i+1][j-1] == "S") ||
				(letters[i-1][j-1] == "S" && letters[i+1][j-1] == "M")) &&
				letters[i-1][j-1] == letters[i-1][j+1] && letters[i+1][j-1] == letters[i+1][j+1]) {
				// Case of M/S above A and S/M below A.
				foundNum += 1
			} else if letters[i][j] == splitWord[1] && ((letters[i-1][j-1] == "M" && letters[i-1][j+1] == "S") ||
				(letters[i-1][j-1] == "S" && letters[i-1][j+1] == "M")) &&
				letters[i-1][j-1] == letters[i+1][j-1] && letters[i-1][j+1] == letters[i+1][j+1] {
				// Case of M/S left of A and S/M right of A.
				foundNum += 1
			}
		}
	}
	fmt.Println(foundNum)
}

func main() {
	firstProblem()
	secondProblem()
}
