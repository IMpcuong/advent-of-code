package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const ALPHABET string = "abcdefghijklmnopqrstuvwxyz"

// NOTE: rucksack := { [leftCompartment, rightCompartment] | len(leftCompartment) = len(rigthCompartment) }

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
		compartments := []string{line[:halfLine], line[halfLine+1:]}
		rucksackList = append(rucksackList, compartments)
	}
	return rucksackList
}

func detectIndeticalChar(left, right string) string {
	// HACK(day3): Both strings have an identical length.
	// HACK(day3): Only return first identical charater because 2 compartments only contain one.
	for _, char := range left {
		if strings.ContainsRune(right, char) {
			return string(char)
		}
	}

	return *new(string)
}

func solvingDay3(path string) int {
	var prioritiesSum int

	rucksackList := segregateLineByLength(path)
	for _, compartments := range rucksackList[:] {
		fmt.Println(compartments)
		identicalChar := detectIndeticalChar(compartments[0], compartments[1])
		fmt.Println(identicalChar)

		if identicalChar == "" {
			continue
		}

		var itemTypeVal int
		for i := range ALPHABET {
			charAsRune, _ := utf8.DecodeLastRuneInString(identicalChar)
			if unicode.IsLower(charAsRune) && identicalChar == string(ALPHABET[i]) {
				itemTypeVal = i + 1
				fmt.Println(itemTypeVal)
			}

			if unicode.IsUpper(charAsRune) && strings.ToLower(identicalChar) == string(ALPHABET[i]) {
				itemTypeVal = i + 1 + len(ALPHABET)
				fmt.Println(itemTypeVal)
			}
		}
		prioritiesSum += itemTypeVal
		fmt.Println(prioritiesSum)
	}

	return prioritiesSum
}
