package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type report struct {
	levels []int
	safe   bool
}

func panick(e error) {
	if e != nil {
		panic(e)
	}
}

func isSaveAfterRemovingOne(num []int) bool {
	for i := 0; i < len(num); i++ {
		numWith := append([]int{}, num[:i]...)
		numWith = append(numWith, num[i+1:]...)
		if isIncreasing(numWith) || isDecreasing(numWith) {
			return true
		}
	}
	return false
}

func isDecreasing(num []int) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] < num[i+1] {
			return false
		}
		if num[i]-num[i+1] < 1 || num[i]-num[i+1] > 3 {
			return false
		}
	}
	return true
}

func isIncreasing(num []int) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			return false
		}

		if num[i+1]-num[i] < 1 || num[i+1]-num[i] > 3 {
			return false
		}
	}
	return true
}

func parseFile(file *os.File) []report {
	scanner := bufio.NewScanner(file)
	var reportsss []report
	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Split(line, " ")
		var rep report
		rep.safe = false
		for _, v := range str {
			num, err := strconv.Atoi(v)
			panick(err)
			rep.levels = append(rep.levels, num)
		}
		reportsss = append(reportsss, rep)
	}
	return reportsss
}

func main() {
	file, err := os.Open("input.txt")
	panick(err)
	reports := parseFile(file)
	for i := range reports {
		if isIncreasing(reports[i].levels) || isDecreasing(reports[i].levels) {
			reports[i].safe = true
		} else {
			reports[i].safe = false
		}
	}
	var counterPartOne int
	var counterPartTwo int
	for _, v := range reports {
		if v.safe {
			counterPartOne++
		} else if isSaveAfterRemovingOne(v.levels) {
			counterPartTwo++
		}
	}
	fmt.Printf("part 1: %v\n", counterPartOne)
	fmt.Printf("part 2: %v\n", counterPartOne+counterPartTwo)
}
