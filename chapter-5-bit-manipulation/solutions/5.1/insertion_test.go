package insertion

import (
	"testing"

	"gotest.tools/assert"
)

func TestUpdateBits(t *testing.T) {
	tests := []struct {
		Name   string
		Input  int
		M      int
		I      int
		J      int
		Output int
	}{
		{
			Name:   "Test",
			Input:  0xFF00FF00,
			M:      0xFFAAFFFF,
			I:      2,
			J:      6,
			Output: 0x3FFABFFFC,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := UpdateBits(test.Input, test.M, test.I, test.J)
			assert.Assert(t, test.Output == got)
		})
	}
}
