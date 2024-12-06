package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row       int
	col       int
	direction string
}

type grid struct {
	data             [][]string
	startingPosition position
	currentPosition  position
	dir              string
}

func (g *grid) rows() int {
	return len(g.data)
}

func (g *grid) cols() int {
	return len(g.data[0])
}

func (g *grid) prettyPrint() {
	for _, line := range g.data {
		fmt.Printf("%v\n", line)
	}
	fmt.Println()
	fmt.Printf("Starting Position: %v|%v\nCurrent Position: %v|%v\n", g.startingPosition.row, g.startingPosition.col, g.currentPosition.row, g.currentPosition.col)
}

func (g *grid) isEdge() bool {
	row := g.currentPosition.row
	col := g.currentPosition.col

	// Check if the current position is on any edge
	return row == 0 || col == 0 || row == g.cols()-1 || col == g.rows()-1
}

func (g *grid) isFree() bool {
	switch g.currentPosition.direction {
	case "up":
		if g.data[g.currentPosition.col-1][g.currentPosition.row] == "#" {
			return false
		}
	case "down":
		if g.data[g.currentPosition.col+1][g.currentPosition.row] == "#" {
			return false
		}
	case "left":
		if g.data[g.currentPosition.col][g.currentPosition.row-1] == "#" {
			return false
		}
	case "right":
		if g.data[g.currentPosition.col][g.currentPosition.row+1] == "#" {
			return false
		}
	}
	return true
}

func (g *grid) finalCount() int {
	var countX int

	for _, row := range g.data {
		for _, cell := range row {
			switch cell {
			case "X":
				countX++
			case "O":
				countX++
			}
		}
	}

	return countX
}

func (g *grid) rotateRight() {
	switch g.currentPosition.direction {
	case "up":
		g.currentPosition.direction = "right"
	case "right":
		g.currentPosition.direction = "down"
	case "down":
		g.currentPosition.direction = "left"
	case "left":
		g.currentPosition.direction = "up"
	}
}

func (g *grid) move() {
	switch g.currentPosition.direction {
	case "up":
		g.currentPosition.col -= 1
	case "down":
		g.currentPosition.col += 1
	case "left":
		g.currentPosition.row -= 1
	case "right":
		g.currentPosition.row += 1
	}
}

func (g *grid) setIcon() {
	x := g.currentPosition.col
	y := g.currentPosition.row
	g.data[x][y] = "O"
}

func (g *grid) removeIcon() {
	x := g.currentPosition.col
	y := g.currentPosition.row

	g.data[x][y] = "X"
}

func findStartingPosition(grid [][]string) (x, y int, direction string) {
	for i, l := range grid {
		for j, v := range l {
			switch v {
			case "v":
				return i, j, "down"
			case "^":
				return i, j, "up"
			case ">":
				return i, j, "right"
			case "<":
				return i, j, "left"

			}
		}
	}
	return 0, 0, "error"
}

func parseFile(file *os.File) grid {
	scanner := bufio.NewScanner(file)
	var m [][]string

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for j, c := range line {
			row[j] = string(c)
		}
		m = append(m, row)
	}
	x, y, d := findStartingPosition(m)
	grid := &grid{
		data:             m,
		startingPosition: position{row: y, col: x, direction: d},
		currentPosition:  position{row: y, col: x, direction: d},
	}
	return *grid
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := parseFile(file)
	grid.prettyPrint()

	for {
		grid.removeIcon()
		if !grid.isFree() {
			grid.rotateRight()
		}

		grid.move()

		grid.setIcon()

		if grid.isEdge() {
			fmt.Println("Reached the edge of the grid.")
			break
		}
	}
	grid.prettyPrint()
	fmt.Println(grid.finalCount())
}
