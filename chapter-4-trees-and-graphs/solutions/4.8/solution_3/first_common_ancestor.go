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
	Solution #3: Without Links to Parents
	Alternatively, you could follow a chain in which p and q are on the same side. That is, if p and q are both on
	the left of the node, branch left to look for the common ancestor. If they are both on the right, branch right
	to look for thec ommon ancestors. When p and q are no longer on the same side, you must have found the first
	common ancestor.

	This algorithm runs in O(n) time on a balanced tree. This is because covers is called on 2n nodes in the first call
	(n nodes for the left side, and n nodes for the right side). After that, the algorithm branches left or right, at
	which point covers will be called on (2n/2) nodes, then (2n/4), and so on. This results in a runtime of O(n).

	We know at this point that we cannot do better than that in terms of the asymptotic runtime since we need to
	potentionally look at every node in the tree; however, we may be able to improve it by a constant multiple.
*/

func CommonAncestor[T any](root *TreeNode[T], p *TreeNode[T], q *TreeNode[T]) *TreeNode[T] {
	// error check, one node is not in the tree
	if !Covers(root, p) || !Covers(root, q) {
		return nil
	}
	return ancestorHelper(root, p, q)
}

func ancestorHelper[T any](root *TreeNode[T], p *TreeNode[T], q *TreeNode[T]) *TreeNode[T] {
	if root == nil || root == p || root == q {
		return root
	}
	pIsOnLeft := Covers(root.Left, p)
	qIsOnLeft := Covers(root.Left, q)
	if pIsOnLeft != qIsOnLeft {
		// nodes are on different sides
		return root
	}
	childSide := root.Left
	if !pIsOnLeft {
		childSide = root.Right
	}
	return ancestorHelper(childSide, p, q)
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
