package paths_with_sum

/*
	4.1: Paths with Sum
	You are given a binary tree in which each node contains a integer value (which might be positive or negative).
	Design an algorithm to count the number of paths that sum to a given value. The path does not need to start or end
	at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).
*/

/*
	Solution #2: Optimized

	The runtime for this algorithm is O(N), where N is the number of nodes in the tree. In a balanced tree, the space
	complexity is O(log N) due to the hash table, the space complexity can grow to O(N) in an unbalanced tree.
*/

type TreeNode struct {
	Data  int
	Right *TreeNode
	Left  *TreeNode
}

func CountPathsWithSum(root *TreeNode, targetSum int) int {
	return countPathsWithSum(root, targetSum, 0, make(map[int]int))
}

// countPathsWithSum calculates the number of paths with a given sum in a subtree
func countPathsWithSum(node *TreeNode, targetSum int, runningSum int, pathCount map[int]int) int {
	if node == nil {
		return 0
	}

	runningSum += node.Data
	// Calculate the number of valid paths ending at the current node
	sum := runningSum - targetSum
	totalPaths := pathCount[sum]

	// If runningSum equals targetSum, then one additional path starts at the root
	if runningSum == targetSum {
		totalPaths++
	}

	// Add runningSum to pathCounts
	incrementHashTable(pathCount, runningSum, 1)

	// Count paths with sum on the left and right
	totalPaths += countPathsWithSum(node.Left, targetSum, runningSum, pathCount)
	totalPaths += countPathsWithSum(node.Right, targetSum, runningSum, pathCount)

	// Remove runningSum from pathCounts
	incrementHashTable(pathCount, runningSum, -1)

	return totalPaths
}

// incrementHashTable updates the map with the given key and delta
func incrementHashTable(hashTable map[int]int, key int, delta int) {
	newCount := hashTable[key] + delta
	if newCount == 0 {
		delete(hashTable, key)
	} else {
		hashTable[key] = newCount
	}
}
