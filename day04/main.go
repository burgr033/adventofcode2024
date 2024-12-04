package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintMatrix(matrix [DIMENSION][DIMENSION]string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func checkMatrixForMAS(matrix [DIMENSION][DIMENSION]string) int {
	var counter int
	for i := 1; i < DIMENSION-1; i++ { // Iterate rows, skipping the first and last
		for j := 1; j < DIMENSION-1; j++ { // Iterate columns, skipping the first and last
			if matrix[i][j] == "A" {
				// Check both diagonals for X-MAS pattern
				if ((matrix[i-1][j-1] == "M" && matrix[i+1][j+1] == "S") || (matrix[i-1][j-1] == "S" && matrix[i+1][j+1] == "M")) &&
					((matrix[i-1][j+1] == "M" && matrix[i+1][j-1] == "S") || (matrix[i-1][j+1] == "S" && matrix[i+1][j-1] == "M")) {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
	return counter
}

func buildMatrix(horStrings []string) [DIMENSION][DIMENSION]string {
	var matrix [DIMENSION][DIMENSION]string
	for i, v := range horStrings {
		for j, r := range v {
			matrix[i][j] = string(r)
		}
	}
	return matrix
}

func getDiagonalStrings(horStrings []string) []string {
	var diagonals []string
	rows := len(horStrings)
	cols := len(horStrings[0])

	for d := 0; d < rows+cols-1; d++ {
		var diag1, diag2 string
		for i := 0; i <= d; i++ {
			j := d - i
			if i < rows && j < cols {
				diag1 += string(horStrings[i][j])
				diag2 += string(horStrings[rows-1-i][j])
			}
		}
		diagonals = append(diagonals, diag1, diag2)
	}
	return diagonals
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func checkString(toCheck string, keyword string) int {
	counter := 0
	counter += strings.Count(toCheck, keyword)
	counter += strings.Count(reverseString(toCheck), keyword)
	return counter
}

func parseFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var horStrings []string
	for scanner.Scan() {
		line := scanner.Text()
		horStrings = append(horStrings, line)
	}
	return horStrings
}

func getVerStrings(horStrings []string) []string {
	maxLen := len(horStrings[0])
	var verStrings []string
	for i := 0; i < maxLen; i++ {
		var verString string
		for j := 0; j < len(horStrings); j++ {
			verString += string(horStrings[j][i])
		}
		verStrings = append(verStrings, verString)
	}
	return verStrings
}

// this is very ugly with the constant and the 2D array...
const (
	DIMENSION = 140
	keyword   = "XMAS"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	horStrings := parseFile(file)
	verStrings := getVerStrings(horStrings)
	diagStrings := getDiagonalStrings(horStrings)

	var counter int
	for _, v := range horStrings {
		counter += checkString(v, keyword)
	}
	for _, v := range verStrings {
		counter += checkString(v, keyword)
	}
	for _, v := range diagStrings {
		counter += checkString(v, keyword)
	}
	fmt.Printf("PartOne:%v\n", counter)
	matrix := buildMatrix(horStrings)
	PrintMatrix(matrix)
	fmt.Println("------")
	_ = checkMatrixForMAS(matrix)
}
