package bst_sequence

import (
	"testing"

	"gotest.tools/assert"
)

func TestBSTSequence(t *testing.T) {
	// Create a sample BST
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	expected := [][]int{
		{2, 1, 3},
		{2, 3, 1},
	}
	// Get all possible arrays that could have led to this tree
	sequences := BSTSequences(root)
	assert.DeepEqual(t, expected, sequences)
}
