package stack_min

import (
	"errors"
	"math"
	"stack"

	"golang.org/x/exp/constraints"
)

/*
	3.2: Stack Min: How would you design a stack, which in addition to push and pop, has a function
	min which returns the minimum element? Push, pop, and min should all operate in O(1) time.
*/

/*
The thing with minimums is that they don't change very often. They only change when a smaller element
is added.

One solution is to have just a single int value, minValue, that's a member of the stack struct. When
minValue is popped fromt he stack, we search through the stack to find the new minimum. Unfortunately,
this would break the constraint that push and pop operate in O(1) time.

To further understand this question, let's walk through it with a short example:

push(5);	// stack is {5}, min is 5
push(6);	// stack is {6,5}, min is 5
push(3);	// stack is {3,6,5}, min is 3
push(7);	// stack is {7,3,6,5}, min is 3
pop();		// pops 7, stack is {3, 6, 5}, min is 3
pop();		// pops 3, stack is {6, 5}, min is 5

Observe how once the stack goes back to a prior stack ({6,5}), the minimum also goes back to its prior state (5).
This leads us to our second solution.

If we kept track of the minimum at each state, we would be able to easily know the minimum. We can do this by
having each node record what the minimum beneath itself is. Then, to find the min, you just look at what the top
element thinks is the min.

When you pus an element on to the stack, the element is given the current minimum. It sets its "local min" to be
the min.
*/

type StackNode[T constraints.Ordered] struct {
	Data T
	Min  T
	Next *StackNode[T]
}

type MinStack[T constraints.Ordered] struct {
	Top *StackNode[T]
}

func (ms *MinStack[T]) Pop() (*T, error) {
	if ms.Top == nil {
		return nil, errors.New("empty stack")
	}
	item := ms.Top.Data
	ms.Top = ms.Top.Next
	return &item, nil
}

func (ms *MinStack[T]) Push(item T) {
	t := StackNode[T]{
		Data: item,
		Next: ms.Top,
	}
	if ms.Top == nil || item < ms.Top.Min {
		t.Min = item
	} else {
		t.Min = ms.Top.Min
	}
	ms.Top = &t
}

func (ms *MinStack[T]) Peek() (*StackNode[T], error) {
	if ms.Top == nil {
		return nil, errors.New("empty stack")
	}
	return ms.Top, nil
}

func (ms *MinStack[T]) IsEmpty() bool {
	return ms.Top == nil
}

func (ms *MinStack[T]) Min() *T {
	if ms.IsEmpty() {
		return nil
	} else {
		v, err := ms.Peek()
		if err != nil {
			return nil
		}
		return &v.Min
	}
}

/*
There's just one issue with this, if we have a large stadck, we waste a lot of space by keeping track of the min
for every single element. Can we do better?

We can (maybe) do a bit better than this by using an additional stack which keeps track of the mins.

Why might this be more space efficient? Suppose we had a very large stack and the first element insereted
happened to be the minimum. In the first solution, we would be keeping n integers, where n is the size of
the stack. In the second soution, we just store a few pieces of data: a second stack with one element
and the members within this stack.
*/

type StackWithMin2 struct {
	stack.Stack[int]
	S2 stack.Stack[int]
}

func (s *StackWithMin2) Push(value int) {
	if value <= s.Min() {
		s.S2.Push(value)
	}
	s.Stack.Push(value)
}

func (s *StackWithMin2) Pop() int {
	v, err := s.Stack.Pop()
	if err != nil || v == nil {
		return -1
	}
	if *v == s.Min() {
		s.S2.Pop()
	}
	return *v
}

func (s *StackWithMin2) Min() int {
	if s.S2.IsEmpty() {
		return math.MaxInt64
	} else {
		v, err := s.S2.Peek()
		if err != nil || v == nil {
			return math.MaxInt64
		}
		return *v
	}
}
