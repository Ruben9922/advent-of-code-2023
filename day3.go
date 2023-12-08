package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type vector2d struct {
	x int
	y int
}

type number struct {
	value    int
	position vector2d
	length   int
}

type symbol struct {
	position vector2d
}

func main() {
	lines := readFileLines("input/day3_1.txt")
	numbers, symbols := parse(lines)
	adjacentNumberValues := getNumbersAdjacentToSymbols(numbers, symbols, lines)
	adjacentNumberValuesSum := sumValues(adjacentNumberValues)
	x := slices.ContainsFunc(numbers, func(n number) bool {
		return n.value == 579
	})
	y := slices.Contains(adjacentNumberValues, 579)
	fmt.Println(adjacentNumberValuesSum, x, y)
}

func parse(lines []string) ([]number, []symbol) {
	numbers := make([]number, 0)
	symbols := make([]symbol, 0)
	for i, line := range lines {
		var currentNumberStringBuilder strings.Builder
		for i2, r := range line {
			if unicode.IsDigit(r) {
				currentNumberStringBuilder.WriteRune(r)
			} else {
				if currentNumberStringBuilder.Len() > 0 {
					currentNumberString := currentNumberStringBuilder.String()
					currentNumber, _ := strconv.Atoi(currentNumberString)
					length := utf8.RuneCountInString(currentNumberString)
					n := number{
						value:    currentNumber,
						position: vector2d{x: i2 - length, y: i},
						length:   length,
					}
					numbers = append(numbers, n)

					currentNumberStringBuilder.Reset()
				}

				if r != '.' {
					symbols = append(symbols, symbol{
						vector2d{x: i2, y: i},
					})
				}
			}
		}
	}
	return numbers, symbols
}

// Could probably improve performance by using maps
func getNumbersAdjacentToSymbols(numbers []number, symbols []symbol, lines []string) []int {
	adjacentNumberValues := make([]int, 0, len(numbers))
	for _, n := range numbers {
		numberNeighbors := getNumberNeighbors(n, lines)

	numberNeighbors:
		for _, neighbor := range numberNeighbors {
			for _, s := range symbols {
				if neighbor.x == s.position.x && neighbor.y == s.position.y {
					adjacentNumberValues = append(adjacentNumberValues, n.value)
					break numberNeighbors
				}
			}
		}
	}
	return adjacentNumberValues
}

func getNumberNeighbors(n number, lines []string) []vector2d {
	neighbors := make([]vector2d, 0, 6+(2*n.length))
	for x := -1; x <= n.length; x++ {
		for y := -1; y <= 1; y++ {
			offset := vector2d{x: x, y: y}
			neighbor := vector2d{
				x: n.position.x + offset.x,
				y: n.position.y + offset.y,
			}
			if !(neighbor.y == n.position.y && neighbor.x >= n.position.x && neighbor.x < n.position.x+n.length) &&
				isPointValid(neighbor, lines) {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

func isPointValid(point vector2d, lines []string) bool {
	return point.y >= 0 && point.y < len(lines) && point.x >= 0 && point.x < len(lines[point.x])
}

func sumValues(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
