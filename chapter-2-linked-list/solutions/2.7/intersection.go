package intersection

import (
	"linked_list"
)

/*
	2.7: Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the
	intersecting node. Note that the intersection is defined based on reference, not value. That is,
	if the kth node of the first linked list is the exact same node (by reference) as the jth node
	of the second linked list, then they are intersecting.
*/

/*
Determining if there's an intersection.

How would we detect two linked lists intersect? One approach would be to use a hash table and just throw
all the linked list nodes into there. We would need to be careful to reference the linked list by their
memory location, not by their value.

There's an easier way though. Observe that two intersecting linked list will always have the same last
node. Therefore, we can just traverse to the end of each linked list and compare the laset nodes.

How do we find where the intersection is, though?

Finding the intersecting node.

One thougth is that we could traverse backwards through each linked list. When the linked lists "split", that's
the intersection. Of course, you can't really traverse backwards through a singly linked list.

If the linked list wre the same length, you could just traverse through them at the same time. When they collide,
that's your intersection.

When they're not the same length, we'd like to just "chop off", or ignore, those excess nodes.

How can we do this? Well, if we know the length of the two linked list, then the difference between those two linked
list will tell use how much to chop off.

We can get the length at the same time as we get the tails of the linked lists (which we used in the first step
to determine if there's an intersection).

Putting it all together.

We now have a multistep process.
1. Run through each linked list to get the lengths and the tails.
2. Compare the tails. If they are different (by reference, not by value), return immediately. There is no intersection.
3. Set two pointers to the start of each linked list.
4. On the longer linked list, advance its pointer by the difference in lengths.
5. Now, traverse on each linked list until the pointers are the same.
*/

func FindIntersection(list1 *linked_list.Node[int], list2 *linked_list.Node[int]) *linked_list.Node[int] {
	if list1 == nil || list2 == nil {
		return nil
	}
	// get tail and sizes
	result1 := GetTailAndSize(list1)
	result2 := GetTailAndSize(list2)
	// if different tail nodes, then there's no intersection
	if result1.Tail != result2.Tail {
		return nil
	}
	// set pointers to the start of each linked list
	shorter := list2
	if result1.Size < result2.Size {
		shorter = list1
	}
	longer := list1
	if result1.Size < result2.Size {
		longer = list2
	}
	// advance the pointer for the longer linked list by difference n lengths
	k := result1.Size - result2.Size
	if k < 0 {
		k *= -1
	}
	longer = GetKthNode(longer, k)
	// move both pointers until you have a collision
	for shorter != longer {
		shorter = shorter.Next
		longer = longer.Next
	}
	// return either one
	return longer
}

type Result struct {
	Tail *linked_list.Node[int]
	Size int
}

func GetTailAndSize(list *linked_list.Node[int]) *Result {
	if list == nil {
		return nil
	}
	size := 1
	current := list
	for current.Next != nil {
		size++
		current = current.Next
	}
	return &Result{
		Tail: current,
		Size: size,
	}
}

func GetKthNode(head *linked_list.Node[int], k int) *linked_list.Node[int] {
	current := head
	for k > 0 && current != nil {
		current = current.Next
		k--
	}
	return current
}
