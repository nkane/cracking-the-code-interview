package queue_via_stack

import (
	"log"
	"stack"
)

/*
	3.4: Queue via Stack: Implement a MyQueue struct which implements a queue using two stacks.
*/

/*
Since the major difference between a queue and a stack is the order (first-in first-out vs. last-in first-out),
we know that we need to modify peek() and pop() to go in reverse order. We can use our second stack to reverse
the order of the elements (by popping s1 and pusing the elements on to s2). In such an implementation, on each
peek() and pop() operation, we would pop everything from s1 onto s2, perform the peek / pop operation, and then
push everything back.

This will work, but if pop / peeks are performed back-to-back, we're needlessly moving elements. We can implement
a "lazy" approach where we let the elements sit in s2 until we absolutely must reverse the elements.

In this approach, StackNewest has the newest elements on top and StackOldest has the oldes elements on top. When
we dequeue an element, we want to remove the oldest element first, and so we dequeue from StackOldest. If StackOldest
is empty, then we want to transder all elements from StackNewest into this stack in reverse order. To insert an
element, we push onto StackNewest, since it has the newest elements on top.
*/

type Queue[T any] struct {
	StackNewest *stack.Stack[T]
	StackOldest *stack.Stack[T]
}

func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{
		StackNewest: &stack.Stack[T]{},
		StackOldest: &stack.Stack[T]{},
	}
}

func (q *Queue[T]) Size() int {
	size := q.StackNewest.Size() + q.StackOldest.Size()
	return size
}

func (q *Queue[T]) Add(value T) {
	// push onto StackNewest, which always has the newest elements on top
	q.StackNewest.Push(value)
}

func (q *Queue[T]) ShiftStacks() {
	// move elements from StackNewest into StackOldest
	// this is usually done so that we can do operations on StackOldest
	if q.StackOldest.IsEmpty() {
		for !q.StackNewest.IsEmpty() {
			v, err := q.StackNewest.Pop()
			if err != nil && v != nil {
				log.Printf("stack shift error: %+v\n", err)
			} else {
				q.StackOldest.Push(*v)
			}
		}
	}
}

func (q *Queue[T]) Peek() *T {
	q.ShiftStacks() // ensure StackOldest has the current elements
	v, err := q.StackOldest.Peek()
	if err != nil {
		log.Printf("peek error: %+v\n", err)
	}
	return v
}

func (q *Queue[T]) Remove() *T {
	q.ShiftStacks()               // ensure StackOldest has the current elements
	v, err := q.StackOldest.Pop() // pop the oldest item
	if err != nil {
		log.Printf("remove error: %+v\n", err)
	}
	return v
}
