package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type instruction int

const (
	left instruction = iota
	right
)

type network map[string][2]string

func main() {
	lines := readFileLines("input/day8.txt")
	instructions, n := parse(lines)
	stepCount := run(instructions, n)
	fmt.Println(stepCount)
}

func parse(lines []string) ([]instruction, network) {
	instructionsString := lines[0]
	networkStrings := lines[2:]

	instructions, err := parseInstructions(instructionsString)
	if err != nil {
		log.Fatal(err)
	}

	n, err := parseNetwork(networkStrings)
	if err != nil {
		log.Fatal(err)
	}

	return instructions, n
}

func parseInstructions(instructionsString string) ([]instruction, error) {
	instructions := make([]instruction, 0, utf8.RuneCountInString(instructionsString))
	for _, i := range instructionsString {
		switch unicode.ToUpper(i) {
		case 'L':
			instructions = append(instructions, left)
		case 'R':
			instructions = append(instructions, right)
		default:
			return instructions, fmt.Errorf("invalid instruction: %c", i)
		}
	}
	return instructions, nil
}

func parseNetwork(networkStrings []string) (network, error) {
	n := make(network)
	for _, s := range networkStrings {
		sourceNode := strings.Split(s, " = ")[0]
		destNodesString := strings.Split(s, " = ")[1]
		destNodesString = strings.TrimPrefix(destNodesString, "(")
		destNodesString = strings.TrimSuffix(destNodesString, ")")
		destNodeStrings := strings.Split(destNodesString, ", ")

		if len(destNodeStrings) < 2 {
			return n, fmt.Errorf("network edge does not contain two destination nodes: \"%s\"", s)
		}

		n[sourceNode] = [2]string{destNodeStrings[0], destNodeStrings[1]}
	}
	return n, nil
}

func run(instructions []instruction, n network) int {
	currentNode := "AAA"
	stepCount := 0
	for iIndex := 0; currentNode != "ZZZ"; iIndex = (iIndex + 1) % len(instructions) {
		i := instructions[iIndex]
		destNodes := n[currentNode]
		if i == left {
			currentNode = destNodes[0]
		} else if i == right {
			currentNode = destNodes[1]
		}
		stepCount++
	}
	return stepCount
}
