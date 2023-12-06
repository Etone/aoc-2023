package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	milliseconds int
	record       int
}

var highscores []race

var highscoreP2 race

func initializeGlobals() {
	highscores = []race{}
}

func parseInputP1() {
	data, _ := os.ReadFile("input.txt")
	removeSpaceRegex := regexp.MustCompile(`[ ]+`)
	cleaned := string(data)
	cleaned = removeSpaceRegex.ReplaceAllString(cleaned, ";")
	lines := strings.Split(cleaned, "\n")
	_, cleanedMillisecondsCSV, _ := strings.Cut(lines[0], ";")
	_, cleanedMillimeterCSV, _ := strings.Cut(lines[1], ";")
	cleanedMs := strings.Split(cleanedMillisecondsCSV, ";")
	cleanedMm := strings.Split(cleanedMillimeterCSV, ";")

	for i := 0; i < len(cleanedMs); i++ {
		ms, _ := strconv.Atoi(cleanedMs[i])
		mm, _ := strconv.Atoi(cleanedMm[i])
		highscores = append(highscores, race{milliseconds: ms, record: mm})
	}
}

func parseInputP2() {
	data, _ := os.ReadFile("input.txt")
	removeSpaceRegex := regexp.MustCompile(`[ ]+`)
	cleaned := string(data)
	cleaned = removeSpaceRegex.ReplaceAllString(cleaned, "")
	lines := strings.Split(cleaned, "\n")
	_, cleanMs, _ := strings.Cut(lines[0], ":")
	_, cleanMm, _ := strings.Cut(lines[1], ":")
	ms, _ := strconv.Atoi(cleanMs)
	mm, _ := strconv.Atoi(cleanMm)
	highscoreP2.milliseconds = ms
	highscoreP2.record = mm
}

func main() {
	initializeGlobals()
	parseInputP1()
	parseInputP2()
	part1()
	part2()
}

func part1() {
	multipliedNumberOfWays := 1
	for _, race := range highscores {
		countToBeat := 0
		for t := 0; t <= race.milliseconds; t++ {
			reach := (race.milliseconds - t) * t
			if reach > race.record {
				countToBeat++
			}
		}
		multipliedNumberOfWays *= countToBeat
	}
	fmt.Printf("Part 1: %d\n", multipliedNumberOfWays)
}

func part2() {
	countToBeat := 0
	for t := 0; t <= highscoreP2.milliseconds; t++ {
		reach := (highscoreP2.milliseconds - t) * t
		if reach > highscoreP2.record {
			countToBeat++
		}
	}
	fmt.Printf("Part 2: %d\n", countToBeat)
}
