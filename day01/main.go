package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func panick(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	var aList, bList []int
	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Split(line, "   ")
		a, _ := strconv.Atoi(str[0])
		b, _ := strconv.Atoi(str[1])
		aList = append(aList, a)
		bList = append(bList, b)
	}
	return aList, bList
}

func calcPartOne(sliceA, sliceB []int) {
	sort.Ints(sliceA)
	sort.Ints(sliceB)

	var finalAns int
	for i := range sliceA {
		ans := sliceA[i] - sliceB[i]
		if ans < 0 {
			ans = ans * -1
		}
		finalAns = finalAns + ans
	}
	fmt.Printf("Answer for Part 1: %v\n", finalAns)
}

func calcPartTwo(sliceA, sliceB []int) {
	var totalCounter int
	for _, v := range sliceA {
		var counter int
		for _, seVal := range sliceB {
			if v == seVal {
				counter++
			}
		}
		totalCounter += v * counter
	}
	fmt.Printf("Answer for Part 1: %v\n", totalCounter)
}

func main() {
	file, err := os.Open("input.txt")
	panick(err)
	sliceA, sliceB := parseFile(file)

	// partOne
	// calcPartOne(sliceA, sliceB)

	// partTwo
	calcPartTwo(sliceA, sliceB)
}
