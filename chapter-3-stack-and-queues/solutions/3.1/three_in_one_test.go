package three_in_one

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestThreeInOne_Fixed(t *testing.T) {
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

type IntZero int

func (i IntZero) Zero() IntZero {
	return 0
}

func TestThreeInOne_Multi(t *testing.T) {
	ms := CreateMultiStack[IntZero](2, 10)

	ms.Push(0, 1)
	ms.Push(0, 2)
	ms.Push(0, 3)

	ms.Push(1, 4)
	ms.Push(1, 5)
	ms.Push(1, 6)

	peekValue := ms.Peek(0)
	assert.Assert(t, peekValue != nil)
	assert.Assert(t, *peekValue != 1)
}
