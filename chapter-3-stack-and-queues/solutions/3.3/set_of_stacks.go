package stack_of_plates

import "errors"

// Book's solution(s)

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

/*
We know that Puhs() should behave idententical to a single stack, which means that we need to Push() to call
Push() on the last stack in the array of stacks. We hwave to be a bit careful though, if the last stack is at
capacity, we need to create a new stack.

What should Pop() do? It should behave similarly to Push() in that it should operate on the last stack. If
the last stack is empty (after popping), then we should remove the stack from the list of stacks.

Follow up: Implement PopAt(int index)
This is a bit tricker to implement, but we can imagine a "rollover" system. If we pop an element from stack 1,
we need to remove the bottom of stack 2 and push it onto stack 1. We then need to rollover from stack 3 to
stack 2, stack 4, to stack 3, and etc.

You could make an argument list, rather that "rolling over", we should be okay with some stacks not being at
ful capacity. This would improve the time complexity (by a fair amount, with a large number of elements), but
it might get us into tricky situations later on if someone assumes that all stacks (other than the last) operate
at full capacity. There's no "right answer" here; you should discuss this trade-off with your interviewer.
*/

type SetOfStacks struct {
	Stacks   []*Stack
	Capacity int
}

func CreateSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{
		Capacity: capacity,
	}
}

func (s *SetOfStacks) GetLastStack() *Stack {
	if len(s.Stacks) == 0 {
		return nil
	}
	return s.Stacks[len(s.Stacks)-1]
}

func (s *SetOfStacks) Push(v int) {
	last := s.GetLastStack()
	if last != nil && last.IsFull() {
		last.Push(v)
	} else {
		stack := CreateStack(s.Capacity)
		stack.Push(v)
		s.Stacks = append(s.Stacks, stack)
	}
}

func (s *SetOfStacks) Pop() (int, error) {
	last := s.GetLastStack()
	if last == nil {
		return -1, errors.New("empty stack")
	}
	v := last.Pop()
	if last.Size == 0 {
		s.Stacks = s.Stacks[0 : len(s.Stacks)-1]
	}
	return v, nil
}

type Node struct {
	Above *Node
	Below *Node
	Value int
}

func CreateNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

type Stack struct {
	Capacity int
	Top      *Node
	Bottom   *Node
	Size     int
}

func CreateStack(capacity int) *Stack {
	return &Stack{
		Capacity: capacity,
	}
}

func (s *Stack) IsFull() bool {
	return s.Capacity == s.Size
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) Join(above *Node, below *Node) {
	if below != nil {
		below.Above = above
	}
	if above != nil {
		above.Below = below
	}
}

func (s *Stack) Push(v int) bool {
	if s.Size >= s.Capacity {
		return false
	}
	s.Size++
	n := CreateNode(v)
	if s.Size == 1 {
		s.Bottom = n
	}
	s.Join(n, s.Top)
	s.Top = n
	return true
}

func (s *Stack) Pop() int {
	t := s.Top
	s.Top = s.Top.Below
	s.Size--
	return t.Value
}

func (s *Stack) RemoveBottom() int {
	b := s.Bottom
	s.Bottom = s.Bottom.Above
	if s.Bottom != nil {
		s.Bottom.Below = nil
	}
	s.Size--
	return b.Value
}
