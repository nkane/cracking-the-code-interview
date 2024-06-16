package string_compression

import (
	"strconv"
	"strings"
)

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
