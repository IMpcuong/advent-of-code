package main

import (
	_ "embed"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

//go:embed day4.txt
var day4Data string

func manipulateInput(input string) [][]string {
	var overlapPairs []string
	if runtime.GOOS == "windows" {
		overlapPairs = strings.Split(input, "\r\n")
	} else {
		overlapPairs = strings.Split(input, "\n")
	}

	var overlapSections [][]string
	for _, pair := range overlapPairs {
		assignments := strings.Split(pair, ",")
		firstPair := strings.Split(assignments[0], "-")
		secondPair := strings.Split(assignments[1], "-")
		section := append(firstPair, secondPair...)
		overlapSections = append(overlapSections, section)
	}

	return overlapSections
}

func solvingPart1Day4(sections [][]string) (int, int) {
	var overlappedP1 int
	var overlappedP2 int = len(sections)
	for _, section := range sections {
		if len(section) != 4 {
			fmt.Println("Input data length mismatched!")
			break
		}

		firstOne, _ := strconv.Atoi(section[0])
		firstTwo, _ := strconv.Atoi(section[1])
		secondOne, _ := strconv.Atoi(section[2])
		secondTwo, _ := strconv.Atoi(section[3])

		// Part1+2:
		if firstTwo < secondOne || secondTwo < firstOne {
			overlappedP2--
			continue
		}
		if firstTwo < secondTwo && firstOne < secondOne {
			continue
		}
		if firstTwo > secondTwo && firstOne > secondOne {
			continue
		}
		overlappedP1++
	}

	return overlappedP1, overlappedP2
}
