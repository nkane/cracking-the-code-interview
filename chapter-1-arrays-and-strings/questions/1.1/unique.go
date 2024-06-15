package unique

/*
1.1: Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use
additional data structures?
*/

func IsUnique(s string) bool {
	charMap := [256]int{0}
	for _, c := range s {
		if charMap[int(c)] == 0 {
			charMap[int(c)] = 1
		} else {
			return false
		}
	}
	return true
}
