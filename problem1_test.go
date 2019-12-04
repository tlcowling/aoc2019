package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem1a(t *testing.T) {
	testMasses := []struct {
		Mass string
		Expected int
	} {
		{
			Mass: "12",
			Expected: 2,
		},
		{
			Mass: "14",
			Expected: 2,
		},
		{
			Mass: "1969",
			Expected: 654,
		},
		{
			Mass: "100756",
			Expected: 33583,
		},
	}

	for _, testMass := range testMasses {
		actual := getFuel(testMass.Mass)
		assert.Equal(t, testMass.Expected, actual, "Get Fuel should return the correct amount.")
	}
}

func TestProblem1b(t *testing.T) {
	testMasses := []struct {
		Mass string
		Expected int
	} {
		{
			Mass: "14",
			Expected: 2,
		},
		{
			Mass: "1969",
			Expected: 966,
		},
		{
			Mass: "100756",
			Expected: 50346,
		},
	}

	for _, testMass := range testMasses {
		actual := getFuelImproved(testMass.Mass)
		assert.Equal(t, testMass.Expected, actual, "Get Fuel should return the correct amount.")
	}
}