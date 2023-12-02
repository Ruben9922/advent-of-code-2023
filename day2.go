package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := readFileLines("input/day2.txt")

	gameIdSum := 0
	minPossibleSetPowerSum := 0
	for _, line := range lines {
		gameId, err := parseGameId(line)
		if err != nil {
			log.Fatalf("Failed parsing game ID: %v", err)
		}

		sets, err := parseSets(line)
		if err != nil {
			log.Fatalf("Failed parsing sets: %v", err)
		}

		valid := areSetsValid(sets)
		if valid {
			gameIdSum += gameId
		}

		minPossibleSet := getMinPossibleSet(sets)
		minPossibleSetPowerSum += computeSetPower(minPossibleSet)
	}

	fmt.Println(gameIdSum)
	fmt.Println(minPossibleSetPowerSum)
}

func parseGameId(line string) (int, error) {
	return strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
}

func parseSets(line string) ([]map[string]int, error) {
	setsString := strings.Split(line, ": ")[1]
	setStrings := strings.Split(setsString, "; ")
	sets := make([]map[string]int, 0, len(setStrings))
	for _, setString := range setStrings {
		cubeCountStrings := strings.Split(setString, ", ")
		set := make(map[string]int)
		for _, cubeCountString := range cubeCountStrings {
			countString := strings.Split(cubeCountString, " ")[0]
			count, err := strconv.Atoi(countString)
			if err != nil {
				return sets, err
			}
			colorString := strings.Split(cubeCountString, " ")[1]

			set[colorString] = count
		}
		sets = append(sets, set)
	}
	return sets, nil
}

func areSetsValid(sets []map[string]int) bool {
	maxCubesByColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, set := range sets {
		for color, count := range set {
			if maxCount, colorPresent := maxCubesByColor[color]; colorPresent && count > maxCount {
				return false
			}
		}
	}
	return true
}

func getMinPossibleSet(sets []map[string]int) map[string]int {
	minPossibleCounts := make(map[string]int)

	for _, set := range sets {
		for color, count := range set {
			if minPossibleCount, present := minPossibleCounts[color]; !present || count > minPossibleCount {
				minPossibleCounts[color] = count
			}
		}
	}

	return minPossibleCounts
}

func computeSetPower(set map[string]int) int {
	return set["red"] * set["green"] * set["blue"]
}
