package delete_middle_node

import (
	"linked_list"
)

/*
	2.3: Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but the first and
	last node, not necessarily the exact middle) of a singly linked list, given only access to that node.

	Example
	Input: the node c from the linked list a -> b -> c -> d -> e -> f
	Result: nothing is returned, but the new linked list looks like a -> b -> d -> e -> f
*/

/*
Solution

In this problem, you are not given access to the head of the linked list. You only have access to that node. The
solution is simply to copy the data from the next node over to the current node, and then to delete the next node.

Note that this problem cannot be solved if the node to be deleted is the last node in the linked list. That's okay,
your interviewer wants you to point that out, and to discuss how to handle thsi case. You could, for example, consider
marking the node as dummy.
*/

func DeleteMiddleNode(n *linked_list.Node[int]) bool {
	if n == nil || n.Next == nil {
		return false
	}
	next := n.Next
	n.Data = next.Data
	n.Next = next.Next
	return true
}
