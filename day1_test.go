package main

import (
	"strings"
	"testing"
)

func TestComputeCalibrationValuesSumPart1(t *testing.T) {
	input := strings.Split(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`, "\n")
	actualOutput := computeCalibrationValuesSumPart1(input)
	expectedOutput := 142

	if actualOutput != expectedOutput {
		t.Errorf("computeCalibrationValuesSumPart1(...) = %d; want %d", actualOutput, expectedOutput)
	}
}

func TestComputeCalibrationValuesSumPart2(t *testing.T) {
	input := strings.Split(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, "\n")
	actualOutput := computeCalibrationValuesSumPart2(input)
	expectedOutput := 281

	if actualOutput != expectedOutput {
		t.Errorf("computeCalibrationValuesSumPart2(...) = %d; want %d", actualOutput, expectedOutput)
	}
}
