package main

import (
	_ "embed"
	"regexp"
	"runtime"
	"strconv"
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
			instructions = lines[num+1:]
		}
	}

	return supplyStacks, instructions
}

func mapColStack(matrixData string) map[int][]string {
	var mapColStack = make(map[int][]string)

	var regexCloseBracket = regexp.MustCompile(`\]`)
	lines := strings.Split(matrixData, "\n")
	for _, line := range lines {
		matchedPos := regexCloseBracket.FindAllStringIndex(line, -1)
		for _, pair := range matchedPos {
			pos := pair[0]
			col := (pair[0] - 1) % 35
			mapColStack[col] = append(mapColStack[col], string(line[pos-1]))
		}
	}

	var clonedMap = make(map[int][]string)
	for i := 0; i < 9; i++ {
		clonedMap[i+1] = mapColStack[4*i+1]
	}

	return clonedMap
}

func mapInstructions(instructionsAsStr []string) []Instruction {
	var insObjs []Instruction
	for _, insStr := range instructionsAsStr {
		numInStr := regexp.MustCompile(`\d+`).FindAllString(insStr, -1)

		move, _ := strconv.ParseUint(numInStr[0], 10, 16)
		from, _ := strconv.ParseUint(numInStr[1], 10, 16)
		to, _ := strconv.ParseUint(numInStr[2], 10, 16)
		ins := Instruction{
			Move: uint16(move),
			From: uint16(from),
			To:   uint16(to),
		}
		insObjs = append(insObjs, ins)
	}

	return insObjs
}
