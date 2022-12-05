package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

//go:embed day3.txt
var input string

// EmbedData is a function used to test `embed` package only.
func EmbedData() string {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("Error: Empty *.txt input file!")
	}
	return input
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

func segregateLineByGroup(input string) []string {
	rucksackItems := strings.Split(input, "\n")
	if len(rucksackItems)%3 != 0 {
		return make([]string, 0)
	}

	for i := 0; i+2 < len(rucksackItems); i += 3 {
		first := rucksackItems[i]
		second := rucksackItems[i+1]
		third := rucksackItems[i+2]
		fmt.Println(first, second, third)
	}
	return make([]string, 0)
}

func solvingDay3(path string) int {
	var prioritiesSum int

	rucksackList := segregateLineByLength(path)
	for _, compartments := range rucksackList[:] {
		identicalChar := detectIdenticalChar(compartments[0], compartments[1])
		if identicalChar == "" {
			continue
		}

		var itemTypeVal int
		for i := range ALPHABET {
			charAsRune, _ := utf8.DecodeLastRuneInString(identicalChar)
			if unicode.IsLower(charAsRune) && identicalChar == string(ALPHABET[i]) {
				itemTypeVal = i + 1
			}

			if unicode.IsUpper(charAsRune) && strings.ToLower(identicalChar) == string(ALPHABET[i]) {
				itemTypeVal = i + 1 + len(ALPHABET)
			}
		}
		prioritiesSum += itemTypeVal
	}

	return prioritiesSum
}
