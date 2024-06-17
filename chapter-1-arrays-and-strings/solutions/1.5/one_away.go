package one_away

/*
	1.5: One Away: There are three types of edits that can be performed on strings: insert a character,
	remove a character, or replace a character. Given two strings, write a function to check if they are
	one edit (or zero edits) away.

	Example:
	pale, ple -> true
	pales, pale -> true
	pale, bale -> true
	pale, bake -> false
*/

/*
There is a "brute force" algorithm to do this. We could check all possible strings that are one edit away
by testing the removal of each character (and comparing), testing the replacement of each character (and
comparing), and then testing the insertion of each possible character (and comparing).

That would be too slow, so let's not bother with implementing it.

This is one of those problems where it's helpful to think about the "meaning" of each of these operations.
What does it mean for two strings to be one insertion, replacement, or removal away from each other?

- Replacement: Consider two strings, such as bale and pale, that are one replacement away. Yes, that
does mean that you could replace a character in bale to make pale. But more precisely, it means that
they are different only in one place.

- Insertion: The strings apple and aple are one insertion away. This means that if you compared the strings, they
would be identical, except for a shift at some point in the strings.

- Removal: The strings apple and aple are also one removal away, since removal is just the inverse of insertion.

We can go ahead and implement this algorithm now. We'll merge the insertion and removal check into one step,
and check the replacement step separately.

Observe that you don't need to check the strings for insertion, removal, and replacement edits. The lengths
of the strings will indicate which of these you need to check.
*/

type OneAway func(a string, b string) bool

func OneAwaySolutionOne(a string, b string) bool {
	if len(a) == len(b) {
		return OneEditReplace(a, b)
	} else if len(a)+1 == len(b) {
		return OneEditInsert(a, b)
	} else if len(a)-1 == len(b) {
		return OneEditInsert(b, a)
	}
	return false
}

func OneEditReplace(s1 string, s2 string) bool {
	foundDifference := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}

func OneEditInsert(s1 string, s2 string) bool {
	index1 := 0
	index2 := 0
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}

/*
This algorithm (and almost any reasonable algorithm) takes O(n) time, where n is the length of the shorter string.

Why is the runtime dictated by the shorter string instead of the longer string? If the strings are the same length
(plus or minus one character), then it doesn't matter whehter we use the longer string or the shorter string to
define the runetime. If the strings are different lengths, then the algorithm will terminate in O(1) time. One really,
really long string therefore won't significantly extend then runtime. It increases the runtime only if both strings
are long.

We might notice that the code of OneEditReplace is very similar to that for OneEditInsert. We can merge them into
one method.

To do this, observe that both methods follow similar logic: compare each character and ensure that the strings are
only different by one. The method vary in how they handle the difference. The method OneEditReplace does nothing
other than flag the difference, whereas OneEditInsert increments the pointer to the longer string. We can handle
both of these in the same method.

Some people might argue the first approach is better, as it is clearer and easier to follow. Others, however, will
argue that the second approach is better, since it's more compact and doesn't duplicate code (which can facilitate
maintainability).

You don't need to "pick a side". You can discuss that tradeoffs with your interviewer.
*/
func OneAwaySolutionTwo(a string, b string) bool {
	l := len(a) - len(b)
	s1 := b
	s2 := a
	if len(a) < len(b) {
		s1 = a
		s2 = b
		l = len(b) - len(a)
	}
	if l > 1 {
		return false
	}
	index1 := 0
	index2 := 0
	foundDifference := false
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if foundDifference {
				return false
			}
			foundDifference = true
			if len(s1) == len(s2) {
				index1++
			}
		} else {
			index1++
		}
		index2++
	}
	return true
}
