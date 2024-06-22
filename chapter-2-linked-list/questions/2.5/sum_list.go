package sum_list

import (
	"linked_list"
)

/*
	2.5: Sum lists: You have two numbers represented by a linked list, where each node contains a single
	digit. The digits are stored in reverse order, such that 1's digit is at the head of the list. Write
	a function that adds two numbers and returns the sum as a linked list. (You are not allowed to
	"cheat" and just convert the linked list to an integer).

	Example:
	Input:  (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295
	Output: 2 -> 1 -> 9. That is, 912.

	Follow up:
	Suppose the digits are stored in forward order. Repeat the above problem.
	Example:
	Input:  (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295
	Output: 9 -> 1 -> 2. That is, 912.
*/

func SumList(l1 *linked_list.LinkedList, l2 *linked_list.LinkedList) {
	n1 := l1.Head
	n2 := l2.Head
	if l1.Length > l2.Length {
		n1 = l2.Head
		n2 = l1.Head
	} else if l2.Length > l1.Length {
	}
	for n1 != nil && n2 != nil {
		n2.Data += n1.Data
		div := n2.Data / 10
		modValue := n2.Data % 10
		n2.Data = modValue
		if div > 0 && n2.Next == nil {
			l2.Add(&linked_list.Node{
				Data: div,
			})
		} else {
			if n2.Next == nil {
				n2.Data = modValue
			} else {
				n2.Next.Data += div
			}
		}
		n1 = n1.Next
		n2 = n2.Next
	}
}

func SumListNormal(l1 *linked_list.LinkedList, l2 *linked_list.LinkedList) {
	// n1 := l1.Head
	// n2 := l1.Head
}
