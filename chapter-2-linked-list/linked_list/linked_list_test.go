package linked_list

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

type TestAdd struct {
	Nodes    []*Node[int]
	Expected string
}

func TestLinkListAdd(t *testing.T) {
	tests := []TestAdd{
		{
			Nodes: []*Node[int]{
				{
					Data: 1,
				},
			},
			Expected: `[pos: 1 | data: 1] -> `,
		},
		{
			Nodes: []*Node[int]{
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
			Nodes: []*Node[int]{
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
		l := &LinkedList[int]{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		buf := fmt.Sprintf("%v", l)
		assert.Assert(t, test.Expected == buf, "output doesn't match")
	}
}

type TestFind struct {
	Nodes        []*Node[int]
	FindData     int
	ExpectedData int
}

func TestLinkedListFind(t *testing.T) {
	tests := []TestFind{
		{
			Nodes: []*Node[int]{
				{
					Data: 1,
				},
			},
			FindData:     1,
			ExpectedData: 1,
		},
		{
			Nodes: []*Node[int]{
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
					Data: 4,
				},
				{
					Data: 5,
				},
				{
					Data: 6,
				},
				{
					Data: 7,
				},
			},
			FindData:     4,
			ExpectedData: 4,
		},
	}
	for _, test := range tests {
		l := LinkedList[int]{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		n := l.Find(test.FindData)
		assert.Assert(t, n != nil, "returned nil node")
		assert.Assert(t, n.Data == test.ExpectedData, "expected data doesn't match")
	}
}

type TestRemove struct {
	Nodes          []*Node[int]
	RemoveIndex    int
	ExpectedLength int
	ExpectedString string
}

func TestLinkListRemove(t *testing.T) {
	tests := []TestRemove{
		{
			Nodes: []*Node[int]{
				{
					Data: 1,
				},
			},
			RemoveIndex:    1,
			ExpectedLength: 0,
			ExpectedString: ``,
		},
		{
			Nodes: []*Node[int]{
				{
					Data: 1,
				},
				{
					Data: 2,
				},
			},
			RemoveIndex:    1,
			ExpectedLength: 1,
			ExpectedString: `[pos: 1 | data: 2] -> `,
		},
		{
			Nodes: []*Node[int]{
				{
					Data: 1,
				},
				{
					Data: 2,
				},
			},
			RemoveIndex:    2,
			ExpectedLength: 1,
			ExpectedString: `[pos: 1 | data: 1] -> `,
		},
		{
			Nodes: []*Node[int]{
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
					Data: 4,
				},
				{
					Data: 5,
				},
				{
					Data: 6,
				},
				{
					Data: 7,
				},
			},
			RemoveIndex:    4,
			ExpectedLength: 6,
			ExpectedString: `[pos: 1 | data: 1] -> [pos: 2 | data: 2] -> [pos: 3 | data: 3] -> [pos: 4 | data: 5] -> [pos: 5 | data: 6] -> [pos: 6 | data: 7] -> `,
		},
	}
	for _, test := range tests {
		l := &LinkedList[int]{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		l.Remove(test.RemoveIndex)
		assert.Assert(t, test.ExpectedLength == l.Length, "length do not match")
		buf := fmt.Sprintf("%v", l)
		assert.Assert(t, test.ExpectedString == buf, "output doesn't match")
	}
}
