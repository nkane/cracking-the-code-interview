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
	Solution  #4: Optimized
	Although Solution #3 is optimial in its runtime, we may recognize that there is still some inefficiency in how it
	operates. Specifically, covers searches all nodes under root for p and q, including the nodes in each subtree
	(root.left and root.right). Then, it picks one of those subtrees and searches all of its nodes. Each subtree is
	searched over and over again.
*/

type Result[T any] struct {
	Node       *TreeNode[T]
	IsAncestor bool
}

func CommonAncestor[T any](root *TreeNode[T], p *TreeNode[T], q *TreeNode[T]) *TreeNode[T] {
	r := ancestorHelper(root, p, q)
	if r.IsAncestor {
		return r.Node
	}
	return nil
}

func ancestorHelper[T any](root *TreeNode[T], p *TreeNode[T], q *TreeNode[T]) *Result[T] {
	if root == nil {
		return &Result[T]{
			Node:       nil,
			IsAncestor: false,
		}
	}
	if root == p && root == q {
		return &Result[T]{
			Node:       root,
			IsAncestor: true,
		}
	}
	rx := ancestorHelper(root.Left, p, q)
	if rx.IsAncestor {
		// found common ancestor
		return rx
	}
	ry := ancestorHelper(root.Right, p, q)
	if ry.IsAncestor {
		// found common ancestor
		return ry
	}
	if rx.Node != nil && ry.Node != nil {
		// this is the common ancestor
		return &Result[T]{
			Node:       root,
			IsAncestor: true,
		}
	} else if root == p || root == q {
		// if we're currently at p or q, and we also found one of those nodes in a subtree, then this is truly
		// and ancestor and the flag should be true.
		isAncestor := false
		if rx.Node != nil || ry.Node != nil {
			isAncestor = true
		}
		return &Result[T]{
			Node:       root,
			IsAncestor: isAncestor,
		}
	} else {
		var node *TreeNode[T]
		if rx.Node != nil {
			node = rx.Node
		} else if ry.Node != nil {
			node = ry.Node
		}
		return &Result[T]{
			Node:       node,
			IsAncestor: false,
		}
	}
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
