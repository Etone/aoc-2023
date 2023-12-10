package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed example.txt
var input string

type Direction int

const (
	up Direction = iota
	down
	left
	right
)

type Tile struct {
	x, y  int
	value string
}

type Grid [][]Tile

func (grid Grid) findStart() Tile {
	for _, o := range grid {
		for _, i := range o {
			if i.value == "S" {
				return i
			}
		}
	}
	panic("No start found")
}

func (grid Grid) print() {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val.value)
		}
		fmt.Println()
	}
}

var grid Grid = Grid{}

func parseInput() {
	lines := strings.Split(input, "\n")
	grid = make(Grid, len(lines))
	for row, rowVal := range lines {
		grid[row] = make([]Tile, len(rowVal))
		for col, colVal := range rowVal {
			grid[row][col] = Tile{x: col, y: row, value: string(colVal)}
		}
	}
}

func intersectSlices(a, b []string) (intersection []string) {
	hash := make(map[string]bool)
	for _, v := range a {
		hash[v] = false
	}

	for _, v := range b {
		if val, ok := hash[v]; ok && !val {
			intersection = append(intersection, v)
			hash[v] = true
		}
	}
	return
}

func canAccess(tile Tile) (tiles []Tile) {
	tiles = []Tile{}
	if canAccessDirectional(left, tile) {
		tiles = append(tiles, grid[tile.y][tile.x-1])
	}
	if canAccessDirectional(right, tile) {
		tiles = append(tiles, grid[tile.y][tile.x+1])
	}
	if canAccessDirectional(up, tile) {
		tiles = append(tiles, grid[tile.y-1][tile.x])
	}
	if canAccessDirectional(down, tile) {
		tiles = append(tiles, grid[tile.y+1][tile.x])
	}

	return
}

func canAccessDirectional(direction Direction, tile Tile) (canAccess bool) {
	switch direction {
	case up:
		canAccess = tile.y > 0 && strings.Contains("S|JL", tile.value) && strings.Contains("S|7F", grid[tile.y-1][tile.x].value)
	case down:
		canAccess = tile.y < len(grid) && strings.Contains("S|7F", tile.value) && strings.Contains("S|JL", grid[tile.y+1][tile.x].value)
	case left:
		canAccess = tile.x > 0 && strings.Contains("S-7J", tile.value) && strings.Contains("S-FL", grid[tile.y][tile.x-1].value)
	case right:
		canAccess = tile.x < len(grid[0]) && strings.Contains("S-FL", tile.value) && strings.Contains("S-7J", grid[tile.y][tile.x+1].value)
	default:
		canAccess = false
	}
	return
}

// bfs
func findLoop(start Tile) []Tile {
	seen := []Tile{start}
	queue := []Tile{start}

	for len(queue) > 0 {
		//pop from queue
		current := queue[0]
		queue = queue[1:]
		for _, tile := range canAccess(current) {
			if !slices.Contains(seen, tile) {
				queue = append(queue, tile)
				seen = append(seen, tile)
			}
		}
	}
	return seen
}

func replaceStartTile() (start Tile) {
	possibleValues := []string{"|", "-", "J", "F", "L", "7"}
	start = grid.findStart()

	if canAccessDirectional(left, start) {
		possibleValues = intersectSlices(possibleValues, strings.Split("-7J", ""))
	}
	if canAccessDirectional(right, start) {
		possibleValues = intersectSlices(possibleValues, strings.Split("-FL", ""))
	}
	if canAccessDirectional(up, start) {
		possibleValues = intersectSlices(possibleValues, strings.Split("|JL", ""))
	}
	if canAccessDirectional(down, start) {
		possibleValues = intersectSlices(possibleValues, strings.Split("|7F", ""))
	}
	start.value = possibleValues[0]
	grid[start.y][start.x] = start
	return

}

func replaceNonLoopTiles(loop []Tile) {
	for _, row := range grid {
		for _, col := range row {
			if !slices.Contains(loop, col) {
				grid[col.y][col.x] = Tile{y: col.y, x: col.x, value: "."}
			}
		}
	}
}

func findEnclosedTiles(loop []Tile) (enclosed []Tile) {
	for _, row := range grid {
		for _, val := range row {
			if val.value != "." {
				continue
			}
			if raycast(val)%2 != 0 {
				enclosed = append(enclosed, val)
			}
		}
	}
	return
}

func raycast(from Tile) (intersections int) {
	for i := from.x - 1; i > 0; i-- {
		tileToCheck := grid[from.y][i]
		if strings.Contains("|JLF", tileToCheck.value) {
			intersections++
		}
	}
	return
}

func part1() {
	start := grid.findStart()
	loop := findLoop(start)
	fmt.Printf("Part 1: %d\n", len(loop)/2)
}

func part2() {
	start := replaceStartTile()
	loop := findLoop(start)
	replaceNonLoopTiles(loop)
	grid.print()
	fmt.Printf("Part 2: %d\n", len(findEnclosedTiles(loop)))
}

func main() {
	parseInput()
	part1()
	part2()
}
