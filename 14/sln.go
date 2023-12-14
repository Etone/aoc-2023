package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed example.txt
var input string

type Grid [][]Value

type Value int

type Direction int

const (
	Empty Value = iota
	FixStone
	RollingStone
)

const (
	North Direction = iota
	West
	South
	East
)

var allDirections = []Direction{North, West, South, East}

func (grid Grid) Print() {
	for _, v := range grid {
		fmt.Println(v)
	}
}

func (grid *Grid) Tilt(direction Direction) (moved bool) {
	moved = false
	for r, row := range *grid {
		for c, value := range row {
			if value == RollingStone {
				moved = moved || grid.TiltInDirection(direction, r, c)
			}
		}
	}
	return
}

func (grid *Grid) TiltInDirection(direction Direction, row, col int) bool {
	switch direction {
	case North:
		if row == 0 {
			return false
		}
		if (*grid)[row-1][col] == Empty {
			(*grid)[row][col] = Empty
			(*grid)[row-1][col] = RollingStone
			return true
		}
	case South:
		if row == len(*grid)-1 {
			return false
		}
		if (*grid)[row+1][col] == Empty {
			(*grid)[row][col] = Empty
			(*grid)[row+1][col] = RollingStone
			return true
		}

	case West:
		if col == 0 {
			return false
		}
		if (*grid)[row][col-1] == Empty {
			(*grid)[row][col] = Empty
			(*grid)[row][col-1] = RollingStone
			return true
		}

	case East:
		if col == len((*grid)[0])-1 {
			return false
		}
		if (*grid)[row][col+1] == Empty {
			(*grid)[row][col] = Empty
			(*grid)[row][col+1] = RollingStone
			return true
		}
	}
	return false
}

func (grid Grid) Load() (result int) {
	for r, row := range grid {
		for _, value := range row {
			if value == RollingStone {
				result += len(grid) - r
			}
		}
	}
	return
}

func (grid *Grid) Cycle() Grid{
	for dir := 0; dir < len(allDirections); dir++ {
		for grid.Tilt(allDirections[dir]) {
			// Nothing here, just cycle
		}
	}
	return *grid
}

func parse() (grid Grid) {
	for r, row := range strings.Split(input, "\n") {
		grid = append(grid, []Value{})
		for _, col := range row {
			grid[r] = append(grid[r], toFieldValue(col))
		}
	}
	return
}

func toFieldValue(v rune) Value {
	switch v {
	case '#':
		return FixStone
	case 'O':
		return RollingStone
	default:
		return Empty
	}
}

func part1() {
	grid := parse()

	for {
		if !grid.Tilt(North) {
			break
		}
	}
	fmt.Printf("Part 1: %d\n", grid.Load())
}

func part2() {
	// Currently this would solve Part 2 after ages
	grid := parse()
	for i := 0; i < 3; i++ {
		grid.Cycle()
	}
	fmt.Printf("Part 2: %d", grid.Load())
}

func main() {
	part1()
	part2()
}
