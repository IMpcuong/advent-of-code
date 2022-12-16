package main

import (
	_ "embed"
	"runtime"
	"strconv"
	"strings"
)

//go:embed day7.txt
var day7Data string

// NOTE: This solution was coming from `https://github.com/alokmenghrajani/adventofcode2022/blob/main/day07/day07.go`.
// NOTE: I will do another version by myself later; currently, my headache is terrible.

func parseInputToTree(input string) *Dir {
	var lines []string
	if runtime.GOOS == "windows" {
		lines = strings.Split(input, "\r\n")
	} else {
		lines = strings.Split(input, "\n")
	}

	var rootDir = &Dir{
		Name:    "/",
		Files:   map[string]int{},
		SubDirs: map[string]*Dir{},
		Parent:  nil,
	}

	var curLevelParent = rootDir
	for _, line := range lines {
		commandParts := strings.Split(line, " ")
		if commandParts[0] == "$" {
			if commandParts[1] == "cd" {
				if commandParts[2] == "/" {
					curLevelParent = rootDir
				} else if commandParts[2] == ".." {
					curLevelParent = curLevelParent.Parent
				} else {
					subDirName := commandParts[2]
					curLevelParent = curLevelParent.AppendSubDir(subDirName)
				}
			} else if commandParts[1] == "ls" {
				continue
			} else {
				panic("Unknown Linux command!")
			}
		} else if commandParts[0] == "dir" {
			subDirName := commandParts[1]
			curLevelParent.AppendSubDir(subDirName)
		} else {
			fileSize, _ := strconv.Atoi(commandParts[0])
			fileName := commandParts[1]
			curLevelParent.AppendFile(fileSize, fileName)
		}
	}

	return rootDir
}

type Dir struct {
	Name    string
	Files   map[string]int
	SubDirs map[string]*Dir
	Parent  *Dir
	CurSize int
}

func (dir *Dir) AppendSubDir(dirName string) *Dir {
	newSubDir, exist := dir.SubDirs[dirName]
	if !exist {
		newSubDir = &Dir{
			Name:    dirName,
			Files:   map[string]int{},
			SubDirs: map[string]*Dir{},
			Parent:  dir,
		}
		dir.SubDirs[dirName] = newSubDir
	}
	return newSubDir
}

func (dir *Dir) AppendFile(inputFileSize int, fileName string) {
	newFileSize, exist := dir.Files[fileName]
	if !exist {
		dir.Files[fileName] = inputFileSize
	} else {
		if newFileSize != inputFileSize {
			panic("Input file size was mismatched!")
		}
	}
}

func (dir *Dir) CalculateRecursiveP1(sum *int) int {
	if dir.CurSize != 0 {
		panic("Cannot compute root directory's size!")
	}

	for _, fileSize := range dir.Files {
		dir.CurSize += fileSize
	}

	for _, subDir := range dir.SubDirs {
		dir.CurSize += subDir.CalculateRecursiveP1(sum)
	}

	if dir.CurSize < 100e3 {
		*sum += dir.CurSize
	}

	return dir.CurSize
}

func solutionForP1(input string) int {
	rootDir := parseInputToTree(input)
	sum := 0
	rootDir.CalculateRecursiveP1(&sum)
	return sum
}
