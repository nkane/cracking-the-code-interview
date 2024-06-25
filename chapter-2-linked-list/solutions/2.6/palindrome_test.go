package palindrome

import (
	"linked_list"
	"testing"

	"gotest.tools/assert"
)

func TestIsPalindrome(t *testing.T) {
	l := linked_list.LinkedList[int]{}
	l.Add(&linked_list.Node[int]{
		Data: 0,
	})
	l.Add(&linked_list.Node[int]{
		Data: 1,
	})
	l.Add(&linked_list.Node[int]{
		Data: 2,
	})
	l.Add(&linked_list.Node[int]{
		Data: 1,
	})
	l.Add(&linked_list.Node[int]{
		Data: 0,
	})
	expected := true
	got := IsPalindrome(l.Head)
	assert.Assert(t, expected == got, "expected: %+v, got: %+v", expected, got)
}

func TestIsPalindromeA(t *testing.T) {
	l := linked_list.LinkedList[int]{}
	l.Add(&linked_list.Node[int]{
		Data: 0,
	})
	l.Add(&linked_list.Node[int]{
		Data: 1,
	})
	l.Add(&linked_list.Node[int]{
		Data: 2,
	})
	l.Add(&linked_list.Node[int]{
		Data: 1,
	})
	l.Add(&linked_list.Node[int]{
		Data: 0,
	})
	expected := true
	got := IsPalindromeA(l.Head)
	assert.Assert(t, expected == got, "expected: %+v, got: %+v", expected, got)
}

func TestIsPalindromeB(t *testing.T) {
	l := linked_list.LinkedList[int]{}
	l.Add(&linked_list.Node[int]{
		Data: 0,
	})
	l.Add(&linked_list.Node[int]{
		Data: 1,
	})
	l.Add(&linked_list.Node[int]{
		Data: 2,
	})
	l.Add(&linked_list.Node[int]{
		Data: 1,
	})
	l.Add(&linked_list.Node[int]{
		Data: 0,
	})
	expected := true
	got := IsPalindromeB(l.Head)
	assert.Assert(t, expected == got, "expected: %+v, got: %+v", expected, got)
}
