package string_rotation

import (
	"fmt"
	"strings"
)

/*
	1.9: String Rotation: Assume you have a method isSubstring which checks if one word is a substring
	of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only
	one call to isSubstring (e.g., "watterbottle" is a rotation of "erbottlewat")
*/

func IsStringRotation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	intersectIndex := -1
	firstIdx := 0
	for i := 0; i < len(a); i++ {
		if intersectIndex == -1 && a[firstIdx] == b[i] {
			intersectIndex = i
			firstIdx++
		} else if intersectIndex >= 0 {
			// the next chars have to match
			idx := intersectIndex + (i - intersectIndex)
			if b[idx] != a[firstIdx] {
				return false
			}
			firstIdx++
		}
	}

	x := b[0:intersectIndex]
	fmt.Printf("%s\n", x)

	y := b[intersectIndex:]
	fmt.Printf("%s\n", y)

	sb := strings.Builder{}
	sb.WriteString(y)
	sb.WriteString(x)

	return strings.Contains(a, sb.String())
}
