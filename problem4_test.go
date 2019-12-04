package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem4a(t *testing.T) {
	passwords := []struct{
		password string
		meetsCondition bool
	}{
		{
			password: "111111",
			meetsCondition: true,
		},
		{
			password: "223450",
			meetsCondition: false,
		},
		{
			password: "123789",
			meetsCondition: false,
		},
	}

	for _, tc := range passwords {
		pc :=PasswordChecker{
			start: 111110,
			end:   675869,
		}
		meets, reason := pc.checkPassword(tc.password)
		assert.Equal(t, tc.meetsCondition, meets, fmt.Sprintf("password %v.. %v",tc.password,reason))
	}
}

func TestProblem4b(t *testing.T) {
	passwords := []struct{
		password string
		meetsCondition bool
	}{
		{
			password: "112233",
			meetsCondition: true,
		},
		{
			password: "123444",
			meetsCondition: false,
		},
		{
			password: "111122",
			meetsCondition: true,
		},
	}

	for _, tc := range passwords {
		pc :=PasswordChecker{
			start: 111110,
			end:   675869,
		}
		meets := pc.checkPasswordSecondCriteria(tc.password)
		assert.Equal(t, tc.meetsCondition, meets)
	}
}