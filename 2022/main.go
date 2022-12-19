package main

import (
	"fmt"
	"strings"
)

func getPathByDay(day int) string {
	return fmt.Sprintf("day%d.txt", day)
}

func printAnsWithPattern(pattern string, ans interface{}) {
	ansAsAtr := fmt.Sprintf("%s", ans)
	fmt.Printf("%s>\t%v\n", strings.Repeat(pattern, 5), ansAsAtr)
}

func main() {
	// HACK(embed): Embedded file content into variable using `embed` package.
	// EmbedData()

	// Day1:
	path1 := getPathByDay(1)
	maxVal, sumTopThree := solvingDay1(path1)
	printAnsWithPattern("=", fmt.Sprintf("%v\t%d", maxVal, sumTopThree))

	// Day2:
	path2 := getPathByDay(2)
	totalPointP1, totalPointP2 := solvingDay2(path2)
	printAnsWithPattern("=", fmt.Sprintf("%d\t%d", totalPointP1, totalPointP2))

	// Day3:
	path3 := getPathByDay(3)
	prioritiesSumP1, prioritiesSumP2 := solvingDay3(path3)
	printAnsWithPattern("=", fmt.Sprintf("%d\t%d", prioritiesSumP1, prioritiesSumP2))

	// Day4:
	overlapSections := manipulateInput(day4Data)
	overlappedP1, overlappedP2 := solvingDay4(overlapSections)
	printAnsWithPattern("=", fmt.Sprintf("%d\t%d", overlappedP1, overlappedP2))

	// Day5:
	supplyStacks, instructions := partitionData(day5Data)
	colStack := mapColStack(supplyStacks)
	clonedColStack := make(map[int][]string)
	clonedColStack = func(colMap map[int][]string) map[int][]string {
		for k, v := range colMap {
			clonedColStack[k] = v
		}
		return clonedColStack
	}(colStack)
	instructionObjs := mapInstructions(instructions)
	clonedInstructionObjs := make([]Instruction, 0)
	clonedInstructionObjs = append(clonedInstructionObjs, instructionObjs...)
	topCratesP1 := solvingDay5Part1(colStack, instructionObjs)
	topCratesP2 := solvingDay5Part2(clonedColStack, clonedInstructionObjs)
	printAnsWithPattern("=", fmt.Sprintf("%s\t%s", topCratesP1, topCratesP2))

	// Day6:
	firstPosPacketMarker := detectPacketMarker(day6Data)
	firstPosMsgMarkerV1 := detectMsgMarkerV1(day6Data, 14)
	firstPosMsgMarkerV2 := detectMsgMarkerV2(day6Data, 14)
	printAnsWithPattern("=", fmt.Sprintf("%d\t%d\t%d", firstPosPacketMarker, firstPosMsgMarkerV1, firstPosMsgMarkerV2))

	// Day7:
	part1Ans := solutionForP1(day7Data)
	part2Ans := solutionForP2(day7Data)
	printAnsWithPattern("=", fmt.Sprintf("%d\t%d", part1Ans, part2Ans))

	// Day8:
	matrix := convertToMatrix(day8Data)
	// fmt.Println(matrix[:2])
	fmt.Println(locateVisibleTree(matrix))
}
