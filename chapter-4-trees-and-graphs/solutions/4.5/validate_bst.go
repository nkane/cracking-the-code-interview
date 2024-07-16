package validate_bst

import (
	"tree"
)

/*
	4.5: Validate BST: Implement a function to check if a binary search tree.
*/

/*
We can implement this solution in two different ways. The first leverages the in-order traversal, and the
second builds off the property that left <= current <= right.

Solution #1: In-Order Traversal
Our first thought might be to do an in-order traversal, copy the elements to an array, and then check to see
if the array is sorted. This solution takes up a bit of extra memory, but it works, mostly.

The only problem is that it can't handle duplicate values in the tree properly. For example, the algorithm
cannot distinguish between the two tree below (one of which is invalid) since they have the same in-order
traversal.

However, if we assume that the tree cannot have duplicate values, then this approach works.
*/

var index = 0

func copyBST(root *tree.Node, array []int) {
	if root == nil {
		return
	}
	copyBST(root.Left, array)
	if len(array)-1 < index {
		// TODO(nick): this can be replaced with tracking the tree length
		array = append(array, 0)
	}
	array[index] = root.Data
	index++
	copyBST(root.Right, array)
}

func checkBST(root *tree.Node) bool {
	array := []int{}
	copyBST(root, array)
	for i := 1; i < len(array); i++ {
		if array[i] <= array[i-1] {
			return false
		}
	}
	return true
}

/*
Note that it is necessary to keep track of the logical "end" of the array, since it would be allocated to hold all
the elements.

When we examine this solution, we find that the array is not actually necessary. We never use it other than to compare
an element to the previous element. So why no track the last element we saw and copmare it as we go?
*/

var lastPrinted *int

func checkBST_alternative(n *tree.Node) bool {
	if n == nil {
		return true
	}
	if !checkBST_alternative((n.Left)) {
		return false
	}
	if lastPrinted != nil && n.Data <= *lastPrinted {
		return false
	}
	lastPrinted = &n.Data
	if !checkBST_alternative((n.Right)) {
		return false
	}
	return true
}

/*
Solution #2: The Min / Max Solution
In the second solution, we leverage the definition of a binary search tree.
What does it mean for a tree to be a binary search tree? We know that it must, of course, satisfy the condition
left.data <= current.data < right.data for each node, but this isn't quite sufficient.

The condition that all left nodes must be less than or equal to the current node, which must be less than all the
right nodes.

Using this thought, we can approach the problem by passing down the min and the max values. As we iterate through
the tree, we verify against progressively narrower ranges.

We start with a range of (min = NULL, max = NULL), which the root obviously meets. (NULL indicates that there is no
min or max.) We then branch left, checking that these nodes are within the range (min = NULL, max = 20). Then, we
branch right, checking that the nodes are within range (min = 20, max = NULL).

We proceed through the tree with this approach. When we branch left, max gets updated. When we branch right, min
gets updated. If anything fails these checks, we stop and return false.

The time complexity for this solution is O(N), where N is the number of nodes in the tree. We can prove that this is
the best we can do, since any algorithm must touch all N nodes.

Due to the use of recursion, the space complexity is O(log N) on a balanced tree. There are up to O(log N) recursive
calls on the stack since we may recurse up to the depth of the tree.
*/

func checkBST_2(node *tree.Node) bool {
	return checkBSTMain_2(node, nil, nil)
}

func checkBSTMain_2(n *tree.Node, min *int, max *int) bool {
	if n == nil {
		return true
	}
	if ((min != nil) && n.Data <= *min) || ((max != nil) && n.Data >= *max) {
		return false
	}
	if !checkBSTMain_2(n.Left, min, &n.Data) || !checkBSTMain_2(n.Right, &n.Data, max) {
		return false
	}
	return true
}
