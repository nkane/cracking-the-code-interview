package binary_to_string

import "strings"

/*
	5.2: Binary to String
	Given a real number between 0 and 1 (e.g., 0.72) that is passed in as a double, print the binary representation.
	If the number cannot be represented accurately in binary with at most 32 characters, print "ERROR".
*/

func printBinary(num float64) string {
	if num >= 1.0 || num <= 0.0 {
		return "ERROR"
	}
	sb := strings.Builder{}
	frac := 0.5
	for num > 0 {
		// set a limit on length, 32 characters
		if sb.Len() >= 32 {
			return "ERROR"
		}
		if num >= frac {
			sb.WriteRune('1')
			num -= frac
		} else {
			sb.WriteRune('0')
		}
		frac /= 2
	}
	return sb.String()
}
