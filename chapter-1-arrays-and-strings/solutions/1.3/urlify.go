package urlify

/*
	1.3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
	Example:
	Input:	"Mr John Smith    ", 13
	Output: "Mr%20John%20Smith"
*/

/*
A common approach in string manipulation problems is to edit the string starting from the end and
working backwards. This is useful because we have an extra buffer at the end, which allows us to
change characters without worrying about what we're overwriting.

We will use this approach in this problem. First, we count the number of spaces. We need two extra
characters for each space (since '%20' is two characters longer that ''), so we double this count.
Then, we walk backwards through the string, editing it. When we see a space, we replace it with %20.
If there is no space, then we copy the original character.
*/

func URLify(s []rune, l int) {
	numberOfSpaces := CountOfChar(s, 0, l, ' ')
	newIndex := l - 1 + numberOfSpaces*2

	// if there are excess spaces, add a null character. This indicates that the
	// space after that point have not been replaced with %20.
	if newIndex+1 < len(s) {
		s[newIndex+1] = '\000'
	}
	for oldIndex := l - 1; oldIndex >= 0; oldIndex -= 1 {
		if s[oldIndex] == ' ' {
			// insert %20
			s[newIndex] = '0'
			s[newIndex-1] = '2'
			s[newIndex-2] = '%'
			newIndex -= 3
		} else {
			s[newIndex] = s[oldIndex]
			newIndex -= 1
		}
	}
}

func CountOfChar(s []rune, start int, end int, target rune) int {
	count := 0
	for i := start; i < end; i++ {
		if s[i] == target {
			count++
		}
	}
	return count
}
