package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var input string

type Record struct {
	reading      string
	springGroups []int
}

func parseInput(input string) (records []Record) {
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		records = append(records, toRecord(row))
	}
	return
}

func toRecord(r string) (rec Record) {
	reading, nums, _ := strings.Cut(r, " ")
	rec.reading = reading
	rec.springGroups = toIntSlice(nums)
	return
}

func toIntSlice(s string) (result []int) {
	nums := strings.Split(s, ",")
	for _, v := range nums {
		conv, _ := strconv.Atoi(v)
		result = append(result, conv)
	}
	return
}

func arrangements(record Record) (result int) {

	if len(record.reading) == 0 {
		if len(record.springGroups) == 0 {
			return 1
		}
		return 0
	}

	firstLetter := string(record.reading[0])

	switch firstLetter {
	case "#":
		result = handleSpring(record)
	case ".":
		rec := Record{record.reading[1:], record.springGroups}
		result = arrangements(rec)

	case "?":
		rec := Record{record.reading[1:], record.springGroups}
		result = arrangements(rec) + handleSpring(record)

	}
	return
}

func handleSpring(record Record) int {
	if len(record.springGroups) == 0 {
		return 0
	}

	currentGroupLength := record.springGroups[0]
	if len(record.reading) < currentGroupLength {
		return 0
	}
	potentialGroup := record.reading[0:currentGroupLength]
	if strings.Contains(potentialGroup, ".") {
		return 0
	}
	if len(record.reading) == currentGroupLength {
		if len(record.springGroups) == 1 {
			return 1
		}
		return 0
	}
	if string(record.reading[currentGroupLength]) == "#" {
		return 0
	}
	rec := Record{record.reading[1:], record.springGroups[1:]}
	return arrangements(rec)
}

func part1() {
	records := parseInput(input)
	sumOfArrangements := 0
	for _, record := range records {
		sumOfArrangements += arrangements(record)
	}
	fmt.Printf("Part 1: %d\n", sumOfArrangements)
}

func part2() {
}

func main() {
	part1()
	part2()
}
