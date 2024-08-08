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

// TreeNode represents a node in the binary search tree
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
	Size  int
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetRandomNode returns a random node from the tree
func (n *TreeNode) GetRandomNode() *TreeNode {
	leftSize := 0
	if n.Left != nil {
		leftSize = n.Left.Size
	}
	idx := r.Intn(n.Size)
	if idx < leftSize {
		return n.Left.GetRandomNode()
	} else if idx == leftSize {
		return n
	} else {
		return n.Right.GetRandomNode()
	}
}

// InsertInOrder inserts a node into the tree, maintaining the BST property
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

// Find searches for a node with the given value in the tree
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
