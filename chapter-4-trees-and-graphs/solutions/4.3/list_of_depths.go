package list_of_depths

import (
	"container/list"
	"tree"

	"github.com/golang-collections/collections/queue"
)

/*
	4.3: List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes
	at each depth (e.g., if you have a tree with depth D, you'll have a D linked list)
*/

func ListOfDepths(t tree.Tree) []*list.List {
	linkedLists := []*list.List{}
	q := queue.New()
	q.Enqueue(t.Root)
	for q.Len() > 0 {
		n := q.Dequeue().(*tree.Node)
		l := list.New().Init()
		currentN := n
		var ok bool
		for currentN != nil {
			l.PushFront(currentN)
			currentN, ok = q.Dequeue().(*tree.Node)
			if !ok {
				break
			}
		}
		currentL := l.Front()
		for currentL != nil {
			currentN, ok := currentL.Value.(*tree.Node)
			if !ok {
				break
			}
			if currentN.Left != nil {
				q.Enqueue(currentN.Left)
			}
			if currentN.Right != nil {
				q.Enqueue(currentN.Right)
			}
			currentL = currentL.Next()
		}
		linkedLists = append(linkedLists, l)
		// push all depth nodes
	}
	return linkedLists
}

/*
Though we might think a first glance that this problem requires a level-by-level traversal, this isn't actually
necessary. We can traverse the graph any way that'd we'd like, provided we know which level we're on.

We can implement a sipmle modification of the pre-order traversal algorithm, where we pass in level + 1 to the next
recursive call.
*/

func createLevelLinkedList(root *tree.Node, lists *[]*list.List, level int) {
	if root == nil {
		return
	}
	var l *list.List
	if len(*lists) == level { // level not contained in list
		l = list.New().Init()
		// levels are always traversed in order, so, if this is the first time we've
		// visited level i, we must have seen levels 0 through i - 1. We can therefore
		// safely add the level at the end.
		*lists = append(*lists, l)
	} else {
		ls := *lists
		l = ls[level]
	}
	l.PushFront(root)
	createLevelLinkedList(root.Left, lists, level+1)
	createLevelLinkedList(root.Right, lists, level+1)
}

func createLevelLinkedListFinal(root *tree.Node) []*list.List {
	lists := []*list.List{}
	createLevelLinkedList(root, &lists, 0)
	return lists
}

/*
Alternatively, we can also implement a modification of breadth-first search. With this implementation, we want
to iterate through the root first, then level 2, and then level 3, and so on.

With each level i, we wil already fully visited all nodes on leve i - 1, this means that to get which nodes are on
level i, we can simply look at all children of the nodes of level i - 1.
*/

func createLeveLinkedLinks2(root *tree.Node) []*list.List {
	lists := []*list.List{}
	current := list.New().Init()
	if root != nil {
		current.PushFront(root)
	}
	for current.Len() > 0 {
		lists = append(lists, current)
		parents := current
		current = list.New().Init()
		parent := parents.Front()
		for parent != nil {
			// visit the chilred
			n, ok := parent.Value.(*tree.Node)
			if !ok {
				break
			}
			if n.Left != nil {
				current.PushFront(n.Left)
			}
			if n.Right != nil {
				current.PushFront(n.Right)
			}
			parent = parent.Next()
		}
	}
	return lists
}
