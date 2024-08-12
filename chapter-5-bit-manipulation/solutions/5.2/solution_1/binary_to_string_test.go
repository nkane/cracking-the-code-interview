package binary_to_string

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinaryToString(t *testing.T) {
	tests := []struct {
		Name   string
		Input  float64
		Output string
	}{
		{
			Name:   "Example 1",
			Input:  3.14,
			Output: "ERROR",
		},
		{
			Name:   "Example 2",
			Input:  0.25,
			Output: "01",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := printBinary(test.Input)
			assert.Assert(t, test.Output == got)
		})
	}
}
