package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day6.txt
var day6Data string

// NOTE: This was a completely bogus solution!
func detectPacketMarker(data string) int {
	for i := 0; i < len(data)-3; i++ {
		if !strings.Contains(data[i+1:i+4], string(data[i])) && !strings.Contains(data[i:i+3], string(data[i+3])) && data[i+1] != data[i+2] {
			return i + 4
		}
	}

	return -1
}

func recursiveSearch(char byte, seq string, begin, end int) (pos int) {
	if end < 1 {
		return -1
	}
	if seq[begin] == char {
		return begin + 1
	}
	if seq[end] == char {
		return end + 1
	}
	return recursiveSearch(char, seq, begin+1, end-1)
}

func detectMsgMarker(data string) int {
	var maxUniqueLen int = 0
	// var listMatchedPos []int
	for i := 0; i < len(data)-14; i++ {
		needed := data[i]
		firstMatchedPos := recursiveSearch(needed, data[i+1:], i+1, i+13)
		fmt.Println(string(needed), firstMatchedPos)
	}

	return maxUniqueLen
}
