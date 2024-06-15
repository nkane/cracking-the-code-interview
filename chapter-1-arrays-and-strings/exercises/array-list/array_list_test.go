package array_list

import (
	"testing"

	"gotest.tools/assert"
)

func TestArrayListAppend(t *testing.T) {
	al := CreateArrayList()
	al.Append(1)
	al.Append(2)
	al.Append(3)
	al.Append(4)
	assert.Assert(t, al.Capacity == 4, "array list capacity check failed - capacity: %d", al.Capacity)
	assert.Assert(t, al.Size == 4, "array list size check failed - size: %d", al.Size)
}
