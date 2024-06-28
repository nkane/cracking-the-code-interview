package three_in_one

import (
	"errors"
)

/*
	3.1: Three in One: Describe how you could use a single array to implement three stacks.
*/

/*
Solution

Like many problems, this one somewhat depends on how well we'd like to support these stacks. If we're
okay with simply allocating a fixed amount of space for each stack, we can do that. This may mean though
that one stack runs out of space, while the others are nearly empty.

Alternatively, we can be flexible in our space allocation, but this significantly increases the complexity
of the problem.

Approach 1: Fixed Division
We can divide the array in three equal parts and allow the individual stack to grow in that limited space.
Note: We will use the notation '[' to mean inclusive of an end opint and "(" to mean exclusive of an end point.

- For stack 1, we will use [0, n/3).
- For stack 2, we will use [n/3, 2n/3).
- For stack 3, we will use [2n/3, n).
*/

type FixedMultiStack[T any] struct {
	NumberOfStacks int
	StackCapacity  int
	Values         []T
	Sizes          []int
}

func CreateFixedMultiStack[T any](stackSize int) *FixedMultiStack[T] {
	numberOfStacks := 3
	return &FixedMultiStack[T]{
		NumberOfStacks: numberOfStacks,
		StackCapacity:  stackSize,
		Values:         make([]T, stackSize*numberOfStacks),
		Sizes:          make([]int, numberOfStacks),
	}
}

// Push value onto stack
func (f *FixedMultiStack[T]) Push(stackNum int, value T) error {
	// check that we have space for the next element
	if f.IsFull(stackNum) {
		return errors.New("stack is full")
	}
	// increment stack pointer and then update top value
	f.Sizes[stackNum]++
	f.Values[f.IndexOfTop(stackNum)] = value
	return nil
}

// Pop item from the top stack
func (f *FixedMultiStack[T]) Pop(stackNum int) (*T, error) {
	if f.IsEmpty(stackNum) {
		return nil, errors.New("stack is empty")
	}
	return &f.Values[f.IndexOfTop(stackNum)], nil
}

// return top element
func (f *FixedMultiStack[T]) Peek(stackNum int) (*T, error) {
	if f.IsEmpty(stackNum) {
		return nil, errors.New("stack is empty")
	}
	return &f.Values[f.IndexOfTop(stackNum)], nil
}

func (f *FixedMultiStack[T]) IsEmpty(stackNum int) bool {
	return f.Sizes[stackNum] == 0
}

func (f *FixedMultiStack[T]) IsFull(stackNum int) bool {
	return f.Sizes[stackNum] == f.StackCapacity
}

func (f *FixedMultiStack[T]) IndexOfTop(stackNum int) int {
	offset := stackNum * f.StackCapacity
	size := f.Sizes[stackNum]
	return offset + size - 1
}
