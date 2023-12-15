package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func hash(s string) uint8 {
	value := 0
	for _, rune := range s {
		value += int(rune)
		value *= 17
	}
	return uint8(value)
}

func part1() {
	initializationSequence := strings.Split(input, ",")
	sumOfHashes := 0
	for _, step := range initializationSequence {
		hash := hash(step)
		sumOfHashes += int(hash)
	}
	fmt.Printf("Part 1: %d\n", sumOfHashes)
}

func part2() {
}

func main() {
	part1()
	part2()
}
