package k_last

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Nodes        []*linked_list.Node
	K            int
	ExpectedData int
}

func TestKLast(t *testing.T) {
	tests := []Test{
		{
			Nodes: []*linked_list.Node{
				{
					Data: 128,
				},
			},
			K:            0,
			ExpectedData: 128,
		},
		{
			Nodes: []*linked_list.Node{
				{
					Data: 128,
				},
				{
					Data: 256,
				},
			},
			K:            0,
			ExpectedData: 256,
		},
		{
			Nodes: []*linked_list.Node{
				{
					Data: 128,
				},
				{
					Data: 256,
				},
				{
					Data: 512,
				},
				{
					Data: 1024,
				},
			},
			K:            2,
			ExpectedData: 256,
		},
		{
			Nodes: []*linked_list.Node{
				{
					Data: 128,
				},
				{
					Data: 256,
				},
				{
					Data: 512,
				},
				{
					Data: 1024,
				},
			},
			K:            1,
			ExpectedData: 512,
		},
	}
	for _, test := range tests {
		l := &linked_list.LinkedList{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		n := KLast(l, test.K)
		assert.Assert(t, test.ExpectedData == n.Data, "data doesn't match - expected: %d, got: %d", test.ExpectedData, n.Data)
	}
}
