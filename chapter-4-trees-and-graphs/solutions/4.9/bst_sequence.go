package bst_sequence

/*
	4.9: BST Sequence: A binary search tree was created by traversing through an array from left to right
	and inserting each element. Given a binary search tree with distinct elements, print all possible arrays
	that could have led to this tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to merge two arrays in all possible ways
func WeaveLists(first, second, prefix []int, results *[][]int) {
	if len(first) == 0 || len(second) == 0 {
		// Add remaining elements to prefix and store result
		result := append([]int{}, prefix...)
		result = append(result, first...)
		result = append(result, second...)
		*results = append(*results, result)
		return
	}

	// Recurse with head of first added to prefix, removing the head will be in subsequent calls
	prefix = append(prefix, first[0])
	WeaveLists(first[1:], second, prefix, results)
	prefix = prefix[:len(prefix)-1]

	// Recurse with head of second added to prefix, removing the head will be in subsequent calls
	prefix = append(prefix, second[0])
	WeaveLists(first, second[1:], prefix, results)
	prefix = prefix[:len(prefix)-1]
}

// Function to generate all sequences for a given BST
func BSTSequences(node *TreeNode) [][]int {
	if node == nil {
		return [][]int{{}}
	}

	prefix := []int{node.Val}

	// Recurse on left and right subtrees
	leftSeq := BSTSequences(node.Left)
	rightSeq := BSTSequences(node.Right)

	results := [][]int{}

	// Weave together each list from the left and right side
	for _, left := range leftSeq {
		for _, right := range rightSeq {
			weaved := [][]int{}
			WeaveLists(left, right, prefix, &weaved)
			results = append(results, weaved...)
		}
	}

	return results
}
