package flip_bit_to_win

import (
	"unsafe"
)

/*
	5.3: Flip Bit to Win
	You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to find the length of the longest
	sequence of 1s you could create.

	Optimal Algorithm
	To reduce the space usage, note that we don't need to hang on to the length of each sequence the entire time. We
	only need it long enough to compare each 1s sequence to the immediately preceding 1s sequence.
	Therefore, we can just walk through the integer doing this, tracking the current 1s sequence length and the
	previous 1s sequence length. When we see a zero, update previousLength:
	- if the next bit is a 1, previousLength should be set to currentLength
	- if the next bit is a 0, then we can't merge these sequences together. So, set previousLength to 0.

	The runtime of this algorithm is still O(b), but we use only O(1) additional memory.
*/

var (
	val       = 0
	INT_BYTES = int(unsafe.Sizeof(val))
)

func flipBit(a int) int {
	// if all 1s, this is already the longest sequence
	if ^a == 0 {
		return INT_BYTES * 8
	}
	currentLength := 0
	previousLength := 0
	maxLength := 1 // we can always have a sequence of at least 1
	for a != 0 {
		if (a & 1) == 1 { // current bit is 1
			currentLength++
		} else if (a & 1) == 0 { // current bit is 0
			// update to 0, if next bit is a 0, or currentLength, if next bit is a 1
			previousLength = currentLength
			if (a & 2) == 0 {
				previousLength = 0
			}
		}
		maxLength = Max(previousLength+currentLength+1, maxLength)
		a = a >> 1
	}
	return maxLength
}

func Max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}
