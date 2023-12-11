package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

var grid []Point

var emptyRows []int
var emptyCol []int
var galaxies []Point

func parseToGrid() {
	grid = []Point{}

	for r, row := range strings.Split(input, "\n") {
		for c, col := range strings.Split(row, "") {
			point := Point{}
			point.x = r
			point.y = c
			grid = append(grid, point)

			if col == "#" {
				galaxies = append(galaxies, point)
			}
		}
	}
}

func manhattenDistanceWithScale(scale int, from, to Point) (distance int) {
	minX := min(from.x, to.x)
	maxX := max(from.x, to.x)
	minY := min(from.y, to.y)
	maxY := max(from.y, to.y)

	for row := minX; row < maxX; row++ {
		if slices.Contains(emptyRows, row) {
			distance += scale - 1
		}
		distance += 1
	}
	for col := minY; col < maxY; col++ {
		if slices.Contains(emptyCol, col) {
			distance += scale - 1
		}
		distance += 1
	}
	return
}

func setEmptyRowsAndCols() {
	rows := strings.Split(input, "\n")
	cols := []string{}

	var col strings.Builder
	for y := 0; y < len(rows[0]); y++ {
		for i := 0; i < len(rows); i++ {
			col.WriteByte(rows[i][y])
		}
		cols = append(cols, col.String())
		col.Reset()
	}
	for r, row := range rows {
		if !strings.Contains(row, "#") {
			emptyRows = append(emptyRows, r)
		}
	}
	for c, col := range cols {
		if !strings.Contains(col, "#") {
			emptyCol = append(emptyCol, c)
		}
	}
}

func part1And2() {
	sumOfDistances1 := 0
	sumOfDistances2 := 0

	for _, from := range galaxies {
		for _, to := range galaxies {
			sumOfDistances1 += manhattenDistanceWithScale(2, from, to)
			sumOfDistances2 += manhattenDistanceWithScale(1_000_000, from, to)
		}
	}
	// I count all galaxies twice, no need for that
	fmt.Printf("Part 1: %d\n", sumOfDistances1/2)
	fmt.Printf("Part 2: %d\n", sumOfDistances2/2)
}

func main() {
	parseToGrid()
	setEmptyRowsAndCols()
	part1And2()
}
