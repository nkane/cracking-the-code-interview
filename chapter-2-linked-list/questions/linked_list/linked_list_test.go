package linked_list

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Nodes    []*Node
	Expected string
}

func TestLinkListAdd(t *testing.T) {
	tests := []Test{
		{
			Nodes: []*Node{
				{
					Data: 1,
				},
			},
			Expected: `[pos: 1 | data: 1] -> `,
		},
		{
			Nodes: []*Node{
				{
					Data: 1,
				},
				{
					Data: 2,
				},
			},
			Expected: `[pos: 1 | data: 1] -> [pos: 2 | data: 2] -> `,
		},
		{
			Nodes: []*Node{
				{
					Data: 1,
				},
				{
					Data: 2,
				},
				{
					Data: 3,
				},
			},
			Expected: `[pos: 1 | data: 1] -> [pos: 2 | data: 2] -> [pos: 3 | data: 3] -> `,
		},
	}

	for _, test := range tests {
		l := &LinkedList{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		buf := fmt.Sprintf("%v", l)
		assert.Assert(t, test.Expected == buf, "output doesn't match")
	}
}
