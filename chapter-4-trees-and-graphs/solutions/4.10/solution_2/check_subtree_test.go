package check_subtree

import (
	"testing"

	"gotest.tools/assert"
)

func TestContainsTree(t *testing.T) {
	tests := []struct {
		T1        *TreeNode
		T2        *TreeNode
		Exepected bool
	}{
		{
			T1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			T2: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Exepected: true,
		},
		{
			T1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			T2: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Exepected: false,
		},
	}
	for _, test := range tests {
		got := containsTree(test.T1, test.T2)
		assert.Assert(t, test.Exepected == got)
	}
}
