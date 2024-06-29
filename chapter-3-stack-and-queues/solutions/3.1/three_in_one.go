package three_in_one

import (
	"errors"
	"log"
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

If we had additional information about the expected usages of the stacks, then we could modify this algorithm
accordingly. For example, if we expected Stack 1 to have many more elements than Stack 2, we could allocate
more space to Stack 1 and less space to Stack 2.
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

/*
Approach 2: Flexible Divisions

A second approach is to allow the stack blocks to be flexible in size. When one stack exceeds its initial
capacity, we grow the allowable capacity and shift elements as necessary.

We will also design our array to be circular, such that the final stack may start at the end of the array
and wrap around to the beginning.

Please not that the code for this solution is far more complex than would be appropriate for an interview.
You could be responsible for pseudocode, or perhaps the code of individual components, but the entire
implementatio would be far too much work.
*/

type StackInfo struct {
	Start    int
	Size     int
	Capacity int
	Max      int
}

func CreateStackInfo(start int, capacity int, max int) StackInfo {
	return StackInfo{
		Start:    start,
		Capacity: capacity,
		Size:     0,
	}
}

// check if an index on the array is within the stack boundaries. The stack can wrap around to the start
// of the array.
func (si *StackInfo) IsWithintStackCapacity(index int) bool {
	// if outside of bounds of array, return flase
	if index < 0 || index >= si.Max {
		return false
	}
	// if index wraps around, adjust it
	contiguousIndex := index
	if index < si.Start {
		contiguousIndex = index + si.Max
	}
	end := si.Start + si.Capacity
	return si.Start <= contiguousIndex && contiguousIndex < end
}

func (si *StackInfo) LastCapacityIndex() int {
	return -1
}

func (si *StackInfo) LastElementIndex() int {
	return -1
}

func (si *StackInfo) IsFull() bool {
	return si.Size == si.Capacity
}

func (si *StackInfo) IsEmpty() bool {
	return si.Size == 0
}

type MultiStack[T Zeroer[T]] struct {
	StackInfo []StackInfo
	Values    []T
}

type Zeroer[T any] interface {
	Zero() T
}

func CreateMultiStack[T Zeroer[T]](numberOfStacks int, defaultSize int) *MultiStack[T] {
	// create metadata for all the stack
	max := numberOfStacks * defaultSize
	info := make([]StackInfo, numberOfStacks)
	for i := 0; i < numberOfStacks; i++ {
		info[i] = CreateStackInfo(defaultSize*i, defaultSize, max)
	}
	return &MultiStack[T]{
		StackInfo: info,
		Values:    make([]T, max),
	}
}

// push value onto stack num, shifting/expanding stacks as necessary.
func (ms *MultiStack[T]) Push(stackNum int, value T) error {
	if ms.AllStacksAreFull() {
		return errors.New("all stacks are full")
	}
	stack := &ms.StackInfo[stackNum]
	// if this stack is full, expand it
	if stack.IsFull() {
		ms.Expand(stackNum)
	}
	// find the index of the top element in the array + 1, and
	// increment the stack pointer
	stack.Size++
	ms.Values[stack.LastElementIndex()] = value
	return nil
}

func (ms *MultiStack[T]) Pop(stackNum int) (*T, error) {
	stack := ms.StackInfo[stackNum]
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	value := ms.Values[stack.LastElementIndex()]
	ms.Values[stack.LastElementIndex()] = ms.Values[stack.LastElementIndex()].Zero()
	stack.Size--
	return &value, nil
}

func (ms *MultiStack[T]) Peek(stackNum int) *T {
	stack := ms.StackInfo[stackNum]
	return &ms.Values[stack.LastElementIndex()]
}

// returns the number of items actually present in the stack
func (ms *MultiStack[T]) NumberOfElements() int {
	size := 0
	for _, si := range ms.StackInfo {
		size += si.Size
	}
	return size
}

func (ms *MultiStack[T]) AllStacksAreFull() bool {
	return ms.NumberOfElements() == len(ms.Values)
}

func (ms *MultiStack[T]) Expand(stackNum int) {
	ms.Shift((stackNum + 1) % len(ms.StackInfo))
	ms.StackInfo[stackNum].Capacity++
}

func (ms *MultiStack[T]) AdjustIndex(index int) int {
	max := len(ms.Values)
	return ((index % max) + max) % max
}

func (ms *MultiStack[T]) NextIndex(index int) int {
	return ms.AdjustIndex(index + 1)
}

func (ms *MultiStack[T]) PreviousIndex(index int) int {
	return ms.AdjustIndex(index - 1)
}

// shift items in stack over by one element, if we have available capacity, then
// we'll end up shrinking the stack by one element, if we don't have available
// capacity, then we'll need to shift the next stack over too.
func (ms *MultiStack[Zeroer]) Shift(stackNum int) {
	log.Printf("/// Shifting %d", stackNum)
	stack := &ms.StackInfo[stackNum]
	// if this stack is at its full capacity, then you need to move the next
	// stack over by one element, this stack can now claim the freed index.
	if stack.Size >= stack.Capacity {
		nextStack := (stackNum + 1) % len(ms.StackInfo)
		ms.Shift(nextStack)
		stack.Capacity++ // claim index that next stack lost
	}
	// shift all elements in stack over by one
	index := stack.LastCapacityIndex()
	for stack.IsWithintStackCapacity(index) {
		ms.Values[index] = ms.Values[ms.PreviousIndex(index)]
		index = ms.PreviousIndex(index)
	}
	// adjust stack data
	ms.Values[stack.Start] = ms.Values[stack.Start].Zero()
	stack.Start = ms.NextIndex(stack.Start) // move start
	stack.Capacity--                        // shrink capacity
}
