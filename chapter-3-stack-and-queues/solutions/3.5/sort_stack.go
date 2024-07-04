package sort_stack

import (
	"stack"
)

/*
	3.5: Sort Stack: Write a program to sort a stack such that the smallest items are on the top.
	You can use an additional temporary stack, but you may not copy elements into any other data
	structure (such as an array). The stack supports the following operations: push, pop, peek,
	and IsEmpty.
*/

/*
One approach is to implement a rudimentary sorting algorithm. We search through the entire stack to find
the minimum element and then push that onto a new stack. Then, we find the new minimum element and push
that. This will actually require a total of three stacks: s1 is the original stack, s2 is the final sorted
stack, and s3 acts as a buffer during our searching of s1. To search s1 for each minimum, we need to pop
elements from s1 and push them onto the buffer, s3.

unfortunately, this requires two additional stacks, and we can only use one. Can we do better? Yes.

Rather than searching for the minimum repeatedly, we can sort s1 by inserting each element form s1 in order
into s2.

This algorithm is O(N^2) time and O(N) space.
*/

type SortedStack struct {
	stack.Stack[int]
}

func (s *SortedStack) Sort() {
	r := stack.Stack[int]{}
	for !s.IsEmpty() {
		// insert each element in s in sorted order into r
		tmp, _ := s.Pop()
		v, _ := r.Peek()
		for !r.IsEmpty() && *v > *tmp {
			value, _ := r.Pop()
			s.Push(*value)
		}
		r.Push(*tmp)
	}
	// copy the elements from r back into s
	for !r.IsEmpty() {
		v, _ := r.Pop()
		s.Push(*v)
	}
}
