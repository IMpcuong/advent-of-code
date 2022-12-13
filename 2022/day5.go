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
	MoveAmount int
	FromStack  int
	ToStack    int
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
	for i := 0; i < len(mapColStack); i++ {
		clonedMap[i+1] = mapColStack[4*i+1]
		clonedMap[i+1] = reverseSlice(clonedMap[i+1])
	}
	return clonedMap
}

func mapInstructions(instructionsAsStr []string) []Instruction {
	var insObjs []Instruction
	for _, insStr := range instructionsAsStr {
		numInStr := regexp.MustCompile(`\d+`).FindAllString(insStr, -1)

		move, _ := strconv.ParseInt(numInStr[0], 10, 16)
		from, _ := strconv.ParseInt(numInStr[1], 10, 16)
		to, _ := strconv.ParseInt(numInStr[2], 10, 16)
		ins := Instruction{
			MoveAmount: int(move),
			FromStack:  int(from),
			ToStack:    int(to),
		}
		insObjs = append(insObjs, ins)
	}

	return insObjs
}

func reverseSlice[Type string | int](slices []Type) []Type {
	for i, j := 0, len(slices)-1; i < j; i, j = i+1, j-1 {
		slices[i], slices[j] = slices[j], slices[i]
	}
	return slices
}

func solvingDay5(stacks map[int][]string, instructions []Instruction) string {
	var topCrates []string

	for _, ins := range instructions[:] {
		amount := ins.MoveAmount
		fromStack := ins.FromStack
		toStack := ins.ToStack

		stackHeight := len(stacks[fromStack])
		copyFromIdx := stackHeight - amount
		stacks[toStack] = append(stacks[toStack], reverseSlice(stacks[fromStack][copyFromIdx:stackHeight])...)
		stacks[fromStack] = stacks[fromStack][:copyFromIdx]
	}

	for idx := 1; idx <= len(stacks); idx++ {
		stackHeight := len(stacks[idx])
		topCrates = append(topCrates, stacks[idx][stackHeight-1])
	}

	return strings.Join(topCrates, "")
}
