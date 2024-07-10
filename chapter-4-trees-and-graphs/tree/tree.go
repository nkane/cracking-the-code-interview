package tree

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/golang-collections/collections/queue"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func BuildTree() *Tree {
	aNode := Node{
		Data: 1,
		Left: &Node{
			Data: 2,
			Left: &Node{
				Data: 3,
			},
			Right: &Node{
				Data: 4,
				Left: &Node{
					Data: 5,
				},
			},
		},
		Right: &Node{
			Data: 6,
			Right: &Node{
				Data: 7,
				Left: &Node{
					Data: 8,
					Left: &Node{
						Data: 9,
					},
					Right: &Node{
						Data: 10,
					},
				},
			},
		},
	}
	return &Tree{
		Root: &aNode,
	}
}

func (t *Tree) DrawSVG(name string) {
	gv := graph.New(graph.IntHash, graph.Directed())
	q := queue.New()
	q.Enqueue(t.Root)
	var n *Node
	for q.Len() > 0 {
		n = q.Dequeue().(*Node)
		if n != nil {
			gv.AddVertex(n.Data)
			if n.Left != nil {
				gv.AddVertex(n.Left.Data)
				gv.AddEdge(n.Data, n.Left.Data)
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				gv.AddVertex(n.Right.Data)
				gv.AddEdge(n.Data, n.Right.Data)
				q.Enqueue(n.Right)
			}
		}
	}
	file, err := os.Create(fmt.Sprintf("tree-%s.gv", name))
	if err != nil {
		panic(err)
	}
	_ = draw.DOT(gv, file)
	cmd := exec.Command("/usr/bin/dot", "-Tsvg", "-Kdot", "-O", fmt.Sprintf("tree-%s.gv", name))
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
