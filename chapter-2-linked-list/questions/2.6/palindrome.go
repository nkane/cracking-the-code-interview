package palindrome

import (
	"linked_list"
)

/*
	2.6: Palindrome: Implement a function to check if a linked list is a palindrome.
*/

func Palindrome[T comparable](l *linked_list.LinkedList[T]) *linked_list.LinkedList[T] {
	// create a new list an reverse it
	newList := linked_list.LinkedList[T]{}
	AddNode(&newList, l.Head)
	return &newList
}

func AddNode[T comparable](l *linked_list.LinkedList[T], n *linked_list.Node[T]) *linked_list.Node[T] {
	newNode := linked_list.Node[T]{
		Data: n.Data,
	}
	if n.Next == nil {
		l.Head = &newNode
		return &newNode
	}
	addedNode := AddNode(l, n.Next)
	addedNode.Next = &newNode
	return &newNode
}
