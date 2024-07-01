package stack_of_plates

// My solution is 30 minutes
/*
	3.3: Stack of Plates: Image a (literal) stack of plates. If the stack gets too high, it might topple. Therefore,
	in real life, we would likely start a new start when the previous stack exceeds some threshold. Implement a
	data strcture SetOfStacks that mimics this. SetOfStacks should be composed of several stacks and should create
	a new stack once the previous one exceeds capacity. SetOfStacks.Push() and SetOfStacks.Pop() should behave
	identically to a single stack (that is, pop() should return the same value as if it would if there were just
	a single stack).

	Follow up
	Implement a function PopAt(index int) which performs a pop operation on a specific sub-stack.
*/

type SubStack[T any] struct {
	Stack      []T
	CurrentIdx int
	Size       int
}

func CreateSubStack[T any](capacity int) SubStack[T] {
	return SubStack[T]{
		Stack:      make([]T, capacity),
		CurrentIdx: -1,
		Size:       0,
	}
}

func (subStack *SubStack[T]) Push(item T) {
	subStack.CurrentIdx++
	subStack.Stack[subStack.CurrentIdx] = item
	subStack.Size++
}

func (subStack *SubStack[T]) Pop() *T {
	item := subStack.Stack[subStack.CurrentIdx]
	subStack.CurrentIdx--
	subStack.Size--
	return &item
}

func (substack *SubStack[T]) IsEmpty() bool {
	return substack.Size == len(substack.Stack)
}

func (subStack *SubStack[T]) IsFull() bool {
	return len(subStack.Stack) == subStack.Size
}

type StackSet[T any] struct {
	Stacks          []SubStack[T]
	CurrentStackIdx int
	Capacity        int
}

func CreateStackSet[T any](capacity int) *StackSet[T] {
	return &StackSet[T]{
		Stacks: []SubStack[T]{
			CreateSubStack[T](capacity),
		},
		Capacity:        capacity,
		CurrentStackIdx: 0,
	}
}

func (stackSet *StackSet[T]) Push(item T) {
	if stackSet.Stacks[stackSet.CurrentStackIdx].IsFull() {
		// create new substack
		stackSet.Stacks = append(stackSet.Stacks, CreateSubStack[T](stackSet.Capacity))
		stackSet.CurrentStackIdx++
	}
	stackSet.Stacks[stackSet.CurrentStackIdx].Push(item)
}

func (stackSet *StackSet[T]) Pop() *T {
	if stackSet.Stacks[stackSet.CurrentStackIdx].IsEmpty() {
		// delete current empty stackSet
		stackSet.Stacks = stackSet.Stacks[0 : stackSet.CurrentStackIdx-1]
		stackSet.CurrentStackIdx--
	}
	return stackSet.Stacks[stackSet.CurrentStackIdx].Pop()
}
