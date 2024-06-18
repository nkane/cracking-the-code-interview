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

func CompressString(s string) string {
	sb := strings.Builder{}
	previousChar := rune(s[0])
	count := 1
	for i := 1; i < len(s); i++ {
		currentChar := rune(s[i])
		if currentChar == previousChar {
			count++
		} else {
			sb.WriteRune(previousChar)
			sb.WriteString(strconv.Itoa(count))
			count = 1
			previousChar = currentChar
		}
	}
	sb.WriteRune(previousChar)
	sb.WriteString(strconv.Itoa(count))
	if sb.Len() >= len(s) {
		return s
	}
	return sb.String()
}
