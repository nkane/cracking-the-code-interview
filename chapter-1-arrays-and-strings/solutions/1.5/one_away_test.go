package one_away

import (
	"testing"

	"gotest.tools/assert"
)

type TestInput struct {
	A              string
	B              string
	ExpectedOutput bool
	OneAway        OneAway
}

func TestValidOneAway(t *testing.T) {
	tests := []TestInput{
		{
			A:              "pale",
			B:              "ple",
			ExpectedOutput: true,
			OneAway:        OneAwaySolutionOne,
		},
		{
			A:              "pales",
			B:              "pale",
			ExpectedOutput: true,
			OneAway:        OneAwaySolutionOne,
		},
		{
			A:              "pale",
			B:              "bale",
			ExpectedOutput: true,
			OneAway:        OneAwaySolutionOne,
		},
		{
			A:              "pale",
			B:              "bake",
			ExpectedOutput: false,
			OneAway:        OneAwaySolutionOne,
		},
		{
			A:              "pale",
			B:              "ple",
			ExpectedOutput: true,
			OneAway:        OneAwaySolutionTwo,
		},
		{
			A:              "pales",
			B:              "pale",
			ExpectedOutput: true,
			OneAway:        OneAwaySolutionTwo,
		},
		{
			A:              "pale",
			B:              "bale",
			ExpectedOutput: true,
			OneAway:        OneAwaySolutionTwo,
		},
		{
			A:              "pale",
			B:              "bake",
			ExpectedOutput: false,
			OneAway:        OneAwaySolutionTwo,
		},
	}
	for _, test := range tests {
		v := test.OneAway(test.A, test.B)
		assert.Assert(t, v == test.ExpectedOutput, "failed assert check a: %s - b: %s", test.A, test.B)
	}
}
