package main

import (
	"fmt"
	"unicode"
)

func main() {
	lines := readFileLines("input/day1.txt")
	part1Sum := computeCalibrationValuesSumPart1(lines)
	part2Sum := computeCalibrationValuesSumPart2(lines)
	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sum)
}

func computeCalibrationValuesSumPart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		var firstDigit int
		var secondDigit int
		for _, r := range line {
			if unicode.IsDigit(r) {
				firstDigit = int(r - '0')
				break
			}
		}
		for i := range line {
			r := rune(line[len(line)-1-i])
			if unicode.IsDigit(r) {
				secondDigit = int(r - '0')
				break
			}
		}

		calibrationValue := (firstDigit * 10) + secondDigit
		sum += calibrationValue
	}
	return sum
}

func computeCalibrationValuesSumPart2(lines []string) int {
	digitWords := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	sum := 0
	for _, line := range lines {
		var firstDigit int
		var secondDigit int

	forwardLoop:
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				firstDigit = int(line[i] - '0')
				break forwardLoop
			}

			// Not the most efficient but does the job
			// Could try using a suffix tree or something
			for digit, word := range digitWords {
				if len(word) <= len(line)-i && line[i:i+len(word)] == word {
					firstDigit = digit
					break forwardLoop
				}
			}
		}

	backwardLoop:
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				secondDigit = int(line[i] - '0')
				break backwardLoop
			}

			for digit, word := range digitWords {
				if len(word) <= len(line)-i && line[i:i+len(word)] == word {
					secondDigit = digit
					break backwardLoop
				}
			}
		}

		calibrationValue := (firstDigit * 10) + secondDigit
		sum += calibrationValue
	}
	return sum
}
