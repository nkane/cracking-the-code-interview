package three_in_one

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestThreeInOne(t *testing.T) {
	fs := CreateFixedMultiStack[int](9)
	fs.Push(0, 1)
	fs.Push(0, 2)
	fs.Push(0, 3)

	fs.Push(1, 4)
	fs.Push(1, 5)
	fs.Push(1, 6)

	fs.Push(2, 7)
	fs.Push(2, 8)
	fs.Push(2, 9)

	peekValue, err := fs.Peek(1)
	assert.NilError(t, err)
	assert.Assert(t, peekValue != nil)
	assert.Assert(t, *peekValue != 1)
}
