package unique

/*
1.1: Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use
additional data structures?
*/

/*
You should first ask your interviewer if the string is an ASCII string or a Unicode string. Asking this question
will show an eye for detail and a solid foundation in computer science. We'll assume for simplicity the character
set is ASCII. If this assumption is not valid, we would need to increase the storage size.

One solution is to create an array of bolleans, where the flag at index i indicates whether character i in
the alphabet is contained in the string. The second time you see the character i you can immediately return false.

We can also immediately return false if the string length exceed the number of unique characters in the
alphabet. After all, you can't form a string of 280 unique characters out of a 128-character alphabet.

It's okay to assume 256 characters. This would be the case in extended ASCII. You should
clarify your assumptions with your interviewer.
*/

type UniqueFunc func(str string) bool

func IsUniqueSimple(str string) bool {
	if len(str) > 128 {
		return false
	}
	charSet := [128]bool{}
	for i := 0; i < len(str); i++ {
		val := rune(str[i])
		if charSet[val] {
			return false
		}
		charSet[val] = true
	}
	return true
}

/*
We can reduce our space usage by a factor of eight by using a bit vector. We will assume, in the below code,
that the string only uses lowercase letters a through z. This will allow us to use a single int.
*/

func IsUniqueBitVector(str string) bool {
	checker := 0
	for i := 0; i < len(str); i++ {
		val := rune(str[i]) - 'a'
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
	}
	return true
}

/*
If we can't use additional data structures, we can do the following:

1. Compare every character of string to every other character of string. This will take O(n^2) time and O(1) space.
2. If we are allowed to modify the input string, we could sort the string in O(n log(n)) time and then linearly
check the string for neighboring characters that are identical. Careful, though: many sorting algorithms take up
extra space.

These solutions are not as optimal in some respect, but might be better depending on your constraints.
*/
