package successor

import (
	"tree"
)

/*
	4.6: Successor: Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a
	binary search tree. You may assume that each node has a link to its parent.
*/

/*
Recall that an in-order traversal traverses the left subtree, then the current node, then the right subtree. To
approach this problem, we need to think very, very carefully about what happens.

Let's suppose we have a hypothetical node. We know that the order goes left subtree, then current side, then right
subtree. So, the next node we visit should be on the right side.

But which node on the right subtree? It should be the first node we'd visit if we were doing an in-order
traversal of that subtree. This means that it should be the leftmost node on the right subtree. Easy enough!

But what if the nodes doesn't have a right subtree? This is where it gets a bit trickier.

If a node n doesn't have a right subtree, then we are done traversing n's subtree. We need to pick up where we
left off with n's parent, which we'll call q.

If n was to the left of q, then the next node we should traverse should be q (again, since left -> current -> right).

If n were to the right of q, then we have fully traversed q's subtree as well. We need to traverse upwards from q
until we find a node x that we have not fully traversed. How do we know that we have not fully traversed a node x?
We know we have hit this case when we move from a left node to its parent. The left ndoe is full traversed, but its
parent is not.
*/

func inOrderSuccesor(n *tree.Node) *tree.Node {
	if n == nil {
		return nil
	}
	// found right childre, return left most node of right subtree
	if n.Right != nil {
		return leftMostChild(n.Right)
	} else {
		q := n
		x := q.Parent
		// go up until we're on left instead of right
		for x != nil && x.Left != q {
			q = x
			x = x.Parent
		}
		return x
	}
}

func leftMostChild(n *tree.Node) *tree.Node {
	if n == nil {
		return nil
	}
	for n.Left != nil {
		n = n.Left
	}
	return n
}
