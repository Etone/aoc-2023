package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFileToArray(path string) (lines []string) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func parseGameLine(line string) (winningNumbers []int, guessedNumbers []int) {
	numbers := strings.SplitAfter(line, ":")[1]
	winning, guessed, _ := strings.Cut(numbers, " | ")

	winning = strings.TrimSpace(winning)
	guessed = strings.TrimSpace(guessed)

	winningNumbers = parseNumbersString(winning)
	guessedNumbers = parseNumbersString(guessed)
	return
}

func parseNumbersString(guessed string) (numbers []int) {
	numbersAsText := strings.Split(guessed, " ")
	for _, text := range numbersAsText {
		convert, error := strconv.Atoi(text)
		if error == nil {
			numbers = append(numbers, convert)
		}
	}
	return
}

func main() {
	part1()
	part2()
}

func part1() {
	games := readFileToArray("input_01.txt")

	sumOfPoints := 0

	for _, game := range games {
		win, guessed := parseGameLine(game)
		matches := 0
		for _, guess := range guessed {
			if slices.Contains(win, guess) {
				matches++
			}
		}
		if matches > 0 {
			sumOfPoints += int(math.Pow(2, float64(matches)-1))
		}
	}
	fmt.Printf("Part 1: %d\n", sumOfPoints)
}

func part2() {
}
