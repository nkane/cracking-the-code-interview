package remove_duplciates

import (
	"linked_list"
)

/*
	2.1: Remove Dups: Write code to remove duplicates from an unsorted linked list.

	Follow up
	How would you solve this problem if a temporary buffer is not allowed?
*/

type RemoveDuplciates func(l *linked_list.LinkedList)

func RemoveDuplciatesWithMap(l *linked_list.LinkedList) {
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

func RemoveDuplciatesNoMap(l *linked_list.LinkedList) {
	behind := l.Head
	ahead := l.Head
	idxBehind := 1
	idxAhead := 2
	for behind != nil {
		ahead = behind.Next
		for ahead != nil {
			if behind.Data == ahead.Data {
				l.Remove(idxAhead)
			}
			ahead = ahead.Next
			idxAhead++
		}
		behind = behind.Next
		idxBehind++
		idxAhead = idxBehind + 1
	}
}
