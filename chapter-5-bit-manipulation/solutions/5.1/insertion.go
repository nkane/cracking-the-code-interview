package insertion

/*
	5.1: Insertion:
	You are given two 32-bit numbers, N and M, and two bit positions, i and j. Write a method to insert M into N such
	that M starts at bit j and ends at bit i. You can assume that the bits j through i have enough space to fit all of
	M. That is, if M=1001, you can assume that there are at leasta 5 bits between j and i. You would not, for example,
	have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.

	Example
	Input:  N = 10000000000, M = 10011, i = 2, j = 6
	Output: N = 10001001100
*/

func UpdateBits(n int, m int, i int, j int) int {
	if i > j || i < 0 || j >= 32 {
		return 0
	}
	// create a mask to clear bits i through j in n.
	// Example: i = 2, j = 4
	// Result should be 11100011, for simplicity, we'll use just 8 bits for the example
	allOnes := ^0
	// 1s before position j, then 0s, left = 11100000
	left := 0
	if j < 31 {
		left = (allOnes << (j + 1))
	}
	// 1's after position i, right = 00000011
	right := ((1 << i) - 1)
	// all 1s, except for 0s, between i and j, mask = 11100011
	mask := left | right
	// clear bits j through i then put m in there
	n_clear := n & mask        // clear bits j through i
	m_shifted := m << i        // move m into correct position
	return n_clear | m_shifted // or them, and we're done!
}
