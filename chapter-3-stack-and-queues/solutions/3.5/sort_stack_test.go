package sort_stack

import "testing"

func TestSortedStack(t *testing.T) {
	s := SortedStack{}
	s.Push(10)
	s.Push(3)
	s.Push(9)
	s.Push(1)
	s.Push(12)
	s.Push(2)
	s.Sort()

	s.Pop()
}
