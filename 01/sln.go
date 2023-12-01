package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	readFile, err := os.Open("01/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	calibaration_value := 0

	for scanner.Scan() {
		number := handleInputLine(scanner.Text())
		calibaration_value += number
	}

	fmt.Printf("Calibration Value: %d \n", calibaration_value)

	readFile.Close()
}

func handleInputLine(s string) int {
	first_digit := findFirstDigit(s)
	last_digit := findLastDigit(s)

	println(s, first_digit, last_digit)


	return first_digit * 10 + last_digit

}

func findFirstDigit(s string) int {
	for _, char := range s {
		if(unicode.IsDigit(char)) {
			return int(char - '0')
		}
	}
	return 0
}

func findLastDigit(s string) int {
	return findFirstDigit(reverse(s))
}

func reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
  }
