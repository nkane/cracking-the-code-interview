package first_common_ancestor

import (
	"testing"

	"gotest.tools/assert"
)

func TestDepth(t *testing.T) {
	root := TreeNode[int]{
		Data: 1,
	}
	root.Right = &TreeNode[int]{
		Data:   5,
		Parent: &root,
	}

	expected := 2
	got := Depth(root.Right)
	assert.Assert(t, expected == got)
}

func TestCommonAncestor(t *testing.T) {
	root := TreeNode[int]{
		Data: 1,
	}
	root.Right = &TreeNode[int]{
		Data:   6,
		Parent: &root,
	}
	root.Left = &TreeNode[int]{
		Data:   3,
		Parent: &root,
	}
	root.Left.Right = &TreeNode[int]{
		Data:   10,
		Parent: root.Left,
	}
	root.Left.Left = &TreeNode[int]{
		Data:   9,
		Parent: root.Left,
	}
	root.Left.Left.Left = &TreeNode[int]{
		Data:   7,
		Parent: root.Left.Left,
	}
	got := CommonAncestor(root.Left.Right, root.Left.Left.Left)
	assert.DeepEqual(t, root.Left, got)
}
