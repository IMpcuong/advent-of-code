package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	LOSE int = 0
	DRAW int = 3
	WIN  int = 6
)

var (
	AllyAction  []string       = []string{"X", "Y", "Z"}
	EnemyAction []string       = []string{"A", "B", "C"}
	Outcome     map[string]int = map[string]int{
		"AX": DRAW, "AY": WIN, "AZ": LOSE,
		"BX": LOSE, "BY": DRAW, "BZ": WIN,
		"CX": WIN, "CY": LOSE, "CZ": DRAW,
	}
)

func segregateByLine(path string) [][]string {
	roundsAction, _ := os.Open(path)
	dataScanner := bufio.NewScanner(roundsAction)
	dataScanner.Split(bufio.ScanLines)

	var rawLines []string
	for dataScanner.Scan() {
		rawLines = append(rawLines, dataScanner.Text())
	}

	var formattedArr [][]string
	for _, line := range rawLines {
		// NOTE(string): Looping through a string return a list of runes.
		// NOTE: Quoting run variable string.
		// char := strconv.QuoteRune(r)
		formattedArr = append(formattedArr, strings.Split(line, " "))
	}

	return formattedArr
}

func getPointByShape(our string, shapes []string) int {
	for idx, shape := range shapes {
		if shape == our {
			return idx + 1
		}
	}
	return -1
}

// Rule1 := { Enemy := [A, B, C]; Ally := [X, Y, Z] | A=X < B=Y < C=Z }
// Rule2 := map[string]int { "Lose": 1, "Draw": 3, "Won": 6 }
//
// Idea: A method to compute the final outcome from each round/line.
// + Break the input file into an array of multiple vectors, which is an array of 2 particles (integer type).
// + Comparison function: The most easiest way to deal with 2 discriminated arrays (where the underlying value is the same but only their masks were distinguished) is comparing the index from each particle --> Wrong!.

func solvingDay2(path string) int {
	var totalScore int

	roundsData := segregateByLine(path)
	for _, round := range roundsData[:] {
		allyPos := getPointByShape(round[1], AllyAction)
		totalScore += allyPos + Outcome[round[0]+round[1]]
	}
	return totalScore
}
