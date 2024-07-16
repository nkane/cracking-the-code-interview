package successor

import (
	"testing"
	"tree"

	"gotest.tools/assert"
)

func TestInOrderSuccessor(t *testing.T) {
	root := tree.Node{
		Data: 20,
	}
	left := tree.Node{
		Data:   10,
		Parent: &root,
	}
	root.Left = &left
	leftSub := tree.Node{
		Data:   5,
		Parent: &left,
	}
	left.Left = &leftSub
	leftRightSub := tree.Node{
		Data:   15,
		Parent: &left,
	}
	left.Right = &leftRightSub
	leftRightRightSub := tree.Node{
		Data:   17,
		Parent: &leftRightSub,
	}
	leftRightSub.Right = &leftRightRightSub

	right := tree.Node{
		Data:   30,
		Parent: &root,
	}
	rightSub := tree.Node{
		Data:   35,
		Parent: &right,
	}
	right.Right = &rightSub

	root.Right = &right
	got := inOrderSuccesor(root.Left)
	assert.Assert(t, got != nil)
	assert.Assert(t, got.Data == 15)
	got = inOrderSuccesor(&leftRightRightSub)
	assert.Assert(t, got != nil)
	assert.Assert(t, got.Data == 20)
}
