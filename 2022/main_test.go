package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateChar(t *testing.T) {
	assert := assert.New(t)

	nonDistinctionStr := "jlltwtsszwswgs"
	distinctionStr := "abcdefg"
	assert.True(uniqueChars(nonDistinctionStr), false)
	assert.True(uniqueChars(distinctionStr), true)
}
