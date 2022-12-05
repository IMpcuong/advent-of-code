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
	prioritiesSum := solvingDay3(path3)
	printAnsWithPattern("=", prioritiesSum)
}
