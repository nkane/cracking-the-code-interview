package minimal_tree

import (
	"testing"
	"tree"
)

func TestMinimalBST(t *testing.T) {
	tree := tree.BuildTree()
	tree.DrawSVG("test-start")
	tree.Root = MinimalBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	tree.DrawSVG("test-end")
}
