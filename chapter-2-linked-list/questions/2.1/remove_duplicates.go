package remove_duplciates

import (
	"fmt"
	"strings"
)

/*
	2.1: Remove Dups: Write code to remove duplicates from an unsorted linked list.

	Follow up
	How would you solve this problem if a temporary buffer is not allowed?
*/

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
	for i := 1; i < k; k++ {
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
