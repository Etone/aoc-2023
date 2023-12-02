package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id   int
	sets []set
}

type set struct {
	red   int
	green int
	blue  int
}

func main() {
	part1()
	part2()
}

func readFileToArray(path string) (lines []string) {
	input, _ := os.Open("02/" + path)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func parseLineToGame(line string) (game game) {
	gameText, setsText, _ := strings.Cut(line, ":")
	gameId, _ := strconv.Atoi(strings.TrimPrefix(gameText, "Game "))

	setsLines := strings.Split(setsText, ";")
	for _, setLine := range setsLines {

		set := set{}

		colorTexts := strings.Split(setLine, ",")
		for _, colorText := range colorTexts {
			colorText = strings.TrimSpace(colorText)
			count, color, _ := strings.Cut(colorText, " ")

			switch color {
			case "red":
				set.red, _ = strconv.Atoi(count)
			case "green":
				set.green, _ = strconv.Atoi(count)
			case "blue":
				set.blue, _ = strconv.Atoi(count)
			}
		}
		game.sets = append(game.sets, set)
	}

	game.id = gameId
	return
}

func isGamePlayable(game game, maxAllowed set) bool {
	for _, set := range game.sets {
		if set.blue > maxAllowed.blue || set.red > maxAllowed.red || set.green > maxAllowed.green {
			return false
		}
	}
	return true
}

func parseLinesToGames(lines []string) (games []game) {
	for _, line := range lines {
		games = append(games, parseLineToGame(line))
	}
	return
}

func powerOfSet(set set) int {
	return set.blue * set.green * set.red
}

func minimunSetNeeded(game game) (minimumSet set) {
	for _, set := range game.sets {
		minimumSet.blue = max(minimumSet.blue, set.blue)
		minimumSet.red = max(minimumSet.red, set.red)
		minimumSet.green = max(minimumSet.green, set.green)
	}
	return
}

func part1() {
	lines := readFileToArray("input_01.txt")

	games := parseLinesToGames(lines)

	sumOfPlayableIds := 0
	for _, game := range games {
		if isGamePlayable(game, set{red: 12, green: 13, blue: 14}) {
			sumOfPlayableIds += game.id
		}
	}
	fmt.Printf("Part 1: %d\n", sumOfPlayableIds)

}

func part2() {
	lines := readFileToArray("input_01.txt")
	games := parseLinesToGames(lines)
	sumOfPowersOfMimimumSets := 0
	for _, game := range games {
		minimum := minimunSetNeeded(game)
		sumOfPowersOfMimimumSets += powerOfSet(minimum)
	}
	fmt.Printf("Part 2: %d\n", sumOfPowersOfMimimumSets)
}
