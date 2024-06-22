package linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (ll *LinkedList) Add(n *Node) {
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}
	ll.Length++
}

func (ll *LinkedList) Find(data int) *Node {
	n := ll.Head
	for n != nil {
		if n.Data == data {
			return n
		}
		n = n.Next
	}
	return nil
}

func (ll *LinkedList) Remove(k int) {
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

func (ll *LinkedList) String() string {
	sb := strings.Builder{}
	n := ll.Head
	i := 1
	for n != nil {
		sb.WriteString(fmt.Sprintf("[pos: %d | data: %d] -> ", i, n.Data))
		n = n.Next
		i++
	}
	return sb.String()
}

func (ll *LinkedList) StringSimple() string {
	sb := strings.Builder{}
	n := ll.Head
	for n != nil {
		sb.WriteString(fmt.Sprintf("[data: %d] -> ", n.Data))
		n = n.Next
	}
	return sb.String()
}
