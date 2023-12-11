package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Measurement = []int

//go:embed input.txt
var input string

func parseInput() (measurements []Measurement) {
	historyLines := strings.Split(input, "\n")
	for _, line := range historyLines {
		measurements = append(measurements, parseHistory(line))
	}
	return
}

func parseHistory(line string) (measured Measurement) {
	values := strings.Fields(line)
	for _, v := range values {
		conv, _ := strconv.Atoi(v)
		measured = append(measured, conv)
	}
	return
}

func extrapolate(values []int) (int, int) {
	if allZero(values) {
		return 0, 0
	}

	diffs := buildDifferences(values)
	n, p := extrapolate(diffs)
	return values[len(values)-1] + n, values[0] - p
}

func buildDifferences(values []int) (level []int) {
	for i := 0; i < len(values)-1; i++ {
		level = append(level, values[i+1]-values[i])
	}
	return
}

func allZero(vals []int) (allZero bool) {
	allZero = true
	for _, v := range vals {
		if v != 0 {
			allZero = false
			break
		}
	}
	return
}

func part1And2() {
	histories := parseInput()
	sumFuture := 0
	sumPast := 0
	for _, hist := range histories {
		predictFuture, predictPast := extrapolate(hist)
		sumFuture += predictFuture
		sumPast += predictPast
	}
	fmt.Printf("Part 1: %d\n", sumFuture)
	fmt.Printf("Part 2: %d\n", sumPast)
}

func main() {
	part1And2()
}
