package stack

import (
	"testing"

	"gotest.tools/assert"
)

func TestStack(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	peekedItem, err := stack.Peek()
	assert.NilError(t, err)
	assert.Assert(t, peekedItem != nil)
	assert.Assert(t, *peekedItem == 5)

	poppedItem, err := stack.Pop()
	assert.NilError(t, err)
	assert.Assert(t, poppedItem != nil)
	assert.Assert(t, *poppedItem == 5)

	isEmpty := stack.IsEmpty()
	assert.Assert(t, isEmpty == false)
}
