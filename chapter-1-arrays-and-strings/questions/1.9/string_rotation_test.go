package string_rotation

import (
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	A        string
	B        string
	Expected bool
}

func TestStringRotation(t *testing.T) {
	tests := []Test{
		{
			A:        "waterbottle",
			B:        "erbottlewat",
			Expected: true,
		},
		{
			A:        "waterbottle",
			B:        "erbottlewaz",
			Expected: false,
		},
		{
			A:        "dograce",
			B:        "racedog",
			Expected: true,
		},
	}
	for _, test := range tests {
		v := IsStringRotation(test.A, test.B)
		assert.Assert(t, v == test.Expected, "failed assert - expected: %+v, got: %+v", test.Expected, v)
	}
}
