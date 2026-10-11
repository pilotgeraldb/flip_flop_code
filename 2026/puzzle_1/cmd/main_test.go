package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	actual := process("testData/input.txt")
	assert.Equal(t, 1141, actual)
}
