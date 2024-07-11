package check_balance

import (
	"testing"
	"tree"

	"gotest.tools/v3/assert"
)

func TestCheckBalance(t *testing.T) {
	tree := tree.BuildTree()
	tree.DrawSVG("test-tree")
	got := CheckBalance(tree.Root)
	assert.Assert(t, got == false)
}

func TestCheckBalanceBookSolution(t *testing.T) {
	tree := tree.BuildTree()
	tree.DrawSVG("test-tree")
	got := isBalanced(tree.Root)
	assert.Assert(t, got == false)
}

func TestCheckBalanceBookSolutionRevised(t *testing.T) {
	tree := tree.BuildTree()
	tree.DrawSVG("test-tree")
	got := isBalanced(tree.Root)
	assert.Assert(t, got == false)
}
