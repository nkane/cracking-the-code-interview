package permutation

import (
	"reflect"
	"sort"
)

/*
	1.2: Check Permutation: Given two strings, write a method to decide if one is a permutation of the
	other.
*/

/*
Like many questions, we should confirm some details with our interviewer. We should understand if the
permutation comparison is case sensitive. That is: is God a permutation of dog? Additionally, we should
ask if whitespace is significant. We will assume for this problem that the comparison is case sensitive and
whitespace is significant. So, "god    " is different from "dog".

Observe first that strings of different lengths cannot be permutations of each other. There are two easy ways
to solve this problem, both of which use this optimization.
*/

type IsPermutation func(a string, b string) bool

/*
Solution #1: Sort the strings

If two strings are permutations, then we know they have the same characters, but in different orders.
Therefore, sorting the strings will put the characters from two permutations in the same order. We just
need to compare the sorted versions of the strings

although this algorithm is not as optimal in some sense, it may be preferable in others; It's clean, simple
and easy to understand. In a practical sense, this may very well be a superior way to implement it.

However, if efficiency is very important, we can implement it a different way.
*/

type RuneSlice []rune

func (rs RuneSlice) Len() int {
	return len(rs)
}

func (rs RuneSlice) Less(i int, j int) bool {
	return rs[i] < rs[j]
}

func (rs RuneSlice) Swap(i int, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func IsPermutationSolutionOne(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aRunes := RuneSlice([]rune(a))
	sort.Sort(aRunes)

	bRunes := RuneSlice([]rune(b))
	sort.Sort(bRunes)

	return reflect.DeepEqual(aRunes, bRunes)
}

/*
Solution #2: Check if two strings have identical counts.

We can also use the definition of a permutation, two words with the same character counts, to implement
this algorithm. We create an array that operates somewhat like a hash table, mapping each character to
it's frequency. We increment through the first string, then decrement through the second string. If the
strings are permutations, then the array will be all zeroes at the end.

We can terminate early if a value ever turns negative (once negative, the value will stay negative and
therefore non-zero). If we don't terminate early, then the array must be all zeros. This is because the
strings are the same length and we incremented the same number of times we decremented. The array cannot
have any positive values if it doesn't have any negative values.
*/

func IsPermutationSolutionTwo(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	// assumption: ASCII
	letters := [128]int{0}
	for i := 0; i < len(a); i++ {
		letters[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		letters[b[i]]--
		if letters[b[i]] < 0 {
			return false
		}
	}
	return true
}
