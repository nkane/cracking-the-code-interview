package palindrome_permutation

/*
	1.4: Palindrome Permutation: Given a string, write a function that checks if it is a permutation of a palindrome.
	A palindrome is a word or phrase that is the same forward and backwards. A permutation is a rearrangement of
	letters. The palindrome does not need to be limited to just dictionary words. You can ignore casing and non-letter
	characters.

	Example:
	Input: Tact Coa
	Output: True (permutations: "taco cat", "atco cta", etc.)
*/

/*
This is a question where it helsp to figure out what it means for a string to be a permutation of a palindrome.
This is like asking what the "defining features" of such a string would be.

A palindrome is a string that is the same forward and backwards. Therefore, to decide if a string is a permutation
of a palindrome, we need to know if it can be written such that it's the same forward and backwards.

What does it take to be able to write a set of characters the same way forward and backwards? We need to have an
even number of almost all characters, so that half can be on one side and half can be on the other side. At most
one character (the middle character)  can have an odd count.

For example, we know tactcoapapa is a permutation of a palindrome because it has Ts, four As, and two Cs, two Ps,
and one O. That O would be the center of all possible palindromes.
*/

/*
 To be more precise, strings with even length (after removing all non-letter characters) must have
 all even counts of characters. Strings of an odd length must have exactly one character with an odd
 count. Of course, and "even" string can't have an odd number of exactly one character, otherwise
 it wouldn't ben an even-length string (an odd number + many even numbers = an odd number). Likewise,
 a string with odd length can't have all characters with even counts (sum of evens is even). It's
 therefore sufficient to say that, to be a permutation of a palindrome, a string can have no more than
 one character that is odd. This will cover both the odd and the even cases.
*/

type IsPalindromePermutation func(s string) bool

/*
 Solution #1:

 Implementing this algorithm is fairly straightforward. We use a hash table to count how many times each
 character appears. Then, we iterate through the hash table and ensure that no more than one character
 has an odd count.
*/

func IsPalindromePermutationSolutionOne(s string) bool {
	table := BuildCharFrequencyTable(s)
	return CheckMaxOneOdd(table)
}

func CheckMaxOneOdd(table []int) bool {
	foundOdd := false
	for _, count := range table {
		if count%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

/*
Map each character to a number. a -> 0, b -> 1, c -> 2, etc.
This is case insensitive. Non-letter characters map to -1.
*/
func GetCharNumber(c rune) int {
	a := rune('a')
	z := rune('z')
	if a <= c && c <= z {
		return int(c - a)
	}
	return -1
}

/*
Count how many times each character appears.
*/
func BuildCharFrequencyTable(phrase string) []int {
	length := rune('z') - rune('a') + 1
	table := make([]int, length)
	for _, c := range phrase {
		x := GetCharNumber(c)
		if x != -1 {
			table[x]++
		}
	}
	return table
}

/*
Solution #2:

We can't optimize the big O time here since any algorithm will always have to look through the entire
string. However, we can make some smaller incremental improvements. Because this is a relatively simple
problem, it can be worthwhile to discuss some small optimizations or at least some tweaks.

Instead of checking the number of odd counts at the ned, we can check as we go along. Then, as soon as we
get to the end, we can have our answer.
*/
func IsPalindromePermutationSolutionTwo(s string) bool {
	countOdd := 0
	length := rune('z') - rune('a') + 1
	table := make([]int, length)
	for _, c := range s {
		x := GetCharNumber(c)
		if x != -1 {
			table[x]++
			if table[x]%2 == 1 {
				countOdd++
			} else {
				countOdd--
			}
		}
	}
	return countOdd <= 1
}

/*
It's important to be very clear here that this is not necessarily more optimal. It has the same big O time and
might even be slightly slower. We have eliminated a final iteration through the hash table, but now we
have to run a few extra lines of code for each character in the string.

You should discuss this with your interviewer as an alternative, but not necessarily more optimal, solution.
*/

/*
Solution #3:
If you think more deeply about this problem, you might notice that we don't actually need to know the counts. We
just need to know if the count is even or odd. Think about flipping a light on/off (that is initially off). If the
light winds up in the off state, we don't know how many times we flipped it, but we do know it was an even count.

Given this, we can use a single integer (as a bit vector). We we see a letter, we map it to an integer between 0 and
25 (assuming an English alphabet). Then we toggle the bit at the value. At the end of the iteration, we check that at
most one bit in the integer is set to 1. Let's start with checking that integer has one 1.

Picture an integer like 00010000. We could of course shift the integer repeatedly to check that there's only a single
1. Alternatively, if we subtract 1 from the number, we'll get 00001111. What's notable about this is that there is no
overlap between the numbers (as opposed to say 0010100, which, when we subtract 1, we get 00100111). So, we can check
to see that a number has exactly one 1 because if we subtract 1 from it and then AND it with the new number, we should
get 0.

00010000 - 1 = 00001111
00010000 & 00001111 = 0

Note that, if the same calculation is performed on 0, you will still get 0.

00000000 - 1 = 11111111
00000000 & 11111111 = 0

Therefore, this calculation will correctly assess if at most one bit is set to 1. This leads us to our final
implementation.
*/

/*
Toggle the ith bit in the integer
*/
func Toggle(bitVector int, index int) int {
	if index < 0 {
		return bitVector
	}
	mask := 1 << index
	bitVector ^= mask
	return bitVector
}

/*
Create bit vector for string. For each letter with i,
toggle with ith bit.
*/
func CreateBitVector(phrase string) int {
	bitVector := 0
	for _, c := range phrase {
		x := GetCharNumber(c)
		bitVector = Toggle(bitVector, x)
	}
	return bitVector
}

/*
Check taht at most one bit is set by subtracting one from the
integer and ANDing it with the original integer.
*/
func CheckAtMostOneBitSet(bitVector int) bool {
	return (bitVector & (bitVector - 1)) == 0
}

func IsPalindromePermutationSolutionThree(s string) bool {
	bitVector := CreateBitVector(s)
	return CheckAtMostOneBitSet(bitVector)
}
