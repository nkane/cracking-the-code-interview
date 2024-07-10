package tree

import "testing"

func TestTree(t *testing.T) {
	tree := BuildTree()
	tree.DrawSVG("test")
}
