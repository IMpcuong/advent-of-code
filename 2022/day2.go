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
	AllyShape  []string       = []string{"X", "Y", "Z"} // Rock, Paper, Scissors
	EnemyShape []string       = []string{"A", "B", "C"} // Rock, Paper, Scissors
	OutcomeP1  map[string]int = map[string]int{
		"AX": DRAW, "AY": WIN, "AZ": LOSE,
		"BX": LOSE, "BY": DRAW, "BZ": WIN,
		"CX": WIN, "CY": LOSE, "CZ": DRAW,
	}
)

// NOTE:
// 1. `String`` is an immutable byte sequence.
// 2. `Byte slice` is a mutable byte sequence.
// 3. `Rune slice` is re-grouping of byte slice so that each index is a character.

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

func getShapeIdx(our string, shapes []string) int {
	for idx, shape := range shapes {
		if shape == our {
			return idx // pos := {0, 1, 2}
		}
	}
	return -1
}

// Part1:
// Rule1 := { Enemy := [A, B, C]; Ally := [X, Y, Z] | A=X < B=Y < C=Z }
// Rule2 := map[string]int { "Lose": 1, "Draw": 3, "Won": 6 }
//
// Idea: A method to compute the final outcome from each round/line.
// + Break the input file into an array of multiple vectors, which is an array of 2 particles (integer type).
// + Comparison function: The most easiest way to deal with 2 discriminated arrays (where the underlying value is the same but only their masks were distinguished) is comparing the index from each particle --> Wrong!.

// Part2:
// Rule3 := map[string]int { "X": Lose, "Y": Draw, "Z": Win }
// Exp: [A, Z] := Y (Enemy: 'Rock', Want 'Win' -> Ally: 'Scissors') -> getRoundPoint(A, Y)

var (
	Shapes map[string]int = map[string]int{
		"X": 0, // Rock
		"Y": 1, // Paper
		"Z": 2, // Scissors
	}
	OutcomeP2 = [3][3]string{
		/*L    D    W*/
		{"Z", "X", "Y"}, // Rock
		{"X", "Y", "Z"}, // Paper
		{"Y", "Z", "X"}, // Scissors
	}
)

func solvingDay2(path string) (int, int) {
	roundsData := segregateByLine(path)

	var totalScorePart1 int
	var totalScorePart2 int
	for _, round := range roundsData[:] {
		// Part1:
		allyShapePoint := getShapeIdx(round[1], AllyShape) + 1
		totalScorePart1 += allyShapePoint + OutcomeP1[round[0]+round[1]]

		// Part2:
		enemyShape := getShapeIdx(round[0], EnemyShape)
		outcomeState := getShapeIdx(round[1], AllyShape)
		allyShape := OutcomeP2[enemyShape][outcomeState]
		allyShapePointP2 := Shapes[allyShape] + 1
		totalScorePart2 += allyShapePointP2 + outcomeState*3
	}

	return totalScorePart1, totalScorePart2
}
