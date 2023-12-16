package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed example.txt
var input string

type Tile struct {
	value     TileValue
	energized bool
}

type Grid [][]Tile

func (grid Grid) getNextTilePosition(dir Direction, r, c int) (Position, bool) {
	switch dir {
	case Up:
		if r <= 0 {
			return Position{}, false
		}
		return Position{r - 1, c}, true
	case Down:
		if r >= len(grid) {
			return Position{}, false
		}
		return Position{r + 1, c}, true
	case Left:
		if c <= 0 {
			return Position{}, false
		}
		return Position{r, c - 1}, true
	case Right:
		if c >= len(grid[0]) {
			return Position{}, false
		}
		return Position{r, c + 1}, true
	default:
		panic("Direction not correct")
	}
}

func (grid Grid) energizedCells() (num int) {
	for _, row := range grid {
		for _, tile := range row {
			if tile.energized {
				num++
			}
		}
	}
	return
}

func (grid Grid) printEnergized() {
	for _, row := range grid {
		fmt.Println()
		for _, tile := range row {
			if tile.energized {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

	}
}

func (grid Grid) get(p Position) Tile {
	return grid[p.row][p.col]
}

type Position struct {
	row, col int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Beam struct {
	pos Position
	dir Direction
}

type TileValue int

const (
	Empty TileValue = iota
	HorizontalSplitter
	VerticalSplitter
	LeftDiagonalMirror
	RightDiagonalMirror
)

func toTileValue(v rune) TileValue {
	switch v {
	case '.':
		return Empty
	case '|':
		return VerticalSplitter
	case '-':
		return HorizontalSplitter
	case '/':
		return RightDiagonalMirror
	case '\\':
		return LeftDiagonalMirror
	default:
		panic("No such TileValue")
	}
}

func parse() (grid Grid) {
	for r, row := range strings.Split(input, "\n") {
		grid = append(grid, []Tile{})
		for _, ch := range row {
			tv := toTileValue(ch)
			grid[r] = append(grid[r], Tile{energized: false, value: tv})
		}
	}
	return
}

func getNextDirectionsBasedOnPosition(grid Grid, beam Beam) []Direction {
	tv := grid.get(beam.pos).value
	switch tv {
	case Empty:
		return []Direction{beam.dir}
	case HorizontalSplitter:
		if beam.dir == Up || beam.dir == Down {
			return []Direction{Left, Right}
		}
		return []Direction{beam.dir}
	case VerticalSplitter:
		if beam.dir == Left || beam.dir == Right {
			return []Direction{Up, Down}
		}
		return []Direction{beam.dir}
	case RightDiagonalMirror:
		{
			switch beam.dir {
			case Right:
				return []Direction{Up}
			case Left:
				return []Direction{Down}
			case Up:
				return []Direction{Left}
			case Down:
				return []Direction{Right}
			default:
				panic("NYI")
			}
		}
	case LeftDiagonalMirror:
		{
			switch beam.dir {
			case Right:
				return []Direction{Down}
			case Left:
				return []Direction{Up}
			case Up:
				return []Direction{Right}
			case Down:
				return []Direction{Left}
			default:
				panic("NYI")
			}
		}
	default:
		panic("NYI")
	}
}

func part1() {
	grid := parse()
	lightBeams := []Beam{{pos: Position{0, 0}, dir: Right}}

	seenBeams := []Beam{}

	for len(lightBeams) > 0 {
		beam := lightBeams[0]
		lightBeams = lightBeams[1:]
		grid[beam.pos.row][beam.pos.col].energized = true

		if slices.Contains(seenBeams, beam) {
			continue
		}

		seenBeams = append(seenBeams, beam)
		dirs := getNextDirectionsBasedOnPosition(grid, beam)
		for _, dir := range dirs {
			if nextPos, ok := grid.getNextTilePosition(dir, beam.pos.row, beam.pos.col); ok {
				lightBeams = append(lightBeams, Beam{nextPos, dir})
			}
		}
	}
	grid.printEnergized()
	fmt.Printf("Part 1: %d\n", grid.energizedCells())
}

func part2() {
}

func main() {
	part1()
	part2()
}
