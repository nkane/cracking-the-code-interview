package permutation

import (
	"testing"

	"gotest.tools/v3/assert"
)

type Test struct {
	A        string
	B        string
	Expected bool
}

func TestValidPermutation(t *testing.T) {
	tests := []Test{
		{
			A:        "tech",
			B:        "hetc",
			Expected: true,
		},
		{
			A:        "tech",
			B:        "tetc",
			Expected: false,
		},
		{
			A: "tech",
			B: "tetct",
		},
	}

	for _, test := range tests {
		v := IsPermutation(test.A, test.B)
		assert.Assert(t, v == test.Expected, "expected a and b to be valid permutation")
	}
}
