package string_compression

import (
	"testing"

	"gotest.tools/assert"
)

func TestValidStringCompression(t *testing.T) {
	input := "aabcccccaaa"
	output := "a2b1c5a3"
	v := CompressString(input)
	assert.Assert(t, output == v, "failed output comparison - expected: %s, got: %s", output, v)
}
