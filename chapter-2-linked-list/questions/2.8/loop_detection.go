package loop_detection

import (
	"linked_list"
)

/*
	2.8: Loop Detection: Given a linked list which might contain a loop, implement an algorithm
	that returns the node at the beginning of the loop (if one exists).

	Example:
	Input: a -> b -> c -> d -> e -> c [the same c as earlier]
	Output: c
*/

func LoopDetect(l *linked_list.LinkedList[string]) *linked_list.Node[string] {
	m := map[*linked_list.Node[string]]struct{}{}
	n := l.Head
	for n != nil {
		if _, ok := m[n]; ok {
			return n
		}
		m[n] = struct{}{}
		n = n.Next
	}
	return nil
}
