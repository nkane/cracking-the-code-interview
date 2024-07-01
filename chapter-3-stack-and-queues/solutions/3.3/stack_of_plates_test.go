package stack_of_plates

import "testing"

func TestStackSet(t *testing.T) {
	stackSet := CreateStackSet[int](3)
	// first stack
	stackSet.Push(1)
	stackSet.Push(2)
	stackSet.Push(3)
	// second stack
	stackSet.Push(4)
	stackSet.Push(5)
	stackSet.Push(6)
}
