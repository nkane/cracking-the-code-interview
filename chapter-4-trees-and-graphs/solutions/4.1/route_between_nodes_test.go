package route_between_nodes

import (
	"graph"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSearch(t *testing.T) {
	g := graph.CreateTestDirectedGraph()
	found := Search(g, g.Nodes[0], g.Nodes[1])
	assert.Assert(t, found == true)
	found = Search(g, g.Nodes[4], g.Nodes[1])
	assert.Assert(t, found == false)
}
