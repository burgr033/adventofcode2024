package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	keyword := "XMAS"
	file, err := os.Open("test_input.txt")
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
	fmt.Println(counter)
}
