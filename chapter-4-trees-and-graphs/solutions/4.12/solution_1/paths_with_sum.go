package paths_with_sum

/*
	4.1: Paths with Sum
	You are given a binary tree in which each node contains a integer value (which might be positive or negative).
	Design an algorithm to count the number of paths that sum to a given value. The path does not need to start or end
	at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).
*/

/*
	Solution #1: Brute force
	In the brute force approach, we just look at all possible paths. To do this, we traverse to each node. At each
	node, we recursively try all paths downwards, tracking the sum as we go. As soon as we git our target sum, we
	increment the total.

	The runtime is O(N log N) in a balanced tree, in an unbalanced tree it is O(N^2).
*/

type TreeNode struct {
	Data  int
	Right *TreeNode
	Left  *TreeNode
}

func CountPathsWithSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	pathsFromRoot := CountPathsWithSumFromNode(root, targetSum, 0)
	pathsOnLeft := CountPathsWithSum(root.Left, targetSum)
	pathsOnRight := CountPathsWithSum(root.Right, targetSum)
	return pathsFromRoot + pathsOnLeft + pathsOnRight
}

func CountPathsWithSumFromNode(node *TreeNode, targetSum int, currentSum int) int {
	if node == nil {
		return 0
	}
	currentSum += node.Data
	totalPaths := 0
	if currentSum == targetSum {
		totalPaths++
	}
	totalPaths += CountPathsWithSumFromNode(node.Left, targetSum, currentSum)
	totalPaths += CountPathsWithSumFromNode(node.Right, targetSum, currentSum)
	return totalPaths
}
