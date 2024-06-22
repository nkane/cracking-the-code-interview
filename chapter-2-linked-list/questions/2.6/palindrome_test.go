package palindrome

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Nodes    []*linked_list.Node[string]
	Expected string
}

func TestPalindrome(t *testing.T) {
	tests := []Test{
		{
			Nodes: []*linked_list.Node[string]{
				{
					Data: "r",
				},
				{
					Data: "a",
				},
				{
					Data: "c",
				},
				{
					Data: "e",
				},
				{
					Data: "c",
				},
				{
					Data: "a",
				},
				{
					Data: "r",
				},
			},
			Expected: `[data: r] -> [data: a] -> [data: c] -> [data: e] -> [data: c] -> [data: a] -> [data: r] -> `,
		},

		{
			Nodes: []*linked_list.Node[string]{
				{
					Data: "r",
				},
				{
					Data: "a",
				},
				{
					Data: "t",
				},
			},
			Expected: `[data: t] -> [data: a] -> [data: r] -> `,
		},
	}

	for _, test := range tests {
		l1 := &linked_list.LinkedList[string]{}
		for _, node := range test.Nodes {
			l1.Add(node)
		}
		newList := Palindrome(l1)
		output := newList.StringSimple()
		assert.Assert(t, output == test.Expected, "output does not match - expected: %s, got: %s", test.Expected, output)
	}
}
