package partition

import (
	"fmt"
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Nodes     []*linked_list.Node[int]
	Partition int
}

func TestPartition(t *testing.T) {
	tests := []Test{
		{
			Nodes: []*linked_list.Node[int]{
				{
					Data: 3,
				},
				{
					Data: 5,
				},
				{
					Data: 8,
				},
				{
					Data: 5,
				},
				{
					Data: 10,
				},
				{
					Data: 2,
				},
				{
					Data: 1,
				},
			},
			Partition: 5,
		},
	}

	for _, test := range tests {
		l := &linked_list.LinkedList[int]{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		fmt.Printf("before: %s\n", l.StringSimple())
		l.Head = Partition(l.Head, test.Partition)
		fmt.Printf("after:  %s\n", l.StringSimple())
		assert.Assert(t, CheckPivot(l, test.Partition), "failed pivot check")
	}
}

func TestPartitionSimple(t *testing.T) {
	tests := []Test{
		{
			Nodes: []*linked_list.Node[int]{
				{
					Data: 3,
				},
				{
					Data: 5,
				},
				{
					Data: 8,
				},
				{
					Data: 5,
				},
				{
					Data: 10,
				},
				{
					Data: 2,
				},
				{
					Data: 1,
				},
			},
			Partition: 5,
		},
	}

	for _, test := range tests {
		l := &linked_list.LinkedList[int]{}
		for _, node := range test.Nodes {
			l.Add(node)
		}
		assert.Assert(t, len(test.Nodes) == l.Length, "length do not match")
		fmt.Printf("before: %s\n", l.StringSimple())
		l.Head = PartitionSimple(l.Head, test.Partition)
		fmt.Printf("after:  %s\n", l.StringSimple())
		assert.Assert(t, CheckPivot(l, test.Partition), "failed pivot check")
	}
}

func CheckPivot(l *linked_list.LinkedList[int], p int) bool {
	// ensure all nodes are less than p until p is found
	n := l.Head
	pFound := false
	for n != nil {
		if n.Data == p {
			pFound = true
		}
		if !pFound {
			if n.Data > p {
				return false
			}
		} else {
			if n.Data < p {
				return false
			}
		}
		n = n.Next
	}
	return true
}
