package unique

import (
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	S        string
	Expected bool
	Unique   UniqueFunc
}

func TestRealUnique(t *testing.T) {
	tests := []Test{
		{
			S:        "abcdefg",
			Expected: true,
			Unique:   IsUniqueSimple,
		},
		{
			S:        "abcdb",
			Expected: false,
			Unique:   IsUniqueSimple,
		},
		{
			S:        "abcdefg",
			Expected: true,
			Unique:   IsUniqueBitVector,
		},
		{
			S:        "abcdb",
			Expected: false,
			Unique:   IsUniqueBitVector,
		},
	}
	for _, test := range tests {
		v := test.Unique(test.S)
		assert.Assert(t, v == test.Expected, "expected unique to be true")
	}
}
