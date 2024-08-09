package paths_with_sum

import (
	"testing"

	"gotest.tools/assert"
)

func TestPathsWithSum(t *testing.T) {
	root := TreeNode{
		Data: 10,
		Right: &TreeNode{
			Data: -3,
			Right: &TreeNode{
				Data: 11,
			},
		},
		Left: &TreeNode{
			Data: 5,
			Right: &TreeNode{
				Data: 2,
				Right: &TreeNode{
					Data: 1,
				},
				Left: &TreeNode{
					Data: 3,
					Right: &TreeNode{
						Data: -2,
					},
					Left: &TreeNode{
						Data: 3,
					},
				},
			},
		},
	}
	got := CountPathsWithSum(&root, 8)
	assert.Assert(t, got == 4, "got %+v, expected: %+v", got, 3)
}
