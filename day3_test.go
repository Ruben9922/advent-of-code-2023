package main

import (
	"slices"
	"strings"
	"testing"
)

var lines = strings.Split(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n")

func TestParse(t *testing.T) {
	actualNumbers, actualSymbols := parse(lines)
	expectedNumbers := []number{
		{
			value:    467,
			position: vector2d{0, 0},
			length:   3,
		},
		{
			value:    114,
			position: vector2d{5, 0},
			length:   3,
		},
		{
			value:    35,
			position: vector2d{2, 2},
			length:   2,
		},
		{
			value:    633,
			position: vector2d{6, 2},
			length:   3,
		},
		{
			value:    617,
			position: vector2d{0, 4},
			length:   3,
		},
		{
			value:    58,
			position: vector2d{7, 5},
			length:   2,
		},
		{
			value:    592,
			position: vector2d{2, 6},
			length:   3,
		},
		{
			value:    755,
			position: vector2d{6, 7},
			length:   3,
		},
		{
			value:    664,
			position: vector2d{1, 9},
			length:   3,
		},
		{
			value:    598,
			position: vector2d{5, 9},
			length:   3,
		},
	}
	expectedSymbols := []symbol{
		{vector2d{3, 1}},
		{vector2d{6, 3}},
		{vector2d{3, 4}},
		{vector2d{5, 5}},
		{vector2d{3, 8}},
		{vector2d{5, 8}},
	}

	if !slices.Equal(actualNumbers, expectedNumbers) {
		t.Errorf("parse(...) numbers = %v; want %v", actualNumbers, expectedNumbers)
	}

	if !slices.Equal(actualSymbols, expectedSymbols) {
		t.Errorf("parse(...) symbols = %v; want %v", actualSymbols, expectedSymbols)
	}
}

func TestGetNumbersAdjacentToSymbols(t *testing.T) {
	numbers := []number{
		{
			value:    467,
			position: vector2d{0, 0},
			length:   3,
		},
		{
			value:    114,
			position: vector2d{5, 0},
			length:   3,
		},
		{
			value:    35,
			position: vector2d{2, 2},
			length:   2,
		},
		{
			value:    633,
			position: vector2d{6, 2},
			length:   3,
		},
		{
			value:    617,
			position: vector2d{0, 4},
			length:   3,
		},
		{
			value:    58,
			position: vector2d{7, 5},
			length:   2,
		},
		{
			value:    592,
			position: vector2d{2, 6},
			length:   3,
		},
		{
			value:    755,
			position: vector2d{6, 7},
			length:   3,
		},
		{
			value:    664,
			position: vector2d{1, 9},
			length:   3,
		},
		{
			value:    598,
			position: vector2d{5, 9},
			length:   3,
		},
	}
	symbols := []symbol{
		{vector2d{3, 1}},
		{vector2d{6, 3}},
		{vector2d{3, 4}},
		{vector2d{5, 5}},
		{vector2d{3, 8}},
		{vector2d{5, 8}},
	}
	adjacentNumberValues := getNumbersAdjacentToSymbols(numbers, symbols, lines)

	expectedAdjacentNumberValues := []int{467, 35, 633, 617, 592, 755, 664, 598}

	if !slices.Equal(adjacentNumberValues, expectedAdjacentNumberValues) {
		t.Errorf("getNumbersAdjacentToSymbols(...) = %v; want %v", adjacentNumberValues, expectedAdjacentNumberValues)
	}
}

func TestSum(t *testing.T) {
	values := []int{45, 6, 34, 6, 63, 1908, 3928394, 192929}
	actualSum := sumValues(values)
	expectedSum := 4123385

	if actualSum != expectedSum {
		t.Errorf("sumValues(...) = %d; want %d", actualSum, expectedSum)
	}
}
