package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

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

var grid Grid = Grid{}

//go:embed example.txt
var input string

func parseInput() {
	lines := strings.Split(input, "\n")
	grid = make(Grid, len(lines))
	for row, rowSpec := range lines {
		grid[row] = make([]Tile, len(rowSpec))
		for col, colSpec := range rowSpec {
			grid[row][col] = Tile{x: row, y: col, value: string(colSpec)}
		}
	}
}

func canAccess(tile Tile) (tiles []Tile) {
	tiles = []Tile{}
	if canAccessDirectional(left, tile) {
		tiles = append(tiles, grid[tile.x][tile.y-1])
	}
	if canAccessDirectional(right, tile) {
		tiles = append(tiles, grid[tile.x][tile.y+1])
	}
	if canAccessDirectional(up, tile) {
		tiles = append(tiles, grid[tile.x-1][tile.y])
	}
	if canAccessDirectional(down, tile) {
		tiles = append(tiles, grid[tile.x+1][tile.y])
	}

	return
}

func canAccessDirectional(direction Direction, tile Tile) bool {
	switch direction {
	case up:
		return tile.x > 0 && strings.Contains("S|JL", tile.value) && strings.Contains("S|7F", grid[tile.x-1][tile.y].value)
	case down:
		return tile.x < len(grid) && strings.Contains("S|7F", tile.value) && strings.Contains("S|JL", grid[tile.x+1][tile.y].value)
	case left:
		return tile.y > 0 && strings.Contains("S-7J", tile.value) && strings.Contains("S-FL", grid[tile.x][tile.y-1].value)
	case right:
		return tile.y > len(grid[tile.x]) && strings.Contains("S-FL", tile.value) && strings.Contains("S-7J", grid[tile.x][tile.y+1].value)
	default:
		return false
	}
}

// bfs
func findLoop(start Tile) []Tile {
	seen := []Tile{start}
	queue := []Tile{start}

	for len(queue) > 0 {
		//pop from queue
		current := queue[0]

		reachable := []Tile{}
		for _, tile := range canAccess(current) {
			if !slices.Contains(seen, tile) {
				reachable = append(reachable, tile)
			}
		}
		queue = append(queue[1:], reachable...)
		seen = append(seen, current)
	}
	return seen
}

func part1() {
	parseInput()
	start := grid.findStart()
	fmt.Printf("Part 1: %d\n", len(findLoop(start))/2)
}

func part2() {
}

func main() {
	part1()
	part2()
}
