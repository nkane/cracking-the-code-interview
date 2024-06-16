package palindrome_permutation

import "strings"

/*
	1.4: Palindrome Permutation: Given a string, write a function that checks if it is a permutation of a palindrome.
	A palindrome is a word or phrase that is the same forward and backwards. A permutation is a rearrangement of
	letters. The palindrome does not need to be limited to just dictionary words. YOu can ignore casing and non-letter
	characters.

	Example:
	Input: Tact Coa
	Output: True (permutations: "taco cat", "atco cta", etc.)

	racecar
	Noon
	Deed
	Level
	Kayak
	Mom
	- there basically needs to be two of each besides a single character
*/

func IsPalindromePermutation(s string) bool {
	// for a string tha tis a permutation, there needs to be two of each character besides a single character
	s = strings.ToLower(s)
	m := map[rune]int{}
	for _, v := range s {
		if v == ' ' {
			continue
		}
		m[v]++
	}
	oneLocated := false
	for _, v := range m {
		if v > 2 {
			return false
		} else if v == 1 {
			if oneLocated {
				return false
			}
			oneLocated = true
		}
	}

	return true
}
