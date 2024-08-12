package flip_bit_to_win

import (
	"unsafe"
)

/*
	5.3: Flip Bit to Win
	You have an integer and you can flip exactly one bit from a 0 to a 1. Write code to find the length of the longest
	sequence of 1s you could create.

	Brute Force
	One approach is to convert an integer into an array that rellects the lengths of the 0s and 1s sequences.

	It's O(b) time and O(b) memory, where b is the length of the sequence.
*/

var (
	val       = 0
	INT_BYTES = int(unsafe.Sizeof(val))
)

func longestSequence(n int) int {
	if n == -1 {
		return INT_BYTES * 8
	}
	sequences := getAlternatingSequences(n)
	return findLongestSequence(sequences)
}

func getAlternatingSequences(n int) []int {
	sequences := []int{}
	searchingFor := 0
	counter := 0

	for i := 0; i < INT_BYTES*8; i++ {
		if (n & 1) != searchingFor {
			sequences = append(sequences, counter)
			searchingFor = n & 1 // flip 1 to 0 or 0 to 1
			counter = 0
		}
		counter++
		n = n >> 1
	}

	sequences = append(sequences, counter)
	return sequences
}

func findLongestSequence(seq []int) int {
	maxSeq := 1

	for i := 0; i < len(seq); i += 2 {
		zerosSeq := seq[i]
		onesSeqPrev := 0
		onesSeqNext := 0

		if i-1 >= 0 {
			onesSeqPrev = seq[i-1]
		}
		if i+1 < len(seq) {
			onesSeqNext = seq[i+1]
		}

		thisSeq := 0
		if zerosSeq == 1 { // can merge
			thisSeq = onesSeqPrev + 1 + onesSeqNext
		} else if zerosSeq > 1 { // just add a one to either side
			thisSeq = 1 + max(onesSeqPrev, onesSeqNext)
		} else if zerosSeq == 0 { // no zero, but take either side
			thisSeq = max(onesSeqPrev, onesSeqNext)
		}

		maxSeq = max(maxSeq, thisSeq)
	}

	return maxSeq
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
