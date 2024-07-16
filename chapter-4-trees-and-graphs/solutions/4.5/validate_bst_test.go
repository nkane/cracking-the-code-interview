package validate_bst

import (
	"testing"
	"tree"

	"gotest.tools/assert"
)

func TestCheckBSTAlternative(t *testing.T) {
	goodTree := tree.Tree{
		Root: &tree.Node{
			Data: 10,
			Left: &tree.Node{
				Data: 8,
			},
			Right: &tree.Node{
				Data: 16,
			},
		},
	}
	badTree := tree.Tree{
		Root: &tree.Node{
			Data: 10,
			Left: &tree.Node{
				Data: 20,
			},
			Right: &tree.Node{
				Data: 16,
			},
		},
	}
	assert.Assert(t, checkBST_alternative(goodTree.Root) == true)
	assert.Assert(t, checkBST_alternative(badTree.Root) == false)
	assert.Assert(t, checkBST_2(goodTree.Root) == true)
	assert.Assert(t, checkBST_2(badTree.Root) == false)
}
