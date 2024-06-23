package k_to_last

import (
	"linked_list"
	"log"
)

/*
	2.2: Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
*/

/*
We will approach this problem both recursively and non-recursively. Remember that recursive solutions are often
cleaner but less optimal. For example, in this problem, the recursive implementation is about half the length
of the iterative solution but also takes O(n) space, where n is the number of elements in the linked list.

Note that for this solution, we defined k such that passing k = 1 would return the last element, k = 2 would
return to the second to last element, and so on. It is equally acceptable to define k such that k = 0 would
return the last element.

Solution #1: If linked list size is known

If the size of the linked list is known, then the kth to last element is the (length - k)th element. We can
just iterate through the linked list to find this element. Because this solution is so trivial, we can almost
be sure that this is not what the interviewer intended.

Solution #2: Recursive

This algorithm recurses through the linked list. When it hits the end, the method passes back a counter set to 0,
Each parent call adds 1 to the counter. When the counter equals k, we know we have reached the kth to last element
of the linked list.

Implement this is short and sweet, provided we have a way of "passing back" an integer value through the stack.
Unfortunately, we can't pass back a node and a counter using normal return statements. So how do we handle this?

Approach A Don't Return the Element

One way to do this is to change the problem to simply printing the kth to last element. Then, we can pass back the
value of the counter simply through return values.
*/

func PrintKthToLast(head *linked_list.Node[int], k int) int {
	if head == nil {
		return 0
	}
	index := PrintKthToLast(head.Next, k) + 1
	if index == k {
		log.Printf("%dth to the last node is %d", k, head.Data)
	}
	return index
}

/*
Approach B: Use C++
A second way to solve this is to use C++ to pass values by reference. This allows us to return the node
value, but also update the counter by passing a pointer to it.

node* nthToLast(node *head, int k, int& i) {
	if (head == NULL) {
		return NULL;
	}
	node *nd = nthToLast(head->next, k, i);
	i = i + 1;
	if (i == k) {
		return head;
	}
	return nd;
}

node * nthToLast(node *head, int k) {
	int i = 0;
	return nthToLast(head, k, i);
}
*/

/*
Approach C: Create a Wrapper Class
We described earlier that the issue was that we couldn't simultaneously return a counter and an index. If we wrap
the counter value with a simple class (or even a single element array), we can mimic passing by reference.

Each of these recursive solutions takes O(n) space due to the recursive calls.
*/

type Index struct {
	Value int
}

func KthToLastWrapper(head *linked_list.Node[int], k int) *linked_list.Node[int] {
	idx := Index{}
	return KthToLastWrapperHelper(head, k, &idx)
}

func KthToLastWrapperHelper(head *linked_list.Node[int], k int, idx *Index) *linked_list.Node[int] {
	if head == nil {
		return nil
	}
	node := KthToLastWrapperHelper(head.Next, k, idx)
	idx.Value = idx.Value + 1
	if idx.Value == k {
		return head
	}
	return node
}

/*
Solution #3: Iterative

A more optimal, but less straightforward, solution is to implement this iteratively. We can use two pointers,
p1 and p2. We place them k nodes apart from the linked list by putting p2 at the beginning and moving p1 k
nodes into the list. Then, we move the at the same pace, p1 will hit the end of the linked list after LENGTH - k
steps. At that point p2 will be LENGTH - k nodes into the list, or k nodes from the end.

This algorithm takes O(n) time and O(1) space.
*/

func NthToLast(head *linked_list.Node[int], k int) *linked_list.Node[int] {
	p1 := head
	p2 := head
	// move p1 k nodes into the list
	for i := 0; i < k; i++ {
		if p1 == nil { // out of bounds
			return nil
		}
		p1 = p1.Next
	}
	// move them at the same pace, when p1 hits the end, p2 will be at the right
	// element
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
