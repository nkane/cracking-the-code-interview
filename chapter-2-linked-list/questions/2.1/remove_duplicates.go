package remove_duplciates

import (
	"linked_list"
)

/*
	2.1: Remove Dups: Write code to remove duplicates from an unsorted linked list.

	Follow up
	How would you solve this problem if a temporary buffer is not allowed?
*/

func RemoveDuplciates(l *linked_list.LinkedList) {
	dups := map[int]struct{}{}
	n := l.Head
	idx := 1
	for n != nil {
		if _, ok := dups[n.Data]; ok {
			// duplicate value, remove this node
			l.Remove(idx)
		}
		dups[n.Data] = struct{}{}
		n = n.Next
		idx++
	}
}
