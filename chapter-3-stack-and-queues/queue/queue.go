package queue

import "errors"

type QueueNode[T any] struct {
	Data T
	Next *QueueNode[T]
}

type Queue[T any] struct {
	First *QueueNode[T]
	Last  *QueueNode[T]
}

func (queue *Queue[T]) Add(item T) {
	t := QueueNode[T]{
		Data: item,
	}
	if queue.Last != nil {
		queue.Last.Next = &t
	}
	queue.Last = &t
	if queue.First == nil {
		queue.First = queue.Last
	}
}

func (queue *Queue[T]) Remove() (*T, error) {
	if queue.First == nil {
		return nil, errors.New("queue empty")
	}
	data := queue.First.Data
	queue.First = queue.First.Next
	if queue.First == nil {
		queue.Last = nil
	}
	return &data, nil
}

func (queue *Queue[T]) Peek() (*T, error) {
	if queue.First == nil {
		return nil, errors.New("queue empty")
	}
	return &queue.First.Data, nil
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.First == nil
}
