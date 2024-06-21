package k_last

import (
	"linked_list"
)

/*
	2.2: Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
*/

func KLast(l *linked_list.LinkedList, k int) *linked_list.Node {
	idx := l.Length - k
	n := l.Head
	for idx > 1 {
		n = n.Next
		idx--
	}
	return n
}
