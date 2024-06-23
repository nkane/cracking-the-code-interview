package remove_duplicates

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

func CreateTestListWithDuplicates() *linked_list.LinkedList[int] {
	l := linked_list.LinkedList[int]{}
	nodes := []*linked_list.Node[int]{
		{
			Data: 1,
		},
		{
			Data: 2,
		},
		{
			Data: 1,
		},
		{
			Data: 4,
		},
	}
	for _, n := range nodes {
		l.Add(n)
	}
	return &l
}

func TestDeleteDuplicatesWithMap(t *testing.T) {
	l := CreateTestListWithDuplicates()
	expected := `[data: 1] -> [data: 2] -> [data: 4] -> `
	DeleteDuplicatesWithMap(l.Head)
	buf := l.StringSimple()
	assert.Assert(t, expected == buf, "expected: %s, got: %s", expected, buf)
}

func TestDeleteDuplicatesWithRunner(t *testing.T) {
	l := CreateTestListWithDuplicates()
	expected := `[data: 1] -> [data: 2] -> [data: 4] -> `
	DeleteDuplicatesWithRunner(l.Head)
	buf := l.StringSimple()
	assert.Assert(t, expected == buf, "expected: %s, got: %s", expected, buf)
}
