package list_of_depths

import (
	"testing"
	"tree"
)

var testTree = tree.Tree{
	Root: &tree.Node{
		Data: 1,
		Left: &tree.Node{
			Data: 2,
			Left: &tree.Node{
				Data: 4,
			},
			Right: &tree.Node{
				Data: 5,
			},
		},
		Right: &tree.Node{
			Data: 3,
			Left: &tree.Node{
				Data: 6,
				Right: &tree.Node{
					Data: 7,
				},
			},
		},
	},
}

func TestListOfDepths(t *testing.T) {
	lists := ListOfDepths(testTree)
	if len(lists) != 4 {
		t.FailNow()
	}
	if lists[0].Len() != 1 {
		t.FailNow()
	}
	if lists[1].Len() != 2 {
		t.FailNow()
	}
	if lists[2].Len() != 3 {
		t.FailNow()
	}
	if lists[3].Len() != 1 {
		t.FailNow()
	}
}

func TestListOfDepthsBookSolution1(t *testing.T) {
	lists := createLevelLinkedListFinal(testTree.Root)
	if len(lists) != 4 {
		t.FailNow()
	}
	if lists[0].Len() != 1 {
		t.FailNow()
	}
	if lists[1].Len() != 2 {
		t.FailNow()
	}
	if lists[2].Len() != 3 {
		t.FailNow()
	}
	if lists[3].Len() != 1 {
		t.FailNow()
	}
}

func TestListOfDepthsBookSolution2(t *testing.T) {
	lists := createLeveLinkedLinks2(testTree.Root)
	if len(lists) != 4 {
		t.FailNow()
	}
	if lists[0].Len() != 1 {
		t.FailNow()
	}
	if lists[1].Len() != 2 {
		t.FailNow()
	}
	if lists[2].Len() != 3 {
		t.FailNow()
	}
	if lists[3].Len() != 1 {
		t.FailNow()
	}
}
