package permutation

import (
	"testing"

	"gotest.tools/v3/assert"
)

type Test struct {
	A             string
	B             string
	Expected      bool
	IsPermutation IsPermutation
}

func TestValidPermutation(t *testing.T) {
	tests := []Test{
		{
			A:             "tech",
			B:             "hetc",
			Expected:      true,
			IsPermutation: IsPermutationSolutionOne,
		},
		{
			A:             "tech",
			B:             "tetc",
			Expected:      false,
			IsPermutation: IsPermutationSolutionOne,
		},
		{
			A:             "tech",
			B:             "tetct",
			IsPermutation: IsPermutationSolutionOne,
		},
		{
			A:             "tech",
			B:             "hetc",
			Expected:      true,
			IsPermutation: IsPermutationSolutionTwo,
		},
		{
			A:             "tech",
			B:             "tetc",
			Expected:      false,
			IsPermutation: IsPermutationSolutionTwo,
		},
		{
			A:             "tech",
			B:             "tetct",
			IsPermutation: IsPermutationSolutionTwo,
		},
	}

	for _, test := range tests {
		v := test.IsPermutation(test.A, test.B)
		assert.Assert(t, v == test.Expected, "expected a and b to be valid permutation")
	}
}
