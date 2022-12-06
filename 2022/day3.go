package main

import (
	"bufio"
	_ "embed"
	"os"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"
)

//go:embed day3.txt
var inputData string

// EmbedData is a function used to test `embed` package only.
func EmbedData() string {
	inputData = strings.TrimRight(inputData, "\n")
	if len(inputData) == 0 {
		panic("Error: Empty *.txt input file!")
	}
	return inputData
}

const ALPHABET string = "abcdefghijklmnopqrstuvwxyz"

// NOTE: rucksack := { [leftCompartment, rightCompartment] | len(leftCompartment) = len(rightCompartment) }

func segregateLineByLength(path string) [][]string {
	rucksackItems, _ := os.Open(path)
	dataScanner := bufio.NewScanner(rucksackItems)
	dataScanner.Split(bufio.ScanLines)

	var rawLines []string
	for dataScanner.Scan() {
		rawLines = append(rawLines, dataScanner.Text())
	}

	var rucksackList [][]string
	for _, line := range rawLines {
		halfLine := len(line) / 2
		compartments := []string{line[:halfLine], line[halfLine:]}
		rucksackList = append(rucksackList, compartments)
	}
	return rucksackList
}

func detectIdenticalChar(left, right string) string {
	// HACK(day3): Both strings have an identical length.
	// HACK(day3): Only return first identical character because 2 compartments only contain one.
	for _, char := range left {
		if strings.ContainsRune(right, char) {
			return string(char)
		}
	}

	return *new(string)
}

func segregateLineByGroup(input string) [][]string {
	var rucksackItems []string
	if runtime.GOOS == "windows" {
		rucksackItems = strings.Split(input, "\r\n")
	} else {
		rucksackItems = strings.Split(input, "\n")
	}
	if len(rucksackItems)%3 != 0 {
		return make([][]string, 0)
	}

	var rucksackList [][]string
	for i := 0; i+2 < len(rucksackItems); i += 3 {
		first := rucksackItems[i]
		second := rucksackItems[i+1]
		third := rucksackItems[i+2]
		group := []string{first, second, third}
		rucksackList = append(rucksackList, group)
	}
	return rucksackList
}

// Exp: Looking through 3 segments from each group to retrieve the most recent badge.
//
// ```txt
// map[68:init 72:init 74:init 80:init 81:init 82:init 87:init 102:init 103:init 104:init 106:init 109:init 110:init 113:init 115:init 118:init]
// map[68:init 72:init 74:matched 80:init 81:matched 82:init 87:init 102:init 103:init 104:init 106:init 109:matched 110:matched 113:matched 115:matched 118:init]
// 115
// ```

// From: https://github.com/BaptisteLalanne/AdventOfCode/blob/d24e5fac929f91f9340041cf41c4889b0c5cc875/2022/day3/main.go#L78

func detectGroupIdentical(first, second, third string) string {
	var charMap = make(map[rune]string)
	for _, char := range first {
		charMap[char] = "init"
	}

	for _, char := range second {
		if _, found := charMap[char]; found {
			charMap[char] = "matched"
		}
	}

	var commonChar string
	for _, char := range third {
		if status := charMap[char]; status == "matched" {
			// NOTE: Stopping when the first match appeared.
			commonChar = string(char)
			break
		}
	}
	return commonChar
}

// End From: https://github.com/BaptisteLalanne/AdventOfCode/blob/d24e5fac929f91f9340041cf41c4889b0c5cc875/2022/day3/main.go#L92

func calculateTypePriority(identicalChar string) int {
	var priorityByType int

	if identicalChar == "" {
		return priorityByType
	}

	for i := range ALPHABET {
		charAsRune, _ := utf8.DecodeLastRuneInString(identicalChar)
		if unicode.IsLower(charAsRune) && identicalChar == string(ALPHABET[i]) {
			priorityByType = i + 1
		}

		if unicode.IsUpper(charAsRune) && strings.ToLower(identicalChar) == string(ALPHABET[i]) {
			priorityByType = i + 1 + len(ALPHABET)
		}
	}

	return priorityByType
}

func solvingDay3(path string) (int, int) {
	// Part1:
	var prioritiesSumP1 int
	rucksackListP1 := segregateLineByLength(path)
	for _, compartments := range rucksackListP1[:] {
		identicalChar := detectIdenticalChar(compartments[0], compartments[1])
		prioritiesSumP1 += calculateTypePriority(identicalChar)
	}

	// Part2:
	var prioritiesSumP2 int
	rucksackListP2 := segregateLineByGroup(inputData)
	for _, group := range rucksackListP2[:] {
		identicalCharInGroup := detectGroupIdentical(group[0], group[1], group[2])
		prioritiesSumP2 += calculateTypePriority(identicalCharInGroup)
	}

	return prioritiesSumP1, prioritiesSumP2
}
