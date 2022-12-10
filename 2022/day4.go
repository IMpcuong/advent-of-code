package main

import (
	_ "embed"
	"runtime"
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
		assignments := strings.Split(pair, "-")
		first := assignments[0]
		second := strings.Split(assignments[1], ",")[0]
		third := strings.Split(assignments[1], ",")[1]
		fourth := assignments[2]
		assignments = []string{first, second, third, fourth}
		overlapSections = append(overlapSections, assignments)
	}

	return overlapSections
}
