package first_common_ancestor

/*
	4.8: First Common Ancestor
	Design an algorithm and write code to find the first common ancestor of two nodes in a binary tree. Avoid
	storing additional nodes in a data structure.
	NOTE: this is not necessarily a binary search tree.
*/

/*
	If this were a binary search tree, we could modify the find operation for the two nodes and see where
	the paths diverge. Unfortunately, this is not a binary search tree, so we must try other approaches.

	Let's assume we're looking for the common ancestors of nodes p and q, one question to ask here is if each
	node in our tree has a link to its parents.
*/

type TreeNode[T any] struct {
	Data   T
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Parent *TreeNode[T]
}

/*
	Solution #1: With Links to Parents
	If each node has a link  to its parent, we could trace p and q's paths up until they intersect. This is essentially
	the same problem as question 2.7 which find the intersection of two linked list. The "linked list" in this case is
	the path from each node up to the root.

	This solution is O(P+D) time, where P is the number of projects and D is the number of dependency pairs.

	By the way, this problem is called topological sort, linearly ordering the vertices in a graph such that for every
	edge (a, b), a appears before b in the linear order.
*/

func CommonAncestor[T any](p *TreeNode[T], q *TreeNode[T]) *TreeNode[T] {
	delta := Depth(p) - Depth(q)
	first := p
	second := q
	if delta > 0 {
		first = q
		second = p
	}
	if delta < 0 {
		delta = -delta
	}
	second = GoUpBy(second, delta)
	// find where paths intersect
	for first != second && first != nil && second != nil {
		first = first.Parent
		second = second.Parent
	}
	if first == nil || second == nil {
		return nil
	}
	return first
}

func GoUpBy[T any](node *TreeNode[T], delta int) *TreeNode[T] {
	for delta > 0 && node != nil {
		node = node.Parent
		delta--
	}
	return node
}

func Depth[T any](node *TreeNode[T]) int {
	depth := 0
	for node != nil {
		node = node.Parent
		depth++
	}
	return depth
}
