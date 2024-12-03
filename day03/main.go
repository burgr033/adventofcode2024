package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func findMatchesPartOne(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	var ans int
	for _, match := range matches {
		n, err1 := strconv.Atoi(match[1])
		m, err2 := strconv.Atoi(match[2])
		if err1 != nil || err2 != nil {
			panic("rip")
		}
		ans += n * m
	}
	return ans
}

func findMatchesPartTwo(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	enabled := true
	var ans int
	for _, match := range matches {
		fmt.Println(match[0])
		if match[0] == "don't()" {
			enabled = false
		}
		if match[0] == "do()" {
			enabled = true
			continue
		}
		if enabled {
			n, err1 := strconv.Atoi(match[1])
			m, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				panic(err1)
			}
			ans += n * m
		}
	}
	return ans
}

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(content)
	return text
}

func main() {
	content := ReadFile("input.txt")
	ansPartOne := findMatchesPartOne(content)
	ansPartTwo := findMatchesPartTwo(content)
	fmt.Printf("Answer Part 1: %v\n", ansPartOne)
	fmt.Printf("Answer Part 2: %v\n", ansPartTwo)
}
