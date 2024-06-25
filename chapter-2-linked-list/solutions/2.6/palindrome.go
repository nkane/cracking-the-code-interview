package palindrome

import (
	"linked_list"

	"github.com/golang-collections/collections/stack"
)

/*
	2.6: Palindrome: Implement a function to check if a linked list is a palindrome.
*/

/*
To approach this problem, we can picture a palindrome like 0 -> 1 -> 2 -> 1 -> 0. We know that, since
it's a palidrome, the list must be the same backwards and forwards. This leads us to our first solution.

Solution #1: Reverse and Compare
Our first solution is to reverse the linked list and compare the reversed list to the original list. If they're
the same, the lists are identical.

Note that when we compare the linked list to the reversed list, we only actually need to compare the first half
of the list. If the first half of the normal list matches the first half of the reversed list, then the second half
of the normal list must match the second half of the reversed list.
*/

func IsPalindrome(head *linked_list.Node[int]) bool {
	reversed := ReverseAndClone(head)
	return IsEqual(head, reversed)
}

func ReverseAndClone(node *linked_list.Node[int]) *linked_list.Node[int] {
	var head *linked_list.Node[int]
	for node != nil {
		n := linked_list.Node[int]{
			Data: node.Data,
		}
		n.Next = head
		head = &n
		node = node.Next
	}
	return head
}

func IsEqual(one *linked_list.Node[int], two *linked_list.Node[int]) bool {
	for one != nil && two != nil {
		if one.Data != two.Data {
			return false
		}
		one = one.Next
		two = two.Next
	}
	return one == nil && two == nil
}

/*
Solution #2: Iterative Approach

We want to detect linked lists where the front half of the list is the reverse of the second half. How would we
do that? By reversing the front half of the list. A stack can accomplish this.

We need to push the first half of the elements onto the stack. We can do this is two different ways, depending
on whether or not we know the size of the linked list.

If we know the size of the linked list, we can iterate through the first half of the elements in a standard for
loop, puhsing each element onto a stack. We must be careful, of course, to handle the case where the length
of the linked list is odd.

If we don't know the size of the linked list, we can iterate through the linked list, using the faster runner / slow
runner techinque described in the beginning of the chapter. At each step in the loop, we push the data from the slow
runner onto a stack. When the fast runner hits the end of the list, the slow runner will have reached the middle
of the linked list. By this point, the stack will have all elements from the front of the linked list, but in reverse
order.

Now, we simply iterate through the rest of the linked list. At each iteration, we compare the node to the top of the
stack. If we complete the iteration without finding a difference, then the linked list is a palindrome.
*/

func IsPalindromeA(head *linked_list.Node[int]) bool {
	fast := head
	slow := head
	stack := stack.New()
	// push elements from first half of linked list onto stack. When fast runner
	// (which is moving 2x speed) reaches the end of the linked list, then we
	// know we're at the middle.
	for fast != nil && fast.Next != nil {
		stack.Push(slow.Data)
		slow = slow.Next
		fast = fast.Next
	}
	// has odd number of elements, so skip the middle element
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		top := stack.Pop().(int)
		// if values are different, then it's not a IsPalindrome
		if top != slow.Data {
			return false
		}
		slow = slow.Next
	}
	return true
}

/*
Solution #3: Recursive Approach
First, a word on notation: in this solution, when we use the notion node Kx, the variable K indicates the value of
the node data, and x (which is either f or b) indicates whether we are referring to the front node with the value
or the back node. For example, in the below linked list, node 2b would refer to the second (back) node with value 2.

Now, like many linked list problems, you can approach this problem recursively. We may have some intuitive idea that
we want to compare element 0 and element n - 1, element 1 and element n - 2, element 2 and element n - 3, and so on,
until the middle element(s). For example:

									0 ( 1 ( 2 ( 3 ) 2 ) 1 ) 0
In order to apply this approach, we first need to know we've reached the middle element, as this will form our base
case. We can do this by passing the length - 2 for the length each time. When the length equals 0 or 1, we're at the
center of the linked list. This is because the length is reduced by 2 each time. Once we've recursed N/2 times,
length will be down to 0.
*/

type Result struct {
	Node   *linked_list.Node[int]
	Result bool
}

func IsPalindromeB(head *linked_list.Node[int]) bool {
	length := LengthOfList(head)
	p := IsPalindromeRecurse(head, length)
	return p.Result
}

func IsPalindromeRecurse(head *linked_list.Node[int], length int) Result {
	if head == nil || length <= 0 { // even number of nodes
		return Result{
			Node:   head,
			Result: true,
		}
	} else if length == 1 {
		return Result{
			Node:   head.Next,
			Result: true,
		}
	}
	// recurse on sublist
	res := IsPalindromeRecurse(head.Next, length-2)
	// if child calls are not a palindrome, pass back up
	// a failure
	if !res.Result || res.Node == nil {
		return res
	}
	// check if matches corresponding node on other side
	res.Result = (head.Data == res.Node.Data)
	// return corresponding node
	res.Node = res.Node.Next
	return res
}

func LengthOfList(n *linked_list.Node[int]) int {
	size := 0
	for n != nil {
		size++
		n = n.Next
	}
	return size
}
