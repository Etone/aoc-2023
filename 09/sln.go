package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type History = []int

func calculateNext(hist []int) (next int, prev int) {
	values := hist

	level := buildDifferences(values)

	lastLevel := true
	for _, v := range level {
		if v != 0 {
			lastLevel = false
			break
		}
	}

	if lastLevel {
		next = hist[len(hist)-1]
		prev = hist[0]
	} else {
		n, p := calculateNext(level)
		next = hist[len(hist)-1] + n
		prev = hist[0] - p
	}

	return
}

//go:embed input.txt
var input string

func parseInput() (histories []History) {
	historyLines := strings.Split(input, "\n")
	for _, line := range historyLines {
		histories = append(histories, parseHistory(line))
	}
	return
}

func parseHistory(line string) (hist History) {
	values := strings.Split(line, " ")
	for _, v := range values {
		conv, _ := strconv.Atoi(v)
		hist = append(hist, conv)
	}
	return
}

func buildDifferences(values []int) (level []int) {
	for i := 0; i < len(values)-1; i++ {
		a := values[i]
		b := values[i+1]

		level = append(level, b-a)
	}
	return
}

func part1And2() {
	histories := parseInput()
	sumFuture := 0
	sumPast := 0
	for _, hist := range histories {
		predictFuture, predictPast := calculateNext(hist)
		sumFuture += predictFuture
		sumPast += predictPast
	}
	fmt.Printf("Part 1: %d\n", sumFuture)
	fmt.Printf("Part 2: %d\n", sumPast)
}

func main() {
	part1And2()
}
