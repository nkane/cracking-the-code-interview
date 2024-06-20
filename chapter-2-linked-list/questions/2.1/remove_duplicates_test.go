package remove_duplciates

import (
	"fmt"
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Nodes    []*linked_list.Node
	Expected string
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []Test{
		{
			Nodes: []*linked_list.Node{
				{
					Data: 1,
				},
			},
			Expected: `[pos: 1 | data: 1] -> `,
		},
		{
			Nodes: []*linked_list.Node{
				{
					Data: 1,
				},
				{
					Data: 1,
				},
			},
			Expected: `[pos: 1 | data: 1] -> `,
		},
		{
			Nodes: []*linked_list.Node{
				{
					Data: 1,
				},
				{
					Data: 2,
				},
				{
					Data: 3,
				},
				{
					Data: 1,
				},
				{
					Data: 5,
				},
			},
			Expected: `[pos: 1 | data: 1] -> [pos: 2 | data: 2] -> [pos: 3 | data: 3] -> [pos: 4 | data: 5] -> `,
		},
	}
	for _, test := range tests {
		l := &linked_list.LinkedList{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		RemoveDuplciates(l)
		buf := fmt.Sprintf("%v", l)
		assert.Assert(t, test.Expected == buf, "output doesn't match - expected: %s, got: %s", test.Expected, buf)
	}
}
