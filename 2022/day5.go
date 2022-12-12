package main

import (
	_ "embed"
	"regexp"
	"runtime"
	"strings"
)

//go:embed day5.txt
var day5Data string

type Instruction struct {
	Move uint16
	From uint16
	To   uint16
}

func partitionData(input string) (supplyStacks string, instructions []string) {
	var lines []string
	if runtime.GOOS == "windows" {
		lines = strings.Split(input, "\r\n")
	} else {
		lines = strings.Split(input, "\n")
	}
	for num, line := range lines {
		if strings.TrimSpace(line) == "" {
			supplyStacks = strings.Join(lines[:num-1], "\n")
			instructions = lines[num:]
		}
	}

	return supplyStacks, instructions
}

func mapColStack(matrixData string) map[int][]string {
	var mapColStack = make(map[int][]string)

	var regexBrackets = regexp.MustCompile(`\]`)
	lines := strings.Split(matrixData, "\n")
	for _, line := range lines {
		matchedPos := regexBrackets.FindAllStringIndex(line, -1)
		for _, pair := range matchedPos {
			pos := pair[0]
			col := (pos%35 + 1) / 3
			mapColStack[col] = append(mapColStack[col], string(line[pos-1]))
		}
	}

	return mapColStack
}
