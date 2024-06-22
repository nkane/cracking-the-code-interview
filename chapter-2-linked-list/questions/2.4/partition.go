package partition

import (
	"linked_list"
)

/*
	2.4: Partition: Write code to partition a linked list around a value x, such that all nodes less than x
	come before all nodes greater than or equal to x. (IMPORTANT): The partition element x can appear anywhere
	in the "right partition"; it does not need to appear between the left and right partitions.
	The additional spacing in the example below indicates the partition. Yes, the output below is one
	of many valid outputs!

	Example
	Input:	3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition = 5]
	Output: 3 -> 1 -> 2	-> 10 -> 5 -> 5 -> 8
*/

func Partition(l *linked_list.LinkedList[int], x int) {
	var beforeStart *linked_list.Node[int]
	var beforeEnd *linked_list.Node[int]
	var afterStart *linked_list.Node[int]
	var afterEnd *linked_list.Node[int]

	node := l.Head
	// partition list
	for node != nil {
		next := node.Next
		node.Next = nil
		if node.Data < x {
			// insert node into end of before list
			if beforeStart == nil {
				beforeStart = node
				beforeEnd = beforeStart
			} else {
				beforeEnd.Next = node
				beforeEnd = node
			}
		} else {
			// insert node into end of after list
			if afterStart == nil {
				afterStart = node
				afterEnd = afterStart
			} else {
				afterEnd.Next = node
				afterEnd = node
			}
		}
		node = next
	}
	if beforeStart == nil {
		l.Head = afterStart
		return
	}
	// merge before list and after list
	beforeEnd.Next = afterStart
	l.Head = beforeStart
}
