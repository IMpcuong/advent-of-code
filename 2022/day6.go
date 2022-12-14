package main

import (
	_ "embed"
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

func RecursiveSearch(char byte, seq string, begin, end int) int {
	if end < 1 {
		return -1
	}
	if seq[begin] == char {
		return begin
	}
	if seq[end] == char {
		return end
	}
	return RecursiveSearch(char, seq, begin+1, end-1)
}

func detectMsgMarkerV1(data string, longest int) int {
	for cusPos := longest; cusPos <= len(data); cusPos++ {
		// NOTE: The ordinary implementation for the data structure `map` was built on hashmap/swiss-table inspiration.
		mapChar := make(map[rune]struct{}) // Equals to: `map[rune]struct{}{}`
		for _, r := range data[cusPos-longest : cusPos] {
			mapChar[r] = *new(struct{}) // Equals to: `struct{}{}`.
		}
		if len(mapChar) >= longest {
			return cusPos
		}
	}
	return -1
}

func detectMsgMarkerV2(data string, longest int) int {
	for i, char := range data {
		return i + int(char)
	}
	return -1
}
