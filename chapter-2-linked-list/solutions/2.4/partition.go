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

/*
If this were an array, we would need to be careful about how we shifted elements. Array shifts are very expensive.

However, in a linked list, the situation is much easier. Rather than shifting and swapping elements, we can actually
create two different linked list: one for the elements less than x, and one for elements greater than or equal to x.

We iterate through the linked lsit, inserting elements into our before list and our after list. Once we reach the end
of the linked list and have completed this splitting, we merge the two lists.

This approach is mostly "stable" in that elements stay in their original order, other than the necessary movement
around the partition.
*/

func Partition(node *linked_list.Node[int], x int) *linked_list.Node[int] {
	var beforeStart *linked_list.Node[int]
	var beforeEnd *linked_list.Node[int]
	var afterStart *linked_list.Node[int]
	var afterEnd *linked_list.Node[int]

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
		return afterStart
	}
	// merge before list and after list
	beforeEnd.Next = afterStart
	return beforeStart
}

/*
If it bugs you to keep around four different variables for tracking two linked list, you're not alone. We can make this
code a bit shorter.

If we don't care about making elements of the list "stable" (which there's no obligation to, since the interviewer
hasn't specified that), then we can instead rearrange the elements by growing the list at the head and tail.

In this approach, we start a "new" list (using the existing nodes). Elements bigger than the pivot element are put at
the the tail and elements smaller are put at the head. Each time we insert and element, we update either the head or
tail.
*/

func PartitionSimple(node *linked_list.Node[int], x int) *linked_list.Node[int] {
	head := node
	tail := node
	for node != nil {
		next := node.Next
		if node.Data < x {
			// insert node at head
			node.Next = head
			head = node
		} else {
			// insert at tail
			tail.Next = node
			tail = node
		}
		node = next
	}
	tail.Next = nil
	// the head has changed, so we need to return it to the user
	return head
}
