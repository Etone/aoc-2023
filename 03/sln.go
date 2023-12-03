package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type grid = [][]cell

type cell = rune

type position struct {
	x    int
	y    int
	cell cell
}

func main() {
	part1()
	part2()
}

func readFileToGrid(path string) (grid grid) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		grid = append(grid, runes)
	}
	return
}

func isPartNumber(position position, grid grid) bool {
	adjacent := getAdjacentCells(position, grid)
	for _, cell := range adjacent {
		if !unicode.IsDigit(cell) && cell != '.' {
			return true
		}
	}
	return false
}

func getAdjacentCells(position position, grid grid) (adjacent []cell) {
	xTop := max(position.x-1, 0)
	xBot := min(position.x+1, len(grid)-1)
	yLeft := max(position.y-1, 0)
	yRight := min(position.y+1, len(grid[0])-1)

	topLeft := grid[xTop][yLeft]
	top := grid[xTop][position.y]
	topRight := grid[xTop][yRight]
	left := grid[position.x][yLeft]
	right := grid[position.x][yRight]
	botLeft := grid[xBot][yLeft]
	bot := grid[xBot][position.y]
	botRight := grid[xBot][yRight]

	adjacent = append(adjacent, topLeft, top, topRight, left, right, botLeft, bot, botRight)

	return
}

func part1() {
	grid := readFileToGrid("input_01.txt")

	sumOfPartNumbers := 0

	for row, cells := range grid {
		var currentNumber strings.Builder
		isTouching := false
		for col, cell := range cells {
			pos := position{cell: cell, x: row, y: col}

			if unicode.IsDigit(cell) {
				currentNumber.WriteRune(cell)
				println(currentNumber.String())
				isTouching = isTouching || isPartNumber(pos, grid)
			} else if currentNumber.Len() != 0 {
				if isTouching {
					partNumber, _ := strconv.Atoi(currentNumber.String())
					sumOfPartNumbers += partNumber
					println(partNumber, sumOfPartNumbers)
				}
				currentNumber.Reset()
				isTouching = false
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sumOfPartNumbers)
}

func part2() {
	grid := readFileToGrid("example.txt")

	sumOfGearRatios := 0

	for row, cells := range grid {
		var currentNumber strings.Builder
		isTouching := false
		for col, cell := range cells {
			pos := position{cell: cell, x: row, y: col}
		}
	}

	fmt.Printf("Part 2: %d\n", sumOfGearRatios)
}
