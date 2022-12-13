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
			return i + 3
		}
	}

	return -1
}
