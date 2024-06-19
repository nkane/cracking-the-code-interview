package remove_duplciates

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestLinkListAdd(t *testing.T) {
	expected := `[pos: 1 | data: 1] -> [pos: 2 | data: 2] -> [pos: 3 | data: 3] -> `
	l := &LinkedList{}
	l.Add(&Node{
		Data: 1,
	})
	l.Add(&Node{
		Data: 2,
	})
	l.Add(&Node{
		Data: 3,
	})
	buf := fmt.Sprintf("%v", l)
	assert.Assert(t, expected == buf, "output doesn't match")
}

func TestRemoveDuplicates(t *testing.T) {
}
