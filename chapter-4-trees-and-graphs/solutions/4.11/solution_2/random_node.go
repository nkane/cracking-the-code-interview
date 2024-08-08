package random_node

import (
	"math/rand"
	"time"
)

/*
	4.11: Random Node: You are implementing a binary search tree class from scratch, which, in addition to insert,
	find, and delete, has a method getRandomNode() which returns a random node from the tree. All nodes should be
	equally likely to be chosen. Design and implement an algorithm for getRandomNode, and explain how you would
	implement the rest of the methods.
*/

type Tree struct {
	Root *TreeNode
	Size int
}

func (t *Tree) GetRandomNode() *TreeNode {
	if t.Root == nil {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(t.Size)
	return t.Root.GetIthNode(i)
}

func (t *Tree) InsertInOrder(value int) {
	if t.Root == nil {
		t.Root = &TreeNode{
			Data: value,
		}
	} else {
		t.Root.InsertInOrder(value)
	}
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
	Size  int
}

func (n *TreeNode) GetIthNode(i int) *TreeNode {
	leftSize := 0
	if n.Left != nil {
		leftSize = n.Left.Size
	}
	if i < leftSize {
		return n.Left.GetIthNode(i)
	} else if i == leftSize {
		return n
	} else {
		// skipping over leftSize + 1 nodes, so subtract them
		return n.Right.GetIthNode(i - (leftSize + 1))
	}
}

func (n *TreeNode) InsertInOrder(d int) {
	if d <= n.Data {
		if n.Left == nil {
			n.Left = &TreeNode{
				Data: d,
				Size: 1,
			}
		} else {
			n.Left.InsertInOrder(d)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{
				Data: d,
				Size: 1,
			}
		} else {
			n.Right.InsertInOrder(d)
		}
	}
	n.Size++
}

func (n *TreeNode) Find(d int) *TreeNode {
	if d == n.Data {
		return n
	} else if d < n.Data {
		if n.Left != nil {
			return n.Left.Find(d)
		}
		return nil
	} else {
		if n.Right != nil {
			return n.Right.Find(d)
		}
		return nil
	}
}
