package intersection

import (
	"linked_list"
)

/*
	2.7: Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the
	intersecting node. Note that the intersection is defined based on reference, not value. That is,
	if the kth node of the first linked list is the exact same node (by reference) as the jth node
	of the second linked list, then they are intersecting.
*/

func Intersecting(l1 *linked_list.LinkedList[int], l2 *linked_list.LinkedList[int]) *linked_list.Node[int] {
	n1 := l1.Head
	m := map[*linked_list.Node[int]]*linked_list.Node[int]{}
	for n1 != nil {
		m[n1] = n1
		n1 = n1.Next
	}
	n2 := l2.Head
	for n2 != nil {
		if v, ok := m[n2]; ok {
			return v
		}
		n2 = n2.Next
	}
	return nil
}
