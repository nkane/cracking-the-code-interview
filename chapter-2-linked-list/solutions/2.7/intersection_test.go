package intersection

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

func TestIntersection(t *testing.T) {
	iNode := linked_list.Node[int]{
		Data: 99,
	}
	n1 := linked_list.Node[int]{
		Data: 1,
	}
	n2 := linked_list.Node[int]{
		Data: 2,
	}
	n3 := linked_list.Node[int]{
		Data: 3,
	}
	l1 := linked_list.LinkedList[int]{}
	l1.Add(&n1)
	l1.Add(&n2)
	l1.Add(&iNode)
	l1.Add(&n3)

	n4 := linked_list.Node[int]{
		Data: 10,
	}
	n5 := linked_list.Node[int]{
		Data: 20,
	}
	n6 := linked_list.Node[int]{
		Data: 30,
	}
	l2 := linked_list.LinkedList[int]{}
	l2.Add(&n4)
	l2.Add(&n5)
	l2.Add(&n6)
	l2.Add(&iNode)
	intersectingNode := FindIntersection(l1.Head, l2.Head)
	assert.Assert(t, intersectingNode == &iNode, "")
}
