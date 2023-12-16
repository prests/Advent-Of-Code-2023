package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

//go:embed input.txt
var efs embed.FS

var digitsAsWordsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var windowSizes = [3]int{3, 4, 5}

func combineDigits(a, b int) int {
	return (a * 10) + b
}

func swapWordsForDigit(input string) string {
	for _, windowSize := range windowSizes {
		if windowSize > len(input) {
			continue
		}

		for i := 0; i+windowSize <= len(input); i++ {
			substring := input[i : i+windowSize]

			digit, exists := digitsAsWordsMap[substring]
			if exists {
				/**
				* Leave the start and end letters of a number written as a word.
				* This helps prevent cases like "eightthree" becoming "8hree"
				* https://www.reddit.com/r/adventofcode/comments/1884fpl/2023_day_1for_those_who_stuck_on_part_2/
				 */
				remainingInput := string(input[i+windowSize-1:])
				input = input[:i+1] + strconv.Itoa(digit) + remainingInput
			}
		}
	}

	return input
}

func getDigits(input string) []int {
	var digits []int

	for _, letter := range input {
		if unicode.IsDigit(letter) {
			digit, err := strconv.Atoi(string(letter))
			if err != nil {
				fmt.Println("Error converting rune to int", err)
				continue
			}

			digits = append(digits, digit)
		}
	}

	return digits
}

func GetCalibrationValue(input string) int {
	fmt.Println(input)
	input = swapWordsForDigit(input)

	digits := getDigits(input)

	fmt.Println(digits)
	fmt.Print("\n")

	if len(digits) == 0 {
		return 0
	}

	if len(digits) == 1 {
		return combineDigits(digits[0], digits[0])
	}

	return combineDigits(digits[0], digits[len(digits)-1])
}

func GetSum(inputs []int) int {
	sum := 0

	for _, val := range inputs {
		sum += val
	}

	return sum
}

func main() {
	file, err := efs.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var calibrationValues []int
	for scanner.Scan() {
		input := scanner.Text()
		calibrationValues = append(calibrationValues, GetCalibrationValue(input))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	sum := GetSum(calibrationValues)
	fmt.Printf("sum: %d\n", sum)
}
