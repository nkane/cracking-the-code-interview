package string_compression

import (
	"strconv"
	"strings"
)

/*
	1.6: String Compression: Implement a method to perform basic string compression using the counts
	of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the
	"compressed" string would not become smaller than the original string, your method should return
	the original string. You can assume the string has only uppercase and lowercase letters (a - z).
*/

/*
Solution:

At first glance, implementing this method seems fairly straightforward, but perhaps a bit tedious. We
iterate through the string, copying characters to a new string and counting the repeats. At each
iteration, check if the current character is the same as the next character. If not, add its
compressed version to the result. How hard could this be?
*/

type CompressString func(str string) string

func CompressStringBad(str string) string {
	compressedString := ""
	countConsecutive := 0
	for i := 0; i < len(str); i++ {
		countConsecutive++
		// If next character is different than current, append this char to result
		if i+1 >= len(str) || str[i] != str[i+1] {
			compressedString += "" + string(str[i]) + strconv.Itoa(countConsecutive)
			countConsecutive = 0
		}
	}
	if len(compressedString) < len(str) {
		return compressedString
	}
	return str
}

/*
This works. It is efficient, though? Take a look at the runtime of this code.

The runtime is O(p + k^2), where p is the size of the original string and k is the number of character
sequences. For example, if the setring is aabccdeeaa, then there are six character sequences. It's slow
because string concatenation operates in O(n^2) time.

We can fix this by using a StringBuilder.
*/
func CompressStringSB(str string) string {
	compressed := strings.Builder{}
	countConsecutive := 0
	for i := 0; i < len(str); i++ {
		countConsecutive++
		// If next character is different than current, append this char to result
		if i+1 >= len(str) || str[i] != str[i+1] {
			compressed.WriteRune(rune(str[i]))
			compressed.WriteString(strconv.Itoa(countConsecutive))
			countConsecutive = 0
		}
	}
	if compressed.Len() < len(str) {
		return compressed.String()
	}
	return str
}

/*
Both of these solutions create the compressed string first and then return the shorter of the input string
and the compressed string.

Instead, we can check in advance. This will be more optimal in cases where we don't have a larger number
of repeating characters. It will avoid us having to create a string that we never use. The downside of this
is that it causes a second loop through the characters and also adds nearly duplicate code.

One other benefit to this approach is that we can initialize StringBuilder to it's necessary capacity up-front.
Without this, StringBuilder will (behind the scenes) need to double its capacity every time it hits capacity.
The capacity could b double what we ultimately need.
*/

func CompressStringCheckFirst(str string) string {
	finalLength := CountCompression(str)
	if finalLength >= len(str) {
		return str
	}
	compressed := strings.Builder{}
	compressed.Grow(finalLength)
	countConsecutive := 0
	for i := 0; i < len(str); i++ {
		countConsecutive++
		if i+1 >= len(str) || str[i] != str[i+1] {
			compressed.WriteRune(rune(str[i]))
			compressed.WriteString(strconv.Itoa(countConsecutive))
			countConsecutive = 0
		}
	}
	return compressed.String()
}

func CountCompression(str string) int {
	compressedLength := 0
	countConsecutive := 0
	for i := 0; i < len(str); i++ {
		countConsecutive++
		// If next character is different than current, append this char to result
		if i+1 >= len(str) || str[i] != str[i+1] {
			compressedLength += 1 + len(strconv.Itoa(countConsecutive))
			countConsecutive = 0
		}
	}
	return compressedLength
}
