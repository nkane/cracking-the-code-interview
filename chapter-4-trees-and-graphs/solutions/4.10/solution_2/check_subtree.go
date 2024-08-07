package check_subtree

/*
	4.10: Check Subtree: T1 and T2 are two very large binary trees, with T1 much bigger than T2. Create an algorithm
	to determine if T2 is a subtree of T1.

	A tree T2 is a subtree of T1 if there exists a node n in T1 such taht the subtree of n is identical to T2. That is,
	if you cut off the tree at node n, the two trees would be identical.
*/

/*
	An altertnative approach is to search through the larger tree, T1. Each time a node in T1 matches the root of T2,
	call matchTree. The matchTree method will compare the two subtrees to see if they are identical.

	Analyzing the runtime is somewhat complex, a naive answer would be to say that it is O(nm) time, where n is the
	number of nodes in T1 and m is the number of nodes in T2. While this is technically correct, a little more throught
	can produce a tigher bound.

	We do not actually call matchTree on every node in T1. Rather, we call it k times, where k is the number of
	occurences of T2's root in T1. The runtime is closer to O(n + km).

	In fact, even that overstates the runtime. Even if the root were identical, we exist matchTree when we find a
	difference between T1 and T2. We therefore probably do not actually look at m nodes on each call of matchTree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func containsTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	return subTree(t1, t2)
}

func subTree(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil {
		return false
	} else if r1.Val == r2.Val && matchTree(r1, r2) {
		return true
	}
	return subTree(r1.Left, r2) || subTree(r1.Right, r2)
}

func matchTree(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	} else if r1.Val != r2.Val {
		return false
	} else {
		return matchTree(r1.Left, r2.Left) && matchTree(r1.Right, r2.Right)
	}
}
