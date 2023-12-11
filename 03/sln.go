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
	adjacent := getAdjacentCellValues(position, grid)
	for _, cell := range adjacent {
		if !unicode.IsDigit(cell) && cell != '.' {
			return true
		}
	}
	return false
}

func getAdjacentCellValues(position position, grid grid) (adjacent []cell) {
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

func getAdjacentCellPositions(pos position, grid grid) (adjacent []position) {
	xTop := max(pos.x-1, 0)
	xBot := min(pos.x+1, len(grid)-1)
	yLeft := max(pos.y-1, 0)
	yRight := min(pos.y+1, len(grid[0])-1)

	topLeft := position{x: xTop, y: yLeft, cell: grid[xTop][yLeft]}
	top := position{x: xTop, y: pos.y, cell: grid[xTop][pos.y]}
	topRight := position{x: xTop, y: yRight, cell: grid[xTop][yRight]}

	left := position{x: pos.x, y: yLeft, cell: grid[pos.x][yLeft]}
	right := position{x: pos.x, y: yRight, cell: grid[pos.x][yRight]}

	botLeft := position{x: xBot, y: yLeft, cell: grid[xBot][yLeft]}
	bot := position{x: xBot, y: pos.y, cell: grid[xBot][pos.y]}
	botRight := position{x: xBot, y: yRight, cell: grid[xBot][yRight]}

	adjacent = append(adjacent, topLeft, top, topRight, left, right, botLeft, bot, botRight)

	return
}

func gear(pos position, grid grid) (bool, position) {
	adjacent := getAdjacentCellPositions(pos, grid)
	for _, p := range adjacent {
		if p.cell == '*' {
			return true, p
		}
	}
	return false, pos
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
				isTouching = isTouching || isPartNumber(pos, grid)
			} else if currentNumber.Len() != 0 {
				if isTouching {
					partNumber, _ := strconv.Atoi(currentNumber.String())
					sumOfPartNumbers += partNumber
				}
				currentNumber.Reset()
				isTouching = false
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sumOfPartNumbers)
}

func part2() {
	grid := readFileToGrid("input_01.txt")

	sumOfGearRatios := 0

	gearsAndCells := map[position][]int{}

	for row, cells := range grid {
		var currentNumber strings.Builder
		isNextToGear := false
		var gearNextTo position

		for col, cell := range cells {

			pos := position{x: row, y: col, cell: cell}

			if unicode.IsDigit(cell) {
				currentNumber.WriteRune(cell)
				exists, gearPos := gear(pos, grid)
				isNextToGear = isNextToGear || exists
				if exists {
					gearNextTo = gearPos
				}
			} else if currentNumber.Len() != 0 {
				if isNextToGear {
					current, _ := strconv.Atoi(currentNumber.String())
					gearsAndCells[gearNextTo] = append(gearsAndCells[gearNextTo], current)
				}
				currentNumber.Reset()
				isNextToGear = false
			}
		}

	}
	for _, numbers := range gearsAndCells {
		if len(numbers) == 2 {
			sumOfGearRatios += numbers[0] * numbers[1]
		}
	}

	fmt.Printf("Part 2: %d\n", sumOfGearRatios)
}
