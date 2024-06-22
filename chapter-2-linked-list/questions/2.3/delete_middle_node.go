package delete_middle_node

import (
	"linked_list"
)

/*
	2.3: Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but the first and
	last node, not necessarily the exact middle) of a singly linked list, given only access to that node.

	Example
	Input: the node c from the linked list a -> b -> c -> d -> e -> f
	Result: nothing is returned, but the new linked list looks like a -> b -> d -> e -> f
*/

func DeleteMiddleNode[T comparable](l *linked_list.LinkedList[T]) {
	if l.Length <= 2 {
		return
	}
	isEven := (l.Length % 2) == 0
	modifier := 0
	if !isEven {
		modifier = 1
	}
	halfLength := (l.Length / 2) + modifier
	n := l.Head
	idx := 1
	for n != nil && idx < halfLength {
		n = n.Next
		idx++
	}
	l.Remove(idx)
}
