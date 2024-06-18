package zero_matrix

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

type Test struct {
	Input      [][]int
	Expected   [][]int
	ZeroMatrix ZeroMatrix
}

func TestZeroMatrix(t *testing.T) {
	tests := []Test{
		{
			Input: [][]int{
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			Expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 0, 0},
			},
			ZeroMatrix: ZeroMatrixInefficient,
		},
		{
			Input: [][]int{
				{1, 2},
				{3, 4},
			},
			Expected: [][]int{
				{1, 2},
				{3, 4},
			},
			ZeroMatrix: ZeroMatrixInefficient,
		},
		{
			Input: [][]int{
				{1, 2},
				{0, 4},
			},
			Expected: [][]int{
				{0, 2},
				{0, 0},
			},
			ZeroMatrix: ZeroMatrixInefficient,
		},
		{
			Input: [][]int{
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			Expected: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 0, 0},
			},
			ZeroMatrix: ZeroMatrixEfficient,
		},
		{
			Input: [][]int{
				{1, 2},
				{3, 4},
			},
			Expected: [][]int{
				{1, 2},
				{3, 4},
			},
			ZeroMatrix: ZeroMatrixEfficient,
		},
		{
			Input: [][]int{
				{1, 2},
				{0, 4},
			},
			Expected: [][]int{
				{0, 2},
				{0, 0},
			},
			ZeroMatrix: ZeroMatrixEfficient,
		},
	}
	for _, test := range tests {
		test.ZeroMatrix(test.Input)
		assert.Assert(t, reflect.DeepEqual(test.Input, test.Expected) == true, "failed check - expected: %+v, got %+v", test.Expected, test.Input)
	}
}
