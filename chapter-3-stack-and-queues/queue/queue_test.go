package queue

import (
	"testing"

	"gotest.tools/assert"
)

func TestQueue(t *testing.T) {
	queue := Queue[int]{}
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Add(4)
	queue.Add(5)

	peekedItem, err := queue.Peek()
	assert.NilError(t, err)
	assert.Assert(t, peekedItem != nil)
	assert.Assert(t, *peekedItem == 1, "exepcted: 1, got: %d", *peekedItem)

	removedItem, err := queue.Remove()
	assert.NilError(t, err)
	assert.Assert(t, removedItem != nil)
	assert.Assert(t, *removedItem == 1, "exepcted: 1, got: %d", *peekedItem)

	isEmpty := queue.IsEmpty()
	assert.Assert(t, isEmpty == false)
}
