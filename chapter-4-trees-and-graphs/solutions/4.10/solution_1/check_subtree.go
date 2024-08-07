package check_subtree

import (
	"strconv"
	"strings"
)

/*
	4.10: Check Subtree: T1 and T2 are two very large binary trees, with T1 much bigger than T2. Create an algorithm
	to determine if T2 is a subtree of T1.

	A tree T2 is a subtree of T1 if there exists a node n in T1 such taht the subtree of n is identical to T2. That is,
	if you cut off the tree at node n, the two trees would be identical.
*/

/*
	This approach takes O(n * m) time (due to the runtime of Index method), where n and m are the number of nodes in T1
	and T2, respectively. The space complexity is O(n + m).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func containsTree(t1 *TreeNode, t2 *TreeNode) bool {
	sb1 := strings.Builder{}
	sb2 := strings.Builder{}
	getOrderString(t1, &sb1)
	getOrderString(t2, &sb2)
	str1 := sb1.String()
	str2 := sb2.String()
	if strings.Index(str1, str2) > 0 {
		return true
	}
	return false
}

func getOrderString(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteRune('X')
		return
	}
	sb.WriteString(strconv.Itoa(node.Val))
	sb.WriteRune(' ')
	getOrderString(node.Left, sb)
	getOrderString(node.Right, sb)
}
