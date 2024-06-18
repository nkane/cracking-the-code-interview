package rotate_matrix

import (
	"reflect"
	"testing"

	"gotest.tools/v3/assert"
)

type Test struct {
	Input    [][]int
	Expected [][]int
}

func TestValidRotateMatrix(t *testing.T) {
	tests := []Test{
		{
			Input: [][]int{
				{1, 2},
				{3, 4},
			},
			Expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			Input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			Expected: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, test := range tests {
		RotateMatrix(test.Input)
		assert.Assert(t, reflect.DeepEqual(test.Input, test.Expected) == true, "failed equality check - expected: %+v, got: %+v", test.Expected, test.Input)
	}
}
