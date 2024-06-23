package linked_list

import (
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	Next *Node[T]
	Data T
}

type LinkedList[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func (ll *LinkedList[T]) AddHead(n *Node[T]) {
	n.Next = ll.Head
	ll.Head = n
}

func (ll *LinkedList[T]) Add(n *Node[T]) {
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}
	ll.Length++
}

func (ll *LinkedList[T]) Find(data T) *Node[T] {
	n := ll.Head
	for n != nil {
		if n.Data == data {
			return n
		}
		n = n.Next
	}
	return nil
}

func (ll *LinkedList[T]) Remove(k int) {
	if k > ll.Length {
		return
	}
	ll.Length--
	if k == 1 {
		// remove the head
		tmpN := ll.Head
		ll.Head = ll.Head.Next
		tmpN.Next = nil
		return
	}
	n := ll.Head
	prev := ll.Head
	for i := 1; i < k; i++ {
		prev = n
		n = n.Next
	}
	prev.Next = n.Next
	if k == ll.Length {
		ll.Tail = prev
	}
	n.Next = nil
}

func (ll *LinkedList[T]) String() string {
	sb := strings.Builder{}
	n := ll.Head
	i := 1
	for n != nil {
		sb.WriteString(fmt.Sprintf("[pos: %d | data: %v] -> ", i, n.Data))
		n = n.Next
		i++
	}
	return sb.String()
}

func (ll *LinkedList[T]) StringSimple() string {
	sb := strings.Builder{}
	n := ll.Head
	for n != nil {
		sb.WriteString(fmt.Sprintf("[data: %v] -> ", n.Data))
		n = n.Next
	}
	return sb.String()
}
