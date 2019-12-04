package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestProblem3a(t *testing.T) {
	testCases := []struct {
		wire1 string
		wire2 string
		expected int
	} {
		{
			wire1: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2: "U62,R66,U55,R34,D71,R55,D58,R83",
			expected: 159,
		},
		{
			wire1: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expected: 135,
		},
	}

	for _, tc := range testCases {
		wire1Coordinates := parseWireInput(tc.wire1)
		wire2Coordinates := parseWireInput(tc.wire2)
		matches := wireCoordinateMatches(wire1Coordinates, wire2Coordinates)
		assert.Equal(t, shortestManhattanMatch(matches), tc.expected)
	}
}

func TestProblem3b(t *testing.T) {
	testCases := []struct {
		wire1 string
		wire2 string
		expected int
	} {
		{
			wire1: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2: "U62,R66,U55,R34,D71,R55,D58,R83",
			expected: 610,
		},
		{
			wire1: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expected: 410,
		},
	}

	for _, tc := range testCases {
		wire1Coordinates := parseWireInput(tc.wire1)
		wire2Coordinates := parseWireInput(tc.wire2)
		matches := wireCoordinateMatchesAfterStepCount(wire1Coordinates, wire2Coordinates)
		assert.Equal(t, shortestMatchStepCount(matches), tc.expected)
	}
}

