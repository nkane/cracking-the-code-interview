package urlify

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestValidURLify(t *testing.T) {
	v := []rune("Mr John Smith    ")
	l := 13

	URLify(v, l)
	expected := []rune("Mr%20John%20Smith")
	assert.Assert(t, bytes.Equal([]byte(string(v)), []byte(string(expected))), "runes are not equal - got %s - expected %s", string(v), string(expected))
}
