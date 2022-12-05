package main

import (
	"fmt"
	"strings"
)

func printAnsWithPattern(pattern string, ans interface{}) {
	ansAsAtr := fmt.Sprintf("%s", ans)
	fmt.Printf("%s>\t%v\n", strings.Repeat(pattern, 5), ansAsAtr)
}

func main() {
	maxVal, sumTopThree := solvingDay1("day1.txt")
	printAnsWithPattern("=", fmt.Sprintf("%v\t%d", maxVal, sumTopThree))

	totalPointP1, totalPointP2 := solvingDay2("day2.txt")
	printAnsWithPattern("=", fmt.Sprintf("%d\t%d", totalPointP1, totalPointP2))
}
