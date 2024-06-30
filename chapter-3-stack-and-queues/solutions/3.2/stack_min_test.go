package stack_min

import (
	"testing"

	"gotest.tools/assert"
)

func TestStackMin(t *testing.T) {
	minStack := MinStack[int]{}

	minStack.Push(5)
	assert.Assert(t, minStack.Min() != nil)
	assert.Assert(t, *minStack.Min() == 5)

	minStack.Push(6)
	assert.Assert(t, minStack.Min() != nil)
	assert.Assert(t, *minStack.Min() == 5)

	minStack.Push(3)
	assert.Assert(t, minStack.Min() != nil)
	assert.Assert(t, *minStack.Min() == 3)

	minStack.Push(7)
	assert.Assert(t, minStack.Min() != nil)
	assert.Assert(t, *minStack.Min() == 3)

	v, err := minStack.Pop()
	assert.NilError(t, err)
	assert.Assert(t, v != nil)
	assert.Assert(t, *v == 7)
	assert.Assert(t, minStack.Min() != nil)
	assert.Assert(t, *minStack.Min() == 3)

	v, err = minStack.Pop()
	assert.NilError(t, err)
	assert.Assert(t, v != nil)
	assert.Assert(t, *v == 3)
	assert.Assert(t, minStack.Min() != nil)
	assert.Assert(t, *minStack.Min() == 5)
}

func TestStackMin2(t *testing.T) {
	minStack := StackWithMin2{}

	minStack.Push(5)
	assert.Assert(t, minStack.Min() == 5)

	minStack.Push(6)
	assert.Assert(t, minStack.Min() == 5)

	minStack.Push(3)
	assert.Assert(t, minStack.Min() == 3)

	minStack.Push(7)
	assert.Assert(t, minStack.Min() == 3)

	v := minStack.Pop()
	assert.Assert(t, v == 7)
	assert.Assert(t, minStack.Min() == 3)

	v = minStack.Pop()
	assert.Assert(t, v == 3)
	assert.Assert(t, minStack.Min() == 5)
}
