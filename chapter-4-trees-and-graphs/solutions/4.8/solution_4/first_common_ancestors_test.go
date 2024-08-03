package first_common_ancestor

import (
	"testing"

	"gotest.tools/assert"
)

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
	got := CommonAncestor(&root, root.Left.Right, root.Left.Left.Left)
	assert.DeepEqual(t, root.Left, got)

	root = TreeNode[int]{
		Data: 2,
	}
	root.Right = &TreeNode[int]{
		Data:   30,
		Parent: &root,
	}
	root.Left = &TreeNode[int]{
		Data:   10,
		Parent: &root,
	}
	root.Left.Right = &TreeNode[int]{
		Data:   15,
		Parent: root.Left,
	}
	root.Left.Right.Right = &TreeNode[int]{
		Data:   17,
		Parent: root.Left.Right,
	}
	root.Left.Left = &TreeNode[int]{
		Data:   55,
		Parent: root.Left,
	}
	root.Left.Left.Right = &TreeNode[int]{
		Data:   36,
		Parent: root.Left.Left,
	}
	root.Left.Left.Left = &TreeNode[int]{
		Data:   3,
		Parent: root.Left.Left,
	}

	p := root.Left.Left.Right
	q := root.Left.Right.Right
	got = CommonAncestor(&root, p, q)
	assert.DeepEqual(t, root.Left, got)
}
