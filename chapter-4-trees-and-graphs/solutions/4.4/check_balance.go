package check_balance

import (
	"math"
	"tree"

	"github.com/golang-collections/collections/queue"
)

/*
	4.4: Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of this
	question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node
	never differs by more than one.
*/

func CheckBalance(root *tree.Node) bool {
	leftDepth := GetDepth(root.Left)
	rightDepth := GetDepth(root.Right)
	return leftDepth == rightDepth
}

func GetDepth(root *tree.Node) int {
	depth := 0
	q := queue.New()
	depth++
	q.Enqueue(root)
	for q.Len() > 0 {
		node, ok := q.Dequeue().(*tree.Node)
		if !ok || node == nil {
			continue
		}
		if node.Left != nil || node.Right != nil {
			q.Enqueue(node.Left)
			q.Enqueue(node.Right)
			depth++
		}
	}
	return depth
}

/*
In this question, we've been fortunate enough to be told exactly what balanced means: that for each node, the two
subtrees differ in height by no more than one. We can implement a solution based on this definition. We can simply
recurse through the entire tree, and for each node, and for each node, compute the heights of its subtree.
*/

func getHeight(root *tree.Node) int {
	if root == nil {
		return -1
	}
	m := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if m < rightHeight {
		m = rightHeight
	}
	return m + 1
}

func isBalanced(root *tree.Node) bool {
	if root == nil {
		return true
	}
	heightDiff := getHeight(root.Left) - getHeight(root.Right)
	if heightDiff < 0 {
		heightDiff *= -1
	}
	if heightDiff > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

/*
Although this works, it's not very efficient. On each node, we recurse through its entire subtree. This  means
that getHeight is called repeatedly on the same node. The algorithm is O(N log N) since each node is "touched"
once per node above it.

We need to cut out some of the calls to getHeight.

If we inspect this method, we may notice that getHeight could actually check if the tree is balanced at the same time
as it's checking heights. What do we do when we discover that the subtree isn't balanced? Just return an error code.

This improved algorithm works by checking the height of each subtree as we recurse down from the root. On each node,
we recursively get the heights of the left and right subtrees through the checkHeight method. If the subtree is
balanced, then checkHeight will return the actual height of the subtree. If the subtree is not balanced, then
checkHeight will return an error code. We will immediately break and return an error code from the current call.

What do do we use for an error code? The height of a null tree is generally defined to be -1, so that's not a great
idea for an error code. Instead, we'll use min int.

This code runs in O(N) time and O(H) space, where H is the height of the tree.
*/

func checkHeightRevised(root *tree.Node) int {
	if root == nil {
		return -1
	}
	leftHeight := checkHeightRevised(root.Left)
	if leftHeight == math.MinInt {
		return math.MinInt
	}
	rightHeight := checkHeightRevised(root.Right)
	if rightHeight == math.MinInt {
		return math.MinInt
	}
	heightDiff := leftHeight - rightHeight
	if heightDiff < 0 {
		heightDiff *= -1
	}
	if heightDiff > 1 {
		return math.MinInt
	}
	max := leftHeight
	if max < rightHeight {
		max = rightHeight
	}
	return max
}
