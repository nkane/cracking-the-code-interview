package string_rotation

import (
	"strings"
)

/*
	1.9: String Rotation: Assume you have a method isSubstring which checks if one word is a substring
	of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only
	one call to isSubstring (e.g., "watterbottle" is a rotation of "erbottlewat")
*/

/*
If we imagine that s2 is a rotation of s1, then we can ask what the rotation point is. For example, if you
rotate waterbottole after wat, you get erbottolewat. In a rotation, we cut s1 into two parts, x and y,
and rearrange them to get s2.

s1 = xy = waterbottle
x = wat
y = erbottle
s2 = yx = erbottlewat

So, we need to check if there's a way to split s1 into x and y such that xy = s1 and yx = s2. Regardless of
where the division between x and y is, we can see that yx will always be a substring of xyxy. That is, s2 will
always be a substring of s1s1.

And this is precisely how we solve the problem: simple do isSubstring(s1s1, s2)

The runtime of this varies based on the runtime of isSubstring, or in this case strings.Contains. But if you assume
that isSubstring runs in O(A + B) time (on strings of length A and B), then the runtime of IsStringRotation is O(N).
*/

func IsStringRotation(s1 string, s2 string) bool {
	l := len(s1)
	// check that s1 and s2 are equal length and not empty
	if l == len(s2) && l > 0 {
		// concatenate s1 and s1 within new buffer
		s1s1 := s1 + s1
		return strings.Contains(s1s1, s2)
	}
	return false
}
