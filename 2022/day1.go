package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func sumArrOrSlice(intArr []int) int {
	var res = 0
	for _, v := range intArr {
		res += v
	}
	return res
}

func removeZeroVal(intArr []int) []int {
	for idx, v := range intArr {
		if v == 0 {
			intArr = append(intArr[:idx], intArr[idx+1:]...)
			return removeZeroVal(intArr)
		}
	}
	return intArr
}

func solvingDay1() (max int, sumTopThree int) {
	caloriesData, _ := os.Open("day1.txt")
	dataScanner := bufio.NewScanner(caloriesData)
	dataScanner.Split(bufio.ScanLines)

	var rawLines []string
	for dataScanner.Scan() {
		rawLines = append(rawLines, dataScanner.Text())
	}

	linesToNum := func(strArr []string) []int {
		numArr := make([]int, 0, len(strArr))
		for _, str := range strArr {
			num, _ := strconv.Atoi(str)
			numArr = append(numArr, num)
		}
		return numArr
	}(rawLines)

	var arrAns []int
	var sum = 0
	for idx := range linesToNum {
		if linesToNum[idx] != 0 {
			sum += linesToNum[idx]
		} else {
			arrAns = append(arrAns, sum)
			sum = 0
		}
	}
	arrAns = removeZeroVal(arrAns)
	sort.Ints(arrAns[:])
	max = arrAns[len(arrAns)-1]
	sumTopThree = sumArrOrSlice(arrAns[len(arrAns)-3:])

	return max, sumTopThree
}
