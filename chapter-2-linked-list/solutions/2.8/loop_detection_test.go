package loop_detection

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

func TestFindBeginning(t *testing.T) {
	a := linked_list.Node[string]{
		Data: "a",
	}
	b := linked_list.Node[string]{
		Data: "b",
	}
	c := linked_list.Node[string]{
		Data: "c",
	}
	d := linked_list.Node[string]{
		Data: "d",
	}
	e := linked_list.Node[string]{
		Data: "e",
	}
	l := linked_list.LinkedList[string]{}
	l.Add(&a)
	l.Add(&b)
	l.Add(&c)
	l.Add(&d)
	l.Add(&e)
	l.Add(&c)
	loopNode := FindBeginning(l.Head)
	assert.Assert(t, &c == loopNode)
}
