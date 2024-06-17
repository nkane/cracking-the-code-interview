package palindrome_permutation

import (
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Input                   string
	Expected                bool
	IsPalindromePermutation IsPalindromePermutation
}

func TestValidPalindromePermutation(t *testing.T) {
	tests := []Test{
		{
			Input:                   "tact coa",
			Expected:                true,
			IsPalindromePermutation: IsPalindromePermutationSolutionOne,
		},
		{
			Input:                   "tact coa",
			Expected:                true,
			IsPalindromePermutation: IsPalindromePermutationSolutionTwo,
		},
		{
			Input:                   "tact coa",
			Expected:                true,
			IsPalindromePermutation: IsPalindromePermutationSolutionThree,
		},
		{
			Input:                   "tac coa",
			Expected:                false,
			IsPalindromePermutation: IsPalindromePermutationSolutionOne,
		},
		{
			Input:                   "tac coa",
			Expected:                false,
			IsPalindromePermutation: IsPalindromePermutationSolutionTwo,
		},
		{
			Input:                   "tac coa",
			Expected:                false,
			IsPalindromePermutation: IsPalindromePermutationSolutionThree,
		},
	}
	for _, test := range tests {
		v := test.IsPalindromePermutation(test.Input)
		assert.Assert(t, v == test.Expected, "expected valid palindrome permutation")
	}
}
