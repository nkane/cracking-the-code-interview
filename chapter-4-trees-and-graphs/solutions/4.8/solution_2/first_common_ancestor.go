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
	Solution #2: With Links to Parents
	Similar to the earlier approach, we could trace p's path upwards and check if any of the nodes cover q. The first
	node that covers q (we already know that every node on this path will cover p) must be the first common ancestor.

	Observe that we don't need to re-check the entire subtree. As we move from node x to its parent y, all the nodes
	under x have already been check for q. Therefore, we only need to check the new nodes "uncovered", which will be
	the ndoes under x's sibling.

	This algorithm takes O(t) time, where t is the size of the subtree for the first common ancestor. In the worse
	case, this will be O(n), where n is the number of nodes in the tree. We can derive this runtime by noticing that
	each nodein that subtree is searched once.
*/

func CommonAncestor[T any](root *TreeNode[T], p *TreeNode[T], q *TreeNode[T]) *TreeNode[T] {
	// check if either node is not in the tree, or if one covers the other
	if !Covers(root, p) || !Covers(root, q) {
		return nil
	} else if Covers(p, q) {
		return p
	} else if Covers(q, p) {
		return q
	}
	// traverse upwards util you find a node that covers q
	siblings := GetSibling(p)
	parent := p.Parent
	for !Covers(siblings, q) {
		siblings = GetSibling(parent)
		parent = parent.Parent
	}
	return parent
}

func Covers[T any](root *TreeNode[T], p *TreeNode[T]) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}
	return Covers(root.Left, p) || Covers(root.Right, p)
}

func GetSibling[T any](node *TreeNode[T]) *TreeNode[T] {
	if node == nil || node.Parent == nil {
		return nil
	}
	parent := node.Parent
	if parent.Left == node {
		return parent.Right
	}
	return parent.Left
}
