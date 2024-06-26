package loop_detection

import (
	"linked_list"
)

/*
	2.8: Loop Detection: Given a linked list which might contain a loop, implement an algorithm
	that returns the node at the beginning of the loop (if one exists).

	Example:
	Input: a -> b -> c -> d -> e -> c [the same c as earlier]
	Output: c
*/

/*
This is a modification of a classic interview question: detect if a linked list has a loop.
Let's apply a pattern matching approach.

Part 1: Detect If Link List Has A Loop
Any easy way to detect if a linked list has a loop is through the FastRunner / SlowRunner approach.
FastRunner moves two steps at a time, while SlowRunner moves one step. Much like two cars racing
around a track at different steps, they must eventually meet.

An astute reader may wonder if FastRunner might "hop over" SlowRunner completely, without ever
colliding. That's not possible. Suppose that FastRunner did hop over SlowRunner, such that SlowRunner
is at spot i and FastRunner is at spot i + 1. In the previous step, SlowRunner would be at spot i - 1
and FastRunner would at ((i + 1) - 2), or spot i - 1. That is, they would have collided.

Part 2: When Do They Collide?
Let's assume that the linked list has a "non-looped" part of size k.
If we apply our algorithm from part 1, when will FastRunner and SlowRunner collide?
We know that for every p steps that SlowRunner takes, FastRunner has taken 2p steps. Therefore, when
SlowRunner enters the looped portion after k steps, FastRunner has take 2k steps total and must be
2k - k steps, or k steps, into the loop portion. Since k might be much larger than the loop length,
we should actually write this as mod(k, LOOP_SIZE) steps, which we will denote as K.

At each subsequent step, FastRunner and SlowRunner get either one step farther away or one step close,
depending on your perspective. That is, because we are in a circle, when A moves q steps away from B,
it is also moving q steps closer to B.

So now we know the following facts:
1. SlowRunner is 0 steps into the loop.
2. FastRunner is K steps into the loop.
3. SlowRunner is K steps behind the loop.
4. FastRunner is LOOP_SIZE - K steps behind SlowRunner.
5. FastRunner catches up to SlowRunner at a rate of 1 step per unit of time.

So, when do they meet? Well, if FastRunner is LOOP_SIZE - k steps behind SlowRunner, and FastRunner
catches up at a rate of 1 step per unit of time, then they meet after LOOP_SIZE - K steps. At this
point, they will be K steps before the head of the loop. Let's call this point CollisionSpot.

Part 3: How Do You Find The Start of the Loop?
We now know that CollisionSpot is K nodes before the start of the loop. Because K = mod(k, LOOP_SIZE)
(or, in other words, k = k + m * LOOP_SIZE, for any integer M), it is also correct to say that it is k
nodes from the loop start. For example, if node N is 2 nodes into a 5 node loop, it is also correct to
say that it is 7, 12, or even 397 nodes into the loop.

Therefore, both CollisionSpot and LinkedListHead are k nodes from the start of the loop.

Now, if we keep one pointer at CollisionSpot and move the other one to LinkedListHead, they will each
be k nodes from LoopStart. that it is 7, 12, or even 397 nodes into the loop.

Therefore, both CollisionSpot and LinkedListHead are k nodes from the start of the loop.

Now, if we keep one pointer at CollisionSpot and move the other one to LinkedListHead, they will each
be k nodes from LoopStart. Moving the two pointers at the same speed will cause them to collide again,
this time after k steps, at which point they will both be at LoopStart. All we have to do is return this
node.

Part 4: Putting It All Together
To summarize, we move FastPointer twice as fast as SlowPointer. When SlowPointer enters the loop,
after k nodes, FastPointer is k nodes into the loop. This means that FastPointer and SlowPointer
are LOOP_SIZE - k nodes away from each other.

Next, if FastPointer moves two nodes for each node that SlowPointer moves, they move one node closer
to each other on each turn. Therefore, they will meet after LOOP_SIZE - k turns. Both will be k nodes
from the front of the loop.

The head of the linked list is also k nodes from the front of the loop. So, if we keep one pointer where
it is, and move the other pointer to the head of the linked list, then they will meet at the front of
the loop.

Our algorithm is derived from parts 1, 2, and 3
1. Create two pointers, FastPointer and SlowPointer
2. Move FastPointer at a rate of 2 steps and SlowPointer at a rate of 1 step.
3. When they collide, move SlowPointer to LinkedListHead, keep FastPointer where it is.
4. Move SlowPointer and FastPointer at a rate of one step. Return the new collision point.
*/

func FindBeginning(head *linked_list.Node[string]) *linked_list.Node[string] {
	slow := head
	fast := head
	// find meeting point, this will be LOOP_SIZE - k steps into the linked list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // collision
			break
		}
	}
	// error check, no meeting point, and therefor no loop
	if fast == nil || fast.Next == nil {
		return nil
	}
	// move slow to head, keep fast at meeting point, each are k steps from the
	// loop start
	// if they move at the same padce, they must meet at loop start
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
