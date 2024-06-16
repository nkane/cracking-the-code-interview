package one_away

import "strings"

/*
	1.5: One Away: There are three types of edits that can be performed on strings: insert a character,
	remove a character, or replace a character. Given two strings, write a function to check if they are
	one edit (or zero edits) away.

	Example:
	pale, ple -> true
	pales, pale -> true
	pale, bale -> true
	pale, bake -> false
*/

func OneAway(a string, b string) bool {
	a = strings.ToLower(a)
	alen := len(a)
	b = strings.ToLower(b)
	m := map[rune]int{}
	for _, c := range a {
		m[c]++
	}
	// must match at least len(a) - 1
	matches := 0
	for _, c := range b {
		if _, ok := m[c]; ok {
			matches++
		}
	}
	if matches != (alen - 1) {
		return false
	}
	return true
}
