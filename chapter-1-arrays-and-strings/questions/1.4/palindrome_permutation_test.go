package palindrome_permutation

import (
	"testing"

	"gotest.tools/assert"
)

func TestValidPalindromePermutation(t *testing.T) {
	sl := []string{
		"Tact Coa",
		"racecar",
		"Deed",
		"Mom",
		"Kayak",
		"Level",
	}
	for _, s := range sl {
		v := IsPalindromePermutation(s)
		assert.Assert(t, v == true, "expected valid palindrome permutation")
	}
}
