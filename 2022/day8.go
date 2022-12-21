package main

import (
	_ "embed"
	"fmt"
	"math"
	"runtime"
	"strings"
)

//go:embed day8.txt
var day8Data string

// Visualization problem:
/*
	\ 0 1 2 3 4 5
	0 x x x x x x
	1 x - - - - x
	2 x - - - - x
	3 x - - - - x
	4 x - - - - x
	5 x x x x x x
*/

type Matrix [][]int

func (m Matrix) Get(x, y int) int {
	return m[x][y]
}

func (m Matrix) GetRange(x, y int) int {
	var scenicRange = make([]int, 0)
	if x == 0 || y == 0 || x == len(m)-1 || y == len(m)-1 {
		return 0
	}

	for _, direction := range Directions {
		curX := x
		curY := y
		for {
			curX = curX + direction[0]
			curY = curY + direction[1]
			if curX < 0 || curY >= len(m) || curX >= len(m) || curY < 0 {
				break
			}
			if m.Get(x, y) <= m.Get(curX, curY) {
				if curX == x {
					scenicRange = append(scenicRange, int(math.Abs(float64(curY-y))))
				} else if curY == y {
					scenicRange = append(scenicRange, int(math.Abs(float64(curX-x))))
				}
				break
			} else {
				if curX < x {
					scenicRange = append(scenicRange, x)
				} else if x > curX {
					scenicRange = append(scenicRange, len(m)-x)
				} else if curY < y {
					scenicRange = append(scenicRange, y)
				} else {
					scenicRange = append(scenicRange, len(m)-1-y)
				}
				break
			}
		}
	}
	// NOTE: The arduous cases were about the height difference between our and grid's edge (our > edge).
	fmt.Println(scenicRange)

	var visibleSize int = 1
	for _, rangeDir := range scenicRange {
		visibleSize *= rangeDir
	}

	return visibleSize
}

func convertToMatrix(input string) Matrix {
	var lines []string
	if runtime.GOOS == "windows" {
		lines = strings.Split(input, "\r\n")
	} else {
		lines = strings.Split(input, "\n")
	}

	var matrixData = make([][]int, 0)
	for _, line := range lines {
		var row = make([]int, 0)
		for _, char := range line {
			charAsNum := int(char - '0')
			row = append(row, charAsNum)
		}
		matrixData = append(matrixData, row)
	}

	return matrixData
}

var Directions = [][2]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func locateVisibleTree(grid Matrix) int {
	var visibleCount int

	for rowId := 1; rowId < len(grid)-1; rowId++ {
		for colId := 1; colId < len(grid[rowId])-1; colId++ {
			for _, direction := range Directions {
				curX := rowId
				curY := colId
				isVisible := true
				for {
					curX = curX + direction[0]
					curY = curY + direction[1]
					if curX < 0 || curY >= len(grid) || curX >= len(grid) || curY < 0 {
						break
					}
					if grid.Get(rowId, colId) <= grid.Get(curX, curY) {
						isVisible = false
						break
					}
				}
				if isVisible {
					visibleCount++
					break
				}
			}
		}
	}

	visibleCount += len(grid)*4 - 4
	return visibleCount
}

func computeVisibleRange(grid Matrix) int {
	var visibleRange int
	for rowId := 0; rowId < len(grid); rowId++ {
		for colId := 0; colId < len(grid[rowId]); colId++ {
			fmt.Println(grid.Get(rowId, colId))
			visibleRange += grid.GetRange(rowId, colId)
		}
	}

	return visibleRange
}
