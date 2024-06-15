package permutation

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestValidPermutation(t *testing.T) {
	a := "tech"
	b := "hetc"
	v := IsPermutation(a, b)
	assert.Assert(t, v == true, "expected a and b to be valid permutation")
}

func TestInvalidPermutation(t *testing.T) {
	a := "tech"
	b := "tetc"
	v := IsPermutation(a, b)
	assert.Assert(t, v == false, "expected a and b to be invalid permutation")
}

func TestInvalidPermutationLengths(t *testing.T) {
	a := "tech"
	b := "tetct"
	v := IsPermutation(a, b)
	assert.Assert(t, v == false, "expected a and b to be invalid permutation lengths")
}
