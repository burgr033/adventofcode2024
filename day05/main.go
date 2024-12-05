package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkValidity(rules [][2]int, update []int) bool {
	pos := make(map[int]int)
	for i, page := range update {
		pos[page] = i
	}
	for _, rule := range rules {
		if posX, okX := pos[rule[0]]; okX {
			if posY, okY := pos[rule[1]]; okY {
				if posX >= posY {
					return false
				}
			}
		}
	}
	return true
}

func parseFile(file *os.File) ([][2]int, [][]int) {
	scanner := bufio.NewScanner(file)
	var firstSection [][2]int
	var secondSection [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			var newArr [2]int
			l := strings.Split(line, "|")
			newArr[0], _ = strconv.Atoi(l[0])
			newArr[1], _ = strconv.Atoi(l[1])

			firstSection = append(firstSection, newArr)
		}
		if strings.Contains(line, ",") {
			arr := strings.Split(line, ",")
			var newArr []int
			for _, element := range arr {
				num, _ := strconv.Atoi(element)
				newArr = append(newArr, num)
			}
			secondSection = append(secondSection, newArr)
		}

	}
	return firstSection, secondSection
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rules, updates := parseFile(file)
	var ans int
	for _, upd := range updates {
		if checkValidity(rules, upd) {
			mid := upd[len(upd)/2]
			ans += mid
		} else {
			fmt.Println(upd)
		}
	}
	fmt.Printf("Answer Part 1: %v\n", ans)
}
