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

/*
It's useful to remember in this problem how exactly addition works. Image the problem:

	6 1 2
  + 2 9 5
  --------

First, we add 7 and 5 to get 12. The digit 2 becomes the last digit of the number, and 1 gets carried over to
the next step. Second, we add 1, 1 and 9 to get 11. The 1 becomes the second digit, and the other 1 gets carried
over the final step. Third and finally, we add 1, 6 and 2 to get 9. So, our value becomes 912.

We can mimic this process recursively by adding node by node, carrying over any "excess" data to the next node.
Let's walk through this for the below linked list:

		7 -> 1 -> 6
	+	5 -> 9 -> 2
	---------------

We do the following:
1. We add 7 and 5 first, getting a result of 12. 2 becomes the first node in our linked list, and we "carry" the
   1 to the next sum.
	List: 2 -> ?
2. We then add 1 and 9, as well as the "carry", getting a result of 11. 1 becomes the second element of our linked
list, and we carry the 1 to the next sum.
	list: 2 -> 1 -> ?
3. Finally, we add 6, 2 and our "carry", to get 9. This becomes the final element of our linked list.
	list: 2 -> 1 -> 9.
*/

func AddList(l1 *linked_list.Node[int], l2 *linked_list.Node[int]) *linked_list.Node[int] {
	return SumList(l1, l2, 0)
}

func SumList(l1 *linked_list.Node[int], l2 *linked_list.Node[int], carry int) *linked_list.Node[int] {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	result := linked_list.Node[int]{}
	value := carry
	if l1 != nil {
		value += l1.Data
	}
	if l2 != nil {
		value += l2.Data
	}
	result.Data = value % 10 // second digit of number
	// recurse
	if l1 != nil || l2 != nil {
		var l1Next *linked_list.Node[int]
		if l1 != nil {
			l1Next = l1.Next
		}
		var l2Next *linked_list.Node[int]
		if l2 != nil {
			l2Next = l2.Next
		}
		currentCarry := 0
		if value >= 10 {
			currentCarry = 1
		}
		more := SumList(l1Next, l2Next, currentCarry)
		result.Next = more
	}
	return &result
}

/*
Part B is conceptually the same (recurse, carray the excess), but has some additional complications when it comes
to implementation:

1. One list may be shorter than the other, and we cannot handle this "on the fly". For example, suppose we
   were adding (1->2->3->4) and (5->6->7). We need to know that the 5 should be "matched" with the 2, not the 1.
   We can accomplish this by comparing the lengths of the lists in the beginning and padding the shorter list
   with zeros.

2. In the first pat, successive results were added to the tail (i.e., passed forward). This meant that the recursive
   call would be passed the carry, and would return the result (which is then appended to the tail). In this case,
   however, results are added to the head (i.e., passed backward). The recursive call must return the result, as before,
   as well as the the carry. This is not terribly challenging to implement, but it is more cumbersome. We can solve
   this issue by creating a wrapper class called PartialSum.
*/

type PartialSum struct {
	Sum   *linked_list.Node[int]
	Carry int
}

func AddListB(list1 *linked_list.LinkedList[int], list2 *linked_list.LinkedList[int]) *linked_list.Node[int] {
	len1 := list1.Length
	len2 := list2.Length
	l1 := list1.Head
	l2 := list2.Head

	// pad the shorter list with zeros
	if len1 < len2 {
		l1 = PadList(l1, len2-len1)
	} else {
		l2 = PadList(l2, len1-len2)
	}
	// add list
	sum := AddListHelper(l1, l2)
	// if there was a carry value left over, insert this at the front of the list
	// otherwise, just return the linked list
	if sum.Carry == 0 {
		return sum.Sum
	} else {
		result := InsertBefore(sum.Sum, sum.Carry)
		return result
	}
}

func AddListHelper(l1 *linked_list.Node[int], l2 *linked_list.Node[int]) *PartialSum {
	if l1 == nil && l2 == nil {
		sum := PartialSum{}
		return &sum
	}
	// add smaller digits recursively
	sum := AddListHelper(l1.Next, l2.Next)
	// add carry to current data
	val := sum.Carry + l1.Data + l2.Data
	// insert sum of current digits
	fullResult := InsertBefore(sum.Sum, val%10)
	// return sum so far, and the carry value
	sum.Sum = fullResult
	sum.Carry = val / 10
	return sum
}

func PadList(l *linked_list.Node[int], padding int) *linked_list.Node[int] {
	head := l
	for i := 0; i < padding; i++ {
		head = InsertBefore(head, 0)
	}
	return head
}

func InsertBefore(list *linked_list.Node[int], data int) *linked_list.Node[int] {
	node := linked_list.Node[int]{
		Data: data,
	}
	if list != nil {
		node.Next = list
	}
	return &node
}
