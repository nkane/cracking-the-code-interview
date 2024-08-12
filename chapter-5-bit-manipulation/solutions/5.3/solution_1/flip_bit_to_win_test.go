package flip_bit_to_win

import (
	"testing"

	"gotest.tools/assert"
)

func TestLongestSequence(t *testing.T) {
	tests := []struct {
		Name   string
		Input  int
		Output int
	}{
		{
			Name:   "Example 1",
			Input:  -1,
			Output: 64,
		},
		{
			Name:   "Example 2",
			Input:  1775,
			Output: 8,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := longestSequence(test.Input)
			assert.Assert(t, test.Output == got)
		})
	}
}
