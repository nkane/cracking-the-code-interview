package unique

import (
	"testing"

	"gotest.tools/assert"
)

func TestRealUnique(t *testing.T) {
	s := "abcdefg"
	v := IsUnique(s)
	assert.Assert(t, v == true, "expected unique to be true")
}

func TestNotUnique(t *testing.T) {
	s := "abcdb"
	v := IsUnique(s)
	assert.Assert(t, v == false, "expected unique to be false")
}
