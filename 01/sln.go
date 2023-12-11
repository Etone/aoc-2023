package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	input, err := os.Open("01/input_01.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	calibaration_value := 0

	for scanner.Scan() {
		number := handleInputLine(scanner.Text(), false)
		calibaration_value += number
	}

	fmt.Printf("Calibration Value Part 1: %d \n", calibaration_value)

	input.Close()
}

func part2() {
	input, err := os.Open("01/input_02.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	calibaration_value := 0

	for scanner.Scan() {
		number := handleInputLine(scanner.Text(), true)
		calibaration_value += number
	}

	fmt.Printf("Calibration Value Part 2: %d \n", calibaration_value)

	input.Close()
}

func handleInputLine(s string, part_2 bool) int {
	text := s
	if part_2 {
		text = prepareString(s)
	}

	first_digit := findFirstDigit(text)
	last_digit := findFirstDigit(reverse(text))

	return first_digit*10 + last_digit

}

func prepareString(s string) string {
	replaces := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "4",
		"five":  "5e",
		"six":   "6",
		"seven": "7n",
		"eight": "e8t",
		"nine":  "9e",
	}

	for i, _ := range replaces {
		s = strings.Replace(s, i, replaces[i], -1)
	}

	return s
}
func findFirstDigit(s string) (digit int) {

	smallest_index := len(s)

	for i := 0; i < 10; i++ {
		index := strings.Index(s, strconv.Itoa(i))

		if index < 0 || index > smallest_index {
			continue
		}
		smallest_index = index
	}

	digit = int(s[smallest_index] - '0')
	return
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
