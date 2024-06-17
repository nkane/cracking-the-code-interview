package urlify

/*
	1.3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
	Example:
	Input:	"Mr John Smith    ", 13
	Output: "Mr%20John%20Smith"
*/

func URLify(runes []rune, length int) {
	for i := 0; i < len(runes); i++ {
		// check if rune is equal to space character
		if runes[i] == ' ' {
			// shift everything down
			for j := length - 1; j > i; j-- {
				temp := runes[j+2]
				runes[j+2] = runes[j]
				if temp != ' ' {
					runes[j] = temp
				}
			}

			// insert %20
			runes[i] = '%'
			runes[i+1] = '2'
			runes[i+2] = '0'
			length += 2
		}
	}
}
