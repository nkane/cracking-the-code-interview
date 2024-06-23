package remove_duplicates

import (
	"linked_list"
)

/*
	2.1: Remove Dups: Write code to remove duplicates from an unsorted linked list.

	Follow up
	How would you solve this problem if a temporary buffer is not allowed?
*/

/*
In order to remove duplicates from a linked list, we need to be able to track duplciates. A simple hash table
will work well here.

In the below solution, we simply iterate through the linked list, adding each element to a hash table. When we discover
a duplicate element, we remove the element and continue iterating. We can do this all in one pass since we are using a
linked list.

This solution takes O(N) time, where N is the number of elements in the linked list.
*/

func DeleteDuplicatesWithMap(n *linked_list.Node[int]) {
	set := map[int]struct{}{}
	var previous *linked_list.Node[int]
	for n != nil {
		if _, ok := set[n.Data]; ok {
			previous.Next = n.Next
		} else {
			set[n.Data] = struct{}{}
			previous = n
		}
		n = n.Next
	}
}

/*
Follow up: No Buffer Allowed

If we don't have a buffer, we can iterate with two pointers: current which iterates through the linked list,
and runner which checks al subsequent nodes for duplicates.

The code runs in O(1) space, but O(N^2) time.
*/
func DeleteDuplicatesWithRunner(head *linked_list.Node[int]) {
	current := head
	for current != nil {
		// remove all future notes that have the same value
		runner := current
		for runner.Next != nil {
			if runner.Next.Data == current.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
}
