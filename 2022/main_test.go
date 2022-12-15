package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateChar(t *testing.T) {
	assert := assert.New(t)

	nonDistinctionStr := "abbccdefggaaf"
	distinctionStr := "abcdefg"
	assert.True(existDuplicateChar(nonDistinctionStr), true)
	assert.False(existDuplicateChar(distinctionStr), false)
}
