package k_to_last

import (
	"bytes"
	"linked_list"
	"log"
	"os"
	"testing"

	"gotest.tools/assert"
)

func CreateTestList() *linked_list.LinkedList[int] {
	l := linked_list.LinkedList[int]{}
	nodes := []*linked_list.Node[int]{
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
	}
	for _, n := range nodes {
		l.Add(n)
	}
	return &l
}

func TestPrintKthToLast(t *testing.T) {
	// capture output
	buf := bytes.Buffer{}
	log.SetOutput(&buf)
	flags := log.Flags()
	log.SetFlags(0)
	expected := "2th to the last node is 3\n"

	l := CreateTestList()
	PrintKthToLast(l.Head, 2)
	s := string(buf.Bytes())
	log.SetOutput(os.Stderr)
	log.SetFlags(flags)

	assert.Assert(t, s == expected, "expected: %s, got: %s", expected, s)
}

func TestPrintKthToLastWrapper(t *testing.T) {
	l := CreateTestList()

	got := KthToLastWrapper(l.Head, 2)
	expected := l.Head.Next.Next

	assert.Assert(t, expected == got, "expected: %+v, got: %+v", expected, got)
}

func TestNKthToLast(t *testing.T) {
	l := CreateTestList()

	got := NthToLast(l.Head, 2)
	expected := l.Head.Next.Next

	assert.Assert(t, expected == got, "expected: %+v, got: %+v", expected, got)
}
