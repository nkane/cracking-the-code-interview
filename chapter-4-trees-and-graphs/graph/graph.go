package graph

import (
	"os"
	"os/exec"

	graphviz "github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Node[T any] struct {
	Data     T
	Children []*Node[T]
	State    int
}

func (n *Node[T]) AddNeighbor(neighbor *Node[T]) {
	n.Children = append(n.Children, neighbor)
}

type Graph[K comparable, T any] struct {
	Nodes []*Node[T]
	Links map[K]*Node[T]
}

func CreateTestDirectedGraph() *Graph[string, int] {
	g := Graph[string, int]{}
	node0 := Node[int]{
		Data: 0,
	}
	node1 := Node[int]{
		Data: 1,
	}
	node2 := Node[int]{
		Data: 2,
	}
	node3 := Node[int]{
		Data: 3,
	}
	node4 := Node[int]{
		Data: 4,
	}
	node5 := Node[int]{
		Data: 5,
	}
	node0.Children = []*Node[int]{
		&node4,
		&node5,
		&node1,
	}
	node1.Children = []*Node[int]{
		&node3,
		&node4,
	}
	node2.Children = []*Node[int]{
		&node1,
	}
	node3.Children = []*Node[int]{
		&node2,
		&node4,
	}
	g.Nodes = []*Node[int]{
		&node0,
		&node1,
		&node2,
		&node3,
		&node4,
		&node5,
	}
	return &g
}

func CreateTestUndirectedGraph() *Graph[string, int] {
	g := Graph[string, int]{}
	node0 := Node[int]{
		Data: 0,
	}
	node1 := Node[int]{
		Data: 1,
	}
	node2 := Node[int]{
		Data: 2,
	}
	node3 := Node[int]{
		Data: 3,
	}
	node4 := Node[int]{
		Data: 4,
	}
	node5 := Node[int]{
		Data: 5,
	}
	node0.Children = []*Node[int]{
		&node4,
		&node5,
		&node1,
	}
	node1.Children = []*Node[int]{
		&node3,
		&node4,
		&node0,
		&node2,
	}
	node2.Children = []*Node[int]{
		&node1,
	}
	node3.Children = []*Node[int]{
		&node2,
		&node4,
	}
	g.Nodes = []*Node[int]{
		&node0,
		&node1,
		&node2,
		&node3,
		&node4,
		&node5,
	}
	return &g
}

func (g *Graph[K, T]) DrawSVG(directed bool) {
	gv := graphviz.New(graphviz.IntHash)
	if directed {
		gv = graphviz.New(graphviz.IntHash, graphviz.Directed())
	}
	for _, n := range g.Nodes {
		v, ok := any(n.Data).(int)
		if ok {
			gv.AddVertex(v)
		}
	}
	for _, n := range g.Nodes {
		for _, c := range n.Children {
			v, vok := any(n.Data).(int)
			c, cok := any(c.Data).(int)
			if vok && cok {
				gv.AddEdge(v, c)
			}
		}
	}
	file, err := os.Create("graph.gv")
	if err != nil {
		panic(err)
	}
	_ = draw.DOT(gv, file)
	cmd := exec.Command("/usr/bin/dot", "-Tsvg", "-Kneato", "-O", "graph.gv")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
