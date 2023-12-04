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

func numberOfMatches(winning []int, guessed []int) (matches int) {
	for _, guess := range guessed {
		if slices.Contains(winning, guess) {
			matches++
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
		matches := numberOfMatches(win, guessed)
		if matches > 0 {
			sumOfPoints += int(math.Pow(2, float64(matches)-1))
		}
	}
	fmt.Printf("Part 1: %d\n", sumOfPoints)
}

func part2() {
	games := readFileToArray("input_01.txt")

	numberOfGames := map[int]int{}
	numberOfMatchesPerGame := map[int]int{}

	for index := range games {
		numberOfGames[index] = 1
	}

	for index, game := range games {
		win, guessed := parseGameLine(game)
		numberOfMatchesPerGame[index] = numberOfMatches(win, guessed)
	}

	for gameId := 0; gameId < len(games); gameId++ {
		for numberOfTickets := numberOfGames[gameId]; numberOfTickets > 0; numberOfTickets-- {
			for matches := numberOfMatchesPerGame[gameId]; matches > 0; matches-- {
				numberOfGames[gameId+matches]++
			}
		}
	}

	result := 0
	for _, val := range numberOfGames {
		result += val
	}

	fmt.Printf("Part 2: %d\n", result)
}
