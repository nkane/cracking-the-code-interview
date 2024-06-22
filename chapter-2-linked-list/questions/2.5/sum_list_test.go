package sum_list

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	ANodes         []*linked_list.Node
	BNodes         []*linked_list.Node
	Expected       []*linked_list.Node
	ExpectedString string
}

func TestSumlist(t *testing.T) {
	tests := []Test{
		{
			ANodes: []*linked_list.Node{
				{
					Data: 7,
				},
				{
					Data: 1,
				},
				{
					Data: 6,
				},
			},
			BNodes: []*linked_list.Node{
				{
					Data: 5,
				},
				{
					Data: 9,
				},
				{
					Data: 2,
				},
			},
			Expected: []*linked_list.Node{
				{
					Data: 9,
				},
				{
					Data: 1,
				},
				{
					Data: 2,
				},
			},
			ExpectedString: `[data: 2] -> [data: 1] -> [data: 9] -> `,
		},
		{
			ANodes: []*linked_list.Node{
				{
					Data: 5,
				},
				{
					Data: 2,
				},
			},
			BNodes: []*linked_list.Node{
				{
					Data: 0,
				},
				{
					Data: 0,
				},
				{
					Data: 1,
				},
			},
			Expected: []*linked_list.Node{
				{
					Data: 5,
				},
				{
					Data: 2,
				},
				{
					Data: 1,
				},
			},
			ExpectedString: `[data: 5] -> [data: 2] -> [data: 1] -> `,
		},
		{
			ANodes: []*linked_list.Node{
				{
					Data: 9,
				},
				{
					Data: 7,
				},
				{
					Data: 8,
				},
			},
			BNodes: []*linked_list.Node{
				{
					Data: 6,
				},
				{
					Data: 8,
				},
				{
					Data: 5,
				},
			},
			Expected: []*linked_list.Node{
				{
					Data: 5,
				},
				{
					Data: 2,
				},
				{
					Data: 1,
				},
			},
			ExpectedString: `[data: 5] -> [data: 6] -> [data: 4] -> [data: 1] -> `,
		},
	}

	for _, test := range tests {
		l1 := &linked_list.LinkedList{}
		l2 := &linked_list.LinkedList{}
		for _, node := range test.ANodes {
			l1.Add(node)
		}
		for _, node := range test.BNodes {
			l2.Add(node)
		}
		assert.Assert(t, len(test.ANodes) == l1.Length, "length do not match")
		assert.Assert(t, len(test.BNodes) == l2.Length, "length do not match")
		SumList(l1, l2)
		buf := l2.StringSimple()
		assert.Assert(t, buf == test.ExpectedString, "output doesn't match - expected: %s, got: %s", buf, test.ExpectedString)
	}
}
