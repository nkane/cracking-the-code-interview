package graph

import "testing"

func TestPrintGraph(t *testing.T) {
	g := CreateTestDirectedGraph()
	g.DrawSVG(true)
}
