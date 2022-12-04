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
	arrAns, max := solvingDay1("day1.txt")
	printAnsWithPattern("=", fmt.Sprintf("%v\t%d", arrAns, max))

	totalPoint := solvingDay2("day2.txt")
	printAnsWithPattern("=", totalPoint)
}
