package stack

import "errors"

type StackNode[T any] struct {
	Data T
	Next *StackNode[T]
}

type Stack[T any] struct {
	Top *StackNode[T]
}

func (stack *Stack[T]) Pop() (*T, error) {
	if stack.Top == nil {
		return nil, errors.New("invalid pop call")
	}
	item := stack.Top.Data
	stack.Top = stack.Top.Next
	return &item, nil
}

func (stack *Stack[T]) Push(item T) {
	t := StackNode[T]{
		Data: item,
		Next: stack.Top,
	}
	stack.Top = &t
}

func (stack *Stack[T]) Peek() (*T, error) {
	if stack.Top == nil {
		return nil, errors.New("stack is empty")
	}
	return &stack.Top.Data, nil
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.Top == nil
}
