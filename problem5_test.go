package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem5aParseFirstInstruction(t *testing.T) {
	computer := NewIntCodeComputer("1002,4,3,4,33", 1)
	computer.parseCurrentInstruction()

	assert.Equal(t, 2, computer.opcode)
	assert.Equal(t, PositionMode, computer.parameterMode1)
	assert.Equal(t, ImmediateMode, computer.parameterMode2)
	assert.Equal(t, PositionMode, computer.parameterMode3)
}

func TestProblem5aParseFirstInstruction2(t *testing.T) {
	computer := NewIntCodeComputer("3,4,3,4,33", 1)
	computer.parseCurrentInstruction()

	assert.Equal(t, 3, computer.opcode)
	assert.Equal(t, PositionMode, computer.parameterMode1)
	assert.Equal(t, PositionMode, computer.parameterMode2)
	assert.Equal(t, PositionMode, computer.parameterMode3)
}


func TestProblem5aCompute(t *testing.T) {
	computer := NewIntCodeComputer("1002,4,3,4,33,99", 1)
	computer.Begin()

	assert.Equal(t, 99, computer.opcode)
}