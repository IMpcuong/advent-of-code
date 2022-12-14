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

func RecursiveSearch(char byte, seq string, beginIdx, endIdx int) int {
	if endIdx < beginIdx {
		return -1
	}
	if seq[beginIdx] == char {
		return beginIdx
	}
	if seq[endIdx] == char {
		return endIdx
	}
	return RecursiveSearch(char, seq, beginIdx+1, endIdx-1)
}

func detectMsgMarkerV1(data string, longest int) int {
	for curPos := longest; curPos <= len(data); curPos++ {
		// NOTE: The ordinary implementation for the data structure `map` was built on hashmap/swiss-table inspiration.
		mapChar := make(map[rune]struct{}) // Equals to: `map[rune]struct{}{}`.
		for _, r := range data[curPos-longest : curPos] {
			mapChar[r] = *new(struct{}) // Equals to: `struct{}{}`.
		}
		if len(mapChar) >= longest {
			return curPos
		}
	}
	return -1
}

func detectMsgMarkerV2(data string, longest int) int {
	for curPos := longest; curPos <= len(data); curPos++ {
		begin := curPos - longest
		needed := data[begin]
		for chunkIdx, char := range data[begin+1 : curPos] {
			if needed == byte(char) && chunkIdx-begin >= longest {
				return chunkIdx
			}
			begin++
		}
	}

	return -1
}
