package one_away

import (
	"testing"

	"gotest.tools/assert"
)

type TestInput struct {
	A              string
	B              string
	ExpectedOutput bool
}

func TestValidOneAway(t *testing.T) {
	l := []TestInput{
		{
			A:              "pale",
			B:              "ple",
			ExpectedOutput: true,
		},
		{
			A:              "pales",
			B:              "pale",
			ExpectedOutput: true,
		},
		{
			A:              "pale",
			B:              "bale",
			ExpectedOutput: true,
		},
		{
			A:              "pale",
			B:              "bake",
			ExpectedOutput: false,
		},
	}
	for _, ti := range l {
		v := OneAway(ti.A, ti.B)
		assert.Assert(t, v == ti.ExpectedOutput, "failed assert check a: %s - b: %s", ti.A, ti.B)
	}
}
