package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func columns(s string) (cols []string) {
	rows := strings.Split(s, "\n")
	var col strings.Builder
	for y := 0; y < len(rows[0]); y++ {
		for i := 0; i < len(rows); i++ {
			col.WriteByte(rows[i][y])
		}
		cols = append(cols, col.String())
		col.Reset()
	}
	return
}

func rows(s string) []string {
	return strings.Split(s, "\n")
}

func reverse(r []string) (reverse []string) {
	for k := range r {
		reverse = append(reverse, r[len(r)-1-k])
	}
	return
}

func mirror(toMirror []string, compare func([]string, []string) bool) int {
	for i := 1; i < len(toMirror); i++ {
		above := reverse(toMirror[:i])
		below := toMirror[i:]

		length := min(len(above), len(below))
		above = above[:length]
		below = below[:length]

		if compare(above, below) {
			return i
		}
	}
	return 0
}

func differences(a, b string) (differences int) {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}
	return
}

func equalWithSmudges(a, b []string) bool {
	reflectionDifferences := 0
	for i := 0; i < min(len(a), len(b)); i++ {
		reflectionDifferences += differences(a[i], b[i])
	}
	return reflectionDifferences == 1
}

func parse() []string {
	return strings.Split(input, "\n\n")
}

func part1() {
	blocks := parse()
	sum := 0

	for _, block := range blocks {
		sum += mirror(rows(block), slices.Equal) * 100
		sum += mirror(columns(block), slices.Equal)
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2() {
	blocks := parse()
	sum := 0
	for _, block := range blocks {
		sum += mirror(rows(block), equalWithSmudges) * 100
		sum += mirror(columns(block), equalWithSmudges)
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
