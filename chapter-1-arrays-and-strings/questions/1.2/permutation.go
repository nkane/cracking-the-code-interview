package permutation

/*
	1.2: Check Permutation: Given two strings, write a method to decide if one is a permutation of the
	other.
*/

func IsPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	checkMap := make(map[rune]int, 0)
	for _, char := range a {
		checkMap[char]++
	}
	for _, char := range b {
		if v, ok := checkMap[char]; !ok || (v <= 0) {
			return false
		}
		checkMap[char]--
	}
	return true
}
