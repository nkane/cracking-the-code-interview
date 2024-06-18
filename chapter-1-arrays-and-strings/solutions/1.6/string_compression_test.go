package string_compression

import (
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Input          string
	Expected       string
	CompressString CompressString
}

func TestValidStringCompression(t *testing.T) {
	tests := []Test{
		{
			Input:          "aabcccccaaa",
			Expected:       "a2b1c5a3",
			CompressString: CompressStringBad,
		},
		{
			Input:          "abc",
			Expected:       "abc",
			CompressString: CompressStringBad,
		},
		{
			Input:          "aabcccccaaa",
			Expected:       "a2b1c5a3",
			CompressString: CompressStringSB,
		},
		{
			Input:          "abc",
			Expected:       "abc",
			CompressString: CompressStringSB,
		},
		{
			Input:          "aabcccccaaa",
			Expected:       "a2b1c5a3",
			CompressString: CompressStringCheckFirst,
		},
		{
			Input:          "abc",
			Expected:       "abc",
			CompressString: CompressStringCheckFirst,
		},
	}
	for _, test := range tests {
		v := test.CompressString(test.Input)
		assert.Assert(t, test.Expected == v, "failed output comparison - expected: %s, got: %s", test.Expected, v)
	}
}
